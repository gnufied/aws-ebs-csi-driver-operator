// Code generated by go-bindata.
// sources:
// assets/controller_deployment.yaml
// assets/controller_sa.yaml
// assets/csidriver.yaml
// assets/namespace.yaml
// assets/node_daemonset.yaml
// assets/node_sa.yaml
// assets/rbac/attacher_binding.yaml
// assets/rbac/attacher_role.yaml
// assets/rbac/controller_privileged_binding.yaml
// assets/rbac/node_privileged_binding.yaml
// assets/rbac/privileged_role.yaml
// assets/rbac/provisioner_binding.yaml
// assets/rbac/provisioner_role.yaml
// assets/rbac/resizer_binding.yaml
// assets/rbac/resizer_role.yaml
// assets/rbac/snapshotter_binding.yaml
// assets/rbac/snapshotter_role.yaml
// assets/storageclass.yaml
// DO NOT EDIT!

package generated

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _controller_deploymentYaml = []byte(`kind: Deployment
apiVersion: apps/v1
metadata:
  name: aws-ebs-csi-driver
  namespace: openshift-aws-ebs-csi-driver
spec:
  selector:
    matchLabels:
      app: aws-ebs-csi-driver
  serviceName: aws-ebs-csi-driver
  replicas: 1
  template:
    metadata:
      labels:
        app: aws-ebs-csi-driver
    spec:
      hostNetwork: true
      serviceAccount: aws-ebs-csi-driver-controller-sa
      priorityClassName: system-cluster-critical
      tolerations:
        - key: CriticalAddonsOnly
          operator: Exists
      containers:
        - name: ebs-plugin
          image: amazon/aws-ebs-csi-driver:latest
          args:
            - --endpoint=$(CSI_ENDPOINT)
            - --logtostderr
            - --v=5
          env:
            - name: CSI_ENDPOINT
              value: unix:///var/lib/csi/sockets/pluginproxy/csi.sock
            - name: AWS_ACCESS_KEY_ID
              valueFrom:
                secretKeyRef:
                  name: aws-cloud-credentials
                  key: aws_access_key_id
                  optional: true
            - name: AWS_SECRET_ACCESS_KEY
              valueFrom:
                secretKeyRef:
                  name: aws-cloud-credentials
                  key: aws_secret_access_key
                  optional: true
          ports:
            - name: healthz
              containerPort: 19808
              protocol: TCP
          volumeMounts:
            - name: socket-dir
              mountPath: /var/lib/csi/sockets/pluginproxy/
        - name: csi-provisioner
          image: quay.io/k8scsi/csi-provisioner:canary
          args:
            - --provisioner=ebs.csi.aws.com
            - --csi-address=$(ADDRESS)
            - --v=5
            - --feature-gates=Topology=true
          env:
            - name: ADDRESS
              value: /var/lib/csi/sockets/pluginproxy/csi.sock
          volumeMounts:
            - name: socket-dir
              mountPath: /var/lib/csi/sockets/pluginproxy/
        - name: csi-attacher
          image: quay.io/k8scsi/csi-attacher:canary
          args:
            - --csi-address=$(ADDRESS)
            - --v=5
          env:
            - name: ADDRESS
              value: /var/lib/csi/sockets/pluginproxy/csi.sock
          volumeMounts:
            - name: socket-dir
              mountPath: /var/lib/csi/sockets/pluginproxy/
        - name: csi-resizer
          image: quay.io/k8scsi/csi-resizer:canary
          args:
            - --csi-address=$(ADDRESS)
            - --v=5
          env:
            - name: ADDRESS
              value: /var/lib/csi/sockets/pluginproxy/csi.sock
          volumeMounts:
            - name: socket-dir
              mountPath: /var/lib/csi/sockets/pluginproxy/
        - name: csi-snapshotter
          image: quay.io/k8scsi/csi-snapshotter:canary
          args:
            - --csi-address=$(ADDRESS)
            - --v=5
          env:
          - name: ADDRESS
            value: /var/lib/csi/sockets/pluginproxy/csi.sock
          volumeMounts:
          - mountPath: /var/lib/csi/sockets/pluginproxy/
            name: socket-dir
      volumes:
        - name: socket-dir
          emptyDir: {}
`)

