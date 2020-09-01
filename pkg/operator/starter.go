package operator

import (
	"context"
	"fmt"

	dynamicclient "k8s.io/client-go/dynamic"
	kubeclient "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"

	opv1 "github.com/openshift/api/operator/v1"
	"github.com/openshift/library-go/pkg/controller/controllercmd"
	"github.com/openshift/library-go/pkg/operator/csi/csicontrollerset"
	goc "github.com/openshift/library-go/pkg/operator/genericoperatorclient"
	"github.com/openshift/library-go/pkg/operator/v1helpers"

	"github.com/openshift/aws-ebs-csi-driver-operator/pkg/generated"
)

const (
	// Operand and operator run in the same namespace
	defaultNamespace = "openshift-cluster-csi-drivers"
	operatorName     = "aws-ebs-csi-driver-operator"
	operandName      = "aws-ebs-csi-driver"
	instanceName     = "ebs.csi.aws.com"
)

func RunOperator(ctx context.Context, controllerConfig *controllercmd.ControllerContext) error {
	// Create clientsets and informers
	kubeClient := kubeclient.NewForConfigOrDie(rest.AddUserAgent(controllerConfig.KubeConfig, operatorName))
	dynamicClient := dynamicclient.NewForConfigOrDie(rest.AddUserAgent(controllerConfig.KubeConfig, operatorName))
	kubeInformersForNamespaces := v1helpers.NewKubeInformersForNamespaces(kubeClient, defaultNamespace, "")

	// Create GenericOperatorclient. This is used by the library-go controllers created down below
	gvr := opv1.SchemeGroupVersion.WithResource("clustercsidrivers")
	operatorClient, dynamicInformers, err := goc.NewClusterScopedOperatorClientWithConfigName(controllerConfig.KubeConfig, gvr, instanceName)
	if err != nil {
		return err
	}

	csiControllerSet := csicontrollerset.NewCSIControllerSet(
		operatorClient,
		controllerConfig.EventRecorder,
	).WithLogLevelController().WithManagementStateController(
		operandName,
		false,
	).WithStaticResourcesController(
		"AWSEBSDriverStaticResourcesController",
		kubeClient,
		kubeInformersForNamespaces,
		generated.Asset,
		[]string{
			"storageclass.yaml",
			"csidriver.yaml",
			"controller_sa.yaml",
			"node_sa.yaml",
			"rbac/attacher_role.yaml",
			"rbac/attacher_binding.yaml",
			"rbac/privileged_role.yaml",
			"rbac/controller_privileged_binding.yaml",
			"rbac/node_privileged_binding.yaml",
			"rbac/provisioner_role.yaml",
			"rbac/provisioner_binding.yaml",
			"rbac/resizer_role.yaml",
			"rbac/resizer_binding.yaml",
			"rbac/snapshotter_role.yaml",
			"rbac/snapshotter_binding.yaml",
		},
	).WithCredentialsRequestController(
		"AWSEBSDriverCredentialsRequestController",
		defaultNamespace,
		generated.MustAsset,
		"credentials.yaml",
		dynamicClient,
	).WithCSIDriverControllerService(
		"AWSEBSDriverControllerServiceController",
		generated.MustAsset,
		"controller.yaml",
		kubeClient,
		kubeInformersForNamespaces.InformersFor(defaultNamespace),
	).WithCSIDriverNodeService(
		"AWSEBSDriverNodeServiceController",
		generated.MustAsset,
		"node.yaml",
		kubeClient,
		kubeInformersForNamespaces.InformersFor(defaultNamespace),
	)

	if err != nil {
		return err
	}

	klog.Info("Starting informers")
	go kubeInformersForNamespaces.Start(ctx.Done())
	go dynamicInformers.Start(ctx.Done())

	klog.Info("Starting controller set")
	go csiControllerSet.Run(ctx, 1)

	<-ctx.Done()

	return fmt.Errorf("stopped")
}
