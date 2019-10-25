package main

import (
	splitv1alpha1 "github.com/deislabs/smi-sdk-go/pkg/apis/split/v1alpha1"
	certmanagerv1alpha2 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha2"
	controllergen "github.com/rancher/wrangler/pkg/controller-gen"
	"github.com/rancher/wrangler/pkg/controller-gen/args"
	pipelinev1alpha1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
	storagev1 "k8s.io/api/storage/v1"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiv1 "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"
)

func main() {
	controllergen.Run(args.Options{
		OutputPackage: "github.com/rancher/wrangler-api/pkg/generated",
		Boilerplate:   "scripts/boilerplate.go.txt",
		Groups: map[string]args.Group{
			"": {
				Types: []interface{}{
					v1.Node{},
					v1.Namespace{},
					v1.Secret{},
					v1.Service{},
					v1.ServiceAccount{},
					v1.Endpoints{},
					v1.ConfigMap{},
					v1.PersistentVolumeClaim{},
					v1.Pod{},
				},
				InformersPackage: "k8s.io/client-go/informers",
				ClientSetPackage: "k8s.io/client-go/kubernetes",
				ListersPackage:   "k8s.io/client-go/listers",
			},
			"extensions": {
				Types: []interface{}{
					extensionsv1beta1.Ingress{},
				},
				InformersPackage: "k8s.io/client-go/informers",
				ClientSetPackage: "k8s.io/client-go/kubernetes",
				ListersPackage:   "k8s.io/client-go/listers",
			},
			"rbac": {
				Types: []interface{}{
					rbacv1.Role{},
					rbacv1.RoleBinding{},
					rbacv1.ClusterRole{},
					rbacv1.ClusterRoleBinding{},
				},
				InformersPackage: "k8s.io/client-go/informers",
				ClientSetPackage: "k8s.io/client-go/kubernetes",
				ListersPackage:   "k8s.io/client-go/listers",
			},
			"apps": {
				Types: []interface{}{
					appsv1.Deployment{},
					appsv1.DaemonSet{},
					appsv1.StatefulSet{},
				},
				InformersPackage: "k8s.io/client-go/informers",
				ClientSetPackage: "k8s.io/client-go/kubernetes",
				ListersPackage:   "k8s.io/client-go/listers",
			},
			"storage": {
				Types: []interface{}{
					storagev1.StorageClass{},
				},
				InformersPackage: "k8s.io/client-go/informers",
				ClientSetPackage: "k8s.io/client-go/kubernetes",
				ListersPackage:   "k8s.io/client-go/listers",
			},
			"cert-manager.io": {
				Types: []interface{}{
					certmanagerv1alpha2.Certificate{},
					certmanagerv1alpha2.ClusterIssuer{},
					certmanagerv1alpha2.Issuer{},
				},
				PackageName:     "certmanager",
				GenerateClients: true,
			},
			"apiextensions.k8s.io": {
				Types: []interface{}{
					v1beta1.CustomResourceDefinition{},
				},
				PackageName:      "apiextensions",
				ClientSetPackage: "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset",
				InformersPackage: "k8s.io/apiextensions-apiserver/pkg/client/informers/externalversions",
				ListersPackage:   "k8s.io/apiextensions-apiserver/pkg/client/listers",
			},
			"apiregistration.k8s.io": {
				Types: []interface{}{
					apiv1.APIService{},
				},
				PackageName:      "apiregistration",
				ClientSetPackage: "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset",
				InformersPackage: "k8s.io/kube-aggregator/pkg/client/informers/externalversions",
				ListersPackage:   "k8s.io/kube-aggregator/pkg/client/listers",
			},
			"batch": {
				Types: []interface{}{
					batchv1.Job{},
				},
				InformersPackage: "k8s.io/client-go/informers",
				ClientSetPackage: "k8s.io/client-go/kubernetes",
				ListersPackage:   "k8s.io/client-go/listers",
			},
			"tekton.dev": {
				Types: []interface{}{
					pipelinev1alpha1.TaskRun{},
				},
				PackageName:     "pipeline",
				GenerateClients: true,
			},
			"split.smi-spec.io": {
				Types: []interface{}{
					splitv1alpha1.TrafficSplit{},
				},
				PackageName:     "split",
				GenerateClients: true,
			},
		},
	})
}