func controller_deploymentYamlBytes() ([]byte, error) {
	return _controller_deploymentYaml, nil
}

func controller_deploymentYaml() (*asset, error) {
	bytes, err := controller_deploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "controller_deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _controller_saYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: aws-ebs-csi-driver-controller-sa
  namespace: openshift-aws-ebs-csi-driver
`)

func controller_saYamlBytes() ([]byte, error) {
	return _controller_saYaml, nil
}

func controller_saYaml() (*asset, error) {
	bytes, err := controller_saYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "controller_sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _csidriverYaml = []byte(`apiVersion: storage.k8s.io/v1beta1
kind: CSIDriver
metadata:
  name: ebs.csi.aws.com
spec:
  attachRequired: true
  podInfoOnMount: false
`)

func csidriverYamlBytes() ([]byte, error) {
	return _csidriverYaml, nil
}

func csidriverYaml() (*asset, error) {
	bytes, err := csidriverYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "csidriver.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _namespaceYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  name: openshift-aws-ebs-csi-driver
`)

func namespaceYamlBytes() ([]byte, error) {
	return _namespaceYaml, nil
}

func namespaceYaml() (*asset, error) {
	bytes, err := namespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "namespace.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _node_daemonsetYaml = []byte(`kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: aws-ebs-csi-driver
  namespace: openshift-aws-ebs-csi-driver
spec:
  selector:
    matchLabels:
      app: aws-ebs-csi-driver
  template:
    metadata:
      labels:
        app: aws-ebs-csi-driver
    spec:
      hostNetwork: true
      serviceAccount: aws-ebs-csi-driver-node-sa
      priorityClassName: system-node-critical
      tolerations:
        - key: CriticalAddonsOnly
          operator: Exists
      containers:
        - name: ebs-plugin
          securityContext:
            privileged: true
          image: amazon/aws-ebs-csi-driver:latest
          args:
            - --endpoint=$(CSI_ENDPOINT)
            - --logtostderr
            - --v=5
          env:
            - name: CSI_ENDPOINT
              value: unix:/csi/csi.sock
          volumeMounts:
            - name: kubelet-dir
              mountPath: /var/lib/kubelet
              mountPropagation: "Bidirectional"
            - name: plugin-dir
              mountPath: /csi
            - name: device-dir
              mountPath: /dev
          ports:
            - name: healthz
              containerPort: 9808
              protocol: TCP
          livenessProbe:
            httpGet:
              path: /healthz
              port: healthz
            initialDelaySeconds: 10
            timeoutSeconds: 3
            periodSeconds: 10
            failureThreshold: 5
        - name: node-driver-registrar
          securityContext:
            privileged: true
          image: quay.io/k8scsi/csi-node-driver-registrar:canary
          args:
            - --csi-address=$(ADDRESS)
            - --kubelet-registration-path=$(DRIVER_REG_SOCK_PATH)
            - --v=5
          lifecycle:
            preStop:
              exec:
                command: ["/bin/sh", "-c", "rm -rf /registration/ebs.csi.aws.com-reg.sock /csi/csi.sock"]
          env:
            - name: ADDRESS
              value: /csi/csi.sock
            - name: DRIVER_REG_SOCK_PATH
              value: /var/lib/kubelet/plugins/ebs.csi.aws.com/csi.sock
          volumeMounts:
            - name: plugin-dir
              mountPath: /csi
            - name: registration-dir
              mountPath: /registration
        - name: liveness-probe
          image: quay.io/k8scsi/livenessprobe:canary
          args:
            - --csi-address=/csi/csi.sock
            - --probe-timeout=3s
          volumeMounts:
            - name: plugin-dir
              mountPath: /csi
      volumes:
        - name: kubelet-dir
          hostPath:
            path: /var/lib/kubelet
            type: Directory
        - name: plugin-dir
          hostPath:
            path: /var/lib/kubelet/plugins/ebs.csi.aws.com/
            type: DirectoryOrCreate
        - name: registration-dir
          hostPath:
            path: /var/lib/kubelet/plugins_registry/
            type: Directory
        - name: device-dir
          hostPath:
            path: /dev
            type: Directory
`)

func node_daemonsetYamlBytes() ([]byte, error) {
	return _node_daemonsetYaml, nil
}

func node_daemonsetYaml() (*asset, error) {
	bytes, err := node_daemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node_daemonset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _node_saYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: aws-ebs-csi-driver-node-sa
  namespace: openshift-aws-ebs-csi-driver
`)

