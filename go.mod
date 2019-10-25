module github.com/rancher/wrangler-api

go 1.12

replace (
	// if you depend on this module you should copy all these replace statements below to keep your sanity
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.2.0+incompatible
	k8s.io/api => github.com/rancher/kubernetes/staging/src/k8s.io/api v1.16.2-k3s.1
	k8s.io/apiextensions-apiserver => github.com/rancher/kubernetes/staging/src/k8s.io/apiextensions-apiserver v1.16.2-k3s.1
	k8s.io/apimachinery => github.com/rancher/kubernetes/staging/src/k8s.io/apimachinery v1.16.2-k3s.1
	k8s.io/apiserver => github.com/rancher/kubernetes/staging/src/k8s.io/apiserver v1.16.2-k3s.1
	k8s.io/cli-runtime => github.com/rancher/kubernetes/staging/src/k8s.io/cli-runtime v1.16.2-k3s.1
	k8s.io/client-go => github.com/rancher/kubernetes/staging/src/k8s.io/client-go v1.16.2-k3s.1
	k8s.io/cloud-provider => github.com/rancher/kubernetes/staging/src/k8s.io/cloud-provider v1.16.2-k3s.1
	k8s.io/cluster-bootstrap => github.com/rancher/kubernetes/staging/src/k8s.io/cluster-bootstrap v1.16.2-k3s.1
	k8s.io/code-generator => github.com/rancher/kubernetes/staging/src/k8s.io/code-generator v1.16.2-k3s.1
	k8s.io/component-base => github.com/rancher/kubernetes/staging/src/k8s.io/component-base v1.16.2-k3s.1
	k8s.io/cri-api => github.com/rancher/kubernetes/staging/src/k8s.io/cri-api v1.16.2-k3s.1
	k8s.io/csi-translation-lib => github.com/rancher/kubernetes/staging/src/k8s.io/csi-translation-lib v1.16.2-k3s.1
	k8s.io/kube-aggregator => github.com/rancher/kubernetes/staging/src/k8s.io/kube-aggregator v1.16.2-k3s.1
	k8s.io/kube-controller-manager => github.com/rancher/kubernetes/staging/src/k8s.io/kube-controller-manager v1.16.2-k3s.1
	k8s.io/kube-proxy => github.com/rancher/kubernetes/staging/src/k8s.io/kube-proxy v1.16.2-k3s.1
	k8s.io/kube-scheduler => github.com/rancher/kubernetes/staging/src/k8s.io/kube-scheduler v1.16.2-k3s.1
	k8s.io/kubectl => github.com/rancher/kubernetes/staging/src/k8s.io/kubectl v1.16.2-k3s.1
	k8s.io/kubelet => github.com/rancher/kubernetes/staging/src/k8s.io/kubelet v1.16.2-k3s.1
	k8s.io/kubernetes => github.com/rancher/kubernetes v1.16.2-k3s.1
	k8s.io/legacy-cloud-providers => github.com/rancher/kubernetes/staging/src/k8s.io/legacy-cloud-providers v1.16.2-k3s.1
	k8s.io/metrics => github.com/rancher/kubernetes/staging/src/k8s.io/metrics v1.16.2-k3s.1
	k8s.io/node-api => github.com/rancher/kubernetes/staging/src/k8s.io/node-api v1.16.2-k3s.1
	k8s.io/sample-apiserver => github.com/rancher/kubernetes/staging/src/k8s.io/sample-apiserver v1.16.2-k3s.1
	k8s.io/sample-cli-plugin => github.com/rancher/kubernetes/staging/src/k8s.io/sample-cli-plugin v1.16.2-k3s.1
	k8s.io/sample-controller => github.com/rancher/kubernetes/staging/src/k8s.io/sample-controller v1.16.2-k3s.1
)

require (
	contrib.go.opencensus.io/exporter/prometheus v0.1.0 // indirect
	contrib.go.opencensus.io/exporter/stackdriver v0.12.7 // indirect
	github.com/deislabs/smi-sdk-go v0.0.0-20190819154013-e53a9b2d8c1a
	github.com/google/go-containerregistry v0.0.0-20191023194145-7683b4ee5f61 // indirect
	github.com/jetstack/cert-manager v0.11.0
	github.com/markbates/inflect v1.0.4 // indirect
	github.com/rancher/wrangler v0.2.1-0.20191025041946-1fd360590735
	github.com/sirupsen/logrus v1.4.2
	github.com/tektoncd/pipeline v0.8.0
	golang.org/x/xerrors v0.0.0-20191011141410-1b5146add898 // indirect
	k8s.io/api v0.0.0
	k8s.io/apiextensions-apiserver v0.0.0
	k8s.io/apimachinery v0.0.0
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/kube-aggregator v0.0.0
	knative.dev/pkg v0.0.0-20191024223035-2a3fc371d326 // indirect
)