func node_saYamlBytes() ([]byte, error) {
	return _node_saYaml, nil
}

func node_saYaml() (*asset, error) {
	bytes, err := node_saYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node_sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacAttacher_bindingYaml = []byte(`kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ebs-csi-attacher-binding
subjects:
  - kind: ServiceAccount
    name: aws-ebs-csi-driver-controller-sa
    namespace: openshift-aws-ebs-csi-driver
roleRef:
  kind: ClusterRole
  name: ebs-external-attacher-role
  apiGroup: rbac.authorization.k8s.io
`)

func rbacAttacher_bindingYamlBytes() ([]byte, error) {
	return _rbacAttacher_bindingYaml, nil
}

func rbacAttacher_bindingYaml() (*asset, error) {
	bytes, err := rbacAttacher_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/attacher_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacAttacher_roleYaml = []byte(`kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ebs-external-attacher-role
rules:
  - apiGroups: [""]
    resources: ["persistentvolumes"]
    verbs: ["get", "list", "watch", "update", "patch"]
  - apiGroups: [""]
    resources: ["nodes"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["csi.storage.k8s.io"]
    resources: ["csinodeinfos"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["volumeattachments"]
    verbs: ["get", "list", "watch", "update", "patch"]
`)

func rbacAttacher_roleYamlBytes() ([]byte, error) {
	return _rbacAttacher_roleYaml, nil
}

func rbacAttacher_roleYaml() (*asset, error) {
	bytes, err := rbacAttacher_roleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/attacher_role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacController_privileged_bindingYaml = []byte(`kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ebs-controller-privileged-binding
subjects:
  - kind: ServiceAccount
    name: aws-ebs-csi-driver-controller-sa
    namespace: openshift-aws-ebs-csi-driver
roleRef:
  kind: ClusterRole
  name: ebs-privileged-role
  apiGroup: rbac.authorization.k8s.io
`)

func rbacController_privileged_bindingYamlBytes() ([]byte, error) {
	return _rbacController_privileged_bindingYaml, nil
}

func rbacController_privileged_bindingYaml() (*asset, error) {
	bytes, err := rbacController_privileged_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/controller_privileged_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacNode_privileged_bindingYaml = []byte(`kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ebs-node-privileged-binding
subjects:
  - kind: ServiceAccount
    name: aws-ebs-csi-driver-node-sa
    namespace: openshift-aws-ebs-csi-driver
roleRef:
  kind: ClusterRole
  name: ebs-privileged-role
  apiGroup: rbac.authorization.k8s.io
`)

func rbacNode_privileged_bindingYamlBytes() ([]byte, error) {
	return _rbacNode_privileged_bindingYaml, nil
}

func rbacNode_privileged_bindingYaml() (*asset, error) {
	bytes, err := rbacNode_privileged_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/node_privileged_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacPrivileged_roleYaml = []byte(`# TODO: create custom SCC with things that the AWS CSI driver needs

kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ebs-privileged-role
rules:
  - apiGroups: ["security.openshift.io"]
    resourceNames: ["privileged"]
    resources: ["securitycontextconstraints"]
    verbs: ["use"]
`)

func rbacPrivileged_roleYamlBytes() ([]byte, error) {
	return _rbacPrivileged_roleYaml, nil
}

func rbacPrivileged_roleYaml() (*asset, error) {
	bytes, err := rbacPrivileged_roleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/privileged_role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacProvisioner_bindingYaml = []byte(`kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ebs-csi-provisioner-binding
subjects:
  - kind: ServiceAccount
    name: aws-ebs-csi-driver-controller-sa
    namespace: openshift-aws-ebs-csi-driver
roleRef:
  kind: ClusterRole
  name: ebs-external-provisioner-role
  apiGroup: rbac.authorization.k8s.io
`)

func rbacProvisioner_bindingYamlBytes() ([]byte, error) {
	return _rbacProvisioner_bindingYaml, nil
}

func rbacProvisioner_bindingYaml() (*asset, error) {
	bytes, err := rbacProvisioner_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/provisioner_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacProvisioner_roleYaml = []byte(`kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ebs-external-provisioner-role
rules:
  - apiGroups: [""]
    resources: ["persistentvolumes"]
    verbs: ["get", "list", "watch", "create", "delete"]
  - apiGroups: [""]
    resources: ["persistentvolumeclaims"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["storageclasses"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["events"]
    verbs: ["get", "list", "watch", "create", "update", "patch"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["csinodes"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["nodes"]
    verbs: ["get", "list", "watch"]
`)

func rbacProvisioner_roleYamlBytes() ([]byte, error) {
	return _rbacProvisioner_roleYaml, nil
}

func rbacProvisioner_roleYaml() (*asset, error) {
	bytes, err := rbacProvisioner_roleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/provisioner_role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacResizer_bindingYaml = []byte(`kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ebs-csi-resizer-binding
subjects:
  - kind: ServiceAccount
    name: aws-ebs-csi-driver-controller-sa
    namespace: openshift-aws-ebs-csi-driver
roleRef:
  kind: ClusterRole
  name: ebs-external-resizer-role
  apiGroup: rbac.authorization.k8s.io
`)

func rbacResizer_bindingYamlBytes() ([]byte, error) {
	return _rbacResizer_bindingYaml, nil
}

func rbacResizer_bindingYaml() (*asset, error) {
	bytes, err := rbacResizer_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/resizer_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacResizer_roleYaml = []byte(`kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ebs-external-resizer-role
rules:
  - apiGroups: [""]
    resources: ["persistentvolumes"]
    verbs: ["get", "list", "watch", "update", "patch"]
  - apiGroups: [""]
    resources: ["persistentvolumeclaims"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["persistentvolumeclaims/status"]
    verbs: ["update", "patch"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["storageclasses"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["events"]
    verbs: ["list", "watch", "create", "update", "patch"]
`)

func rbacResizer_roleYamlBytes() ([]byte, error) {
	return _rbacResizer_roleYaml, nil
}

func rbacResizer_roleYaml() (*asset, error) {
	bytes, err := rbacResizer_roleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/resizer_role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacSnapshotter_bindingYaml = []byte(`kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ebs-csi-snapshotter-binding
subjects:
  - kind: ServiceAccount
    name: aws-ebs-csi-driver-controller-sa
    namespace: openshift-aws-ebs-csi-driver
roleRef:
  kind: ClusterRole
  name: ebs-external-snapshotter-role
  apiGroup: rbac.authorization.k8s.io
`)

func rbacSnapshotter_bindingYamlBytes() ([]byte, error) {
	return _rbacSnapshotter_bindingYaml, nil
}

func rbacSnapshotter_bindingYaml() (*asset, error) {
	bytes, err := rbacSnapshotter_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/snapshotter_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacSnapshotter_roleYaml = []byte(`kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ebs-external-snapshotter-role
rules:
- apiGroups: [""]
  resources: ["persistentvolumes"]
  verbs: ["get", "list", "watch"]
- apiGroups: [""]
  resources: ["persistentvolumeclaims"]
  verbs: ["get", "list", "watch"]
- apiGroups: ["storage.k8s.io"]
  resources: ["storageclasses"]
  verbs: ["get", "list", "watch"]
- apiGroups: [""]
  resources: ["events"]
  verbs: ["list", "watch", "create", "update", "patch"]
- apiGroups: [""]
  resources: ["secrets"]
  verbs: ["get", "list"]
- apiGroups: ["snapshot.storage.k8s.io"]
  resources: ["volumesnapshotclasses"]
  verbs: ["get", "list", "watch"]
- apiGroups: ["snapshot.storage.k8s.io"]
  resources: ["volumesnapshotcontents"]
  verbs: ["create", "get", "list", "watch", "update", "delete"]
- apiGroups: ["snapshot.storage.k8s.io"]
  resources: ["volumesnapshotcontents/status"]
  verbs: ["update"]
- apiGroups: ["snapshot.storage.k8s.io"]
  resources: ["volumesnapshots"]
  verbs: ["get", "list", "watch", "update"]
- apiGroups: ["apiextensions.k8s.io"]
  resources: ["customresourcedefinitions"]
  verbs: ["create", "list", "watch", "delete"]
- apiGroups: ["coordination.k8s.io"]
  resources: ["leases"]
  verbs: ["get", "watch", "list", "delete", "update", "create"]
`)

func rbacSnapshotter_roleYamlBytes() ([]byte, error) {
	return _rbacSnapshotter_roleYaml, nil
}

func rbacSnapshotter_roleYaml() (*asset, error) {
	bytes, err := rbacSnapshotter_roleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/snapshotter_role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _storageclassYaml = []byte(`apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: ebs-sc
provisioner: ebs.csi.aws.com
reclaimPolicy: "Delete"
volumeBindingMode: WaitForFirstConsumer
`)

func storageclassYamlBytes() ([]byte, error) {
	return _storageclassYaml, nil
}

func storageclassYaml() (*asset, error) {
	bytes, err := storageclassYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "storageclass.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"controller_deployment.yaml":              controller_deploymentYaml,
	"controller_sa.yaml":                      controller_saYaml,
	"csidriver.yaml":                          csidriverYaml,
	"namespace.yaml":                          namespaceYaml,
	"node_daemonset.yaml":                     node_daemonsetYaml,
	"node_sa.yaml":                            node_saYaml,
	"rbac/attacher_binding.yaml":              rbacAttacher_bindingYaml,
	"rbac/attacher_role.yaml":                 rbacAttacher_roleYaml,
	"rbac/controller_privileged_binding.yaml": rbacController_privileged_bindingYaml,
	"rbac/node_privileged_binding.yaml":       rbacNode_privileged_bindingYaml,
	"rbac/privileged_role.yaml":               rbacPrivileged_roleYaml,
	"rbac/provisioner_binding.yaml":           rbacProvisioner_bindingYaml,
	"rbac/provisioner_role.yaml":              rbacProvisioner_roleYaml,
	"rbac/resizer_binding.yaml":               rbacResizer_bindingYaml,
	"rbac/resizer_role.yaml":                  rbacResizer_roleYaml,
	"rbac/snapshotter_binding.yaml":           rbacSnapshotter_bindingYaml,
	"rbac/snapshotter_role.yaml":              rbacSnapshotter_roleYaml,
	"storageclass.yaml":                       storageclassYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"controller_deployment.yaml": {controller_deploymentYaml, map[string]*bintree{}},
	"controller_sa.yaml":         {controller_saYaml, map[string]*bintree{}},
	"csidriver.yaml":             {csidriverYaml, map[string]*bintree{}},
	"namespace.yaml":             {namespaceYaml, map[string]*bintree{}},
	"node_daemonset.yaml":        {node_daemonsetYaml, map[string]*bintree{}},
	"node_sa.yaml":               {node_saYaml, map[string]*bintree{}},
	"rbac": {nil, map[string]*bintree{
		"attacher_binding.yaml":              {rbacAttacher_bindingYaml, map[string]*bintree{}},
		"attacher_role.yaml":                 {rbacAttacher_roleYaml, map[string]*bintree{}},
		"controller_privileged_binding.yaml": {rbacController_privileged_bindingYaml, map[string]*bintree{}},
		"node_privileged_binding.yaml":       {rbacNode_privileged_bindingYaml, map[string]*bintree{}},
		"privileged_role.yaml":               {rbacPrivileged_roleYaml, map[string]*bintree{}},
		"provisioner_binding.yaml":           {rbacProvisioner_bindingYaml, map[string]*bintree{}},
		"provisioner_role.yaml":              {rbacProvisioner_roleYaml, map[string]*bintree{}},
		"resizer_binding.yaml":               {rbacResizer_bindingYaml, map[string]*bintree{}},
		"resizer_role.yaml":                  {rbacResizer_roleYaml, map[string]*bintree{}},
		"snapshotter_binding.yaml":           {rbacSnapshotter_bindingYaml, map[string]*bintree{}},
		"snapshotter_role.yaml":              {rbacSnapshotter_roleYaml, map[string]*bintree{}},
	}},
	"storageclass.yaml": {storageclassYaml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
