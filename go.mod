module github.com/rancher/wrangler-api

go 1.12

replace (
	github.com/knative/pkg => github.com/rancher/pkg v0.0.0-20190514055449-b30ab9de040e
	github.com/matryer/moq => github.com/rancher/moq v0.0.0-20190404221404-ee5226d43009
)

require (
	github.com/deislabs/smi-sdk-go v0.0.0-20190819154013-e53a9b2d8c1a
	github.com/google/go-containerregistry v0.0.0-20190617215043-876b8855d23c // indirect
	github.com/jetstack/cert-manager v0.7.2
	github.com/knative/build v0.6.0
	github.com/knative/pkg v0.0.0-20190514205332-5e4512dcb2ca
	github.com/knative/serving v0.6.1
	github.com/mattbaird/jsonpatch v0.0.0-20171005235357-81af80346b1a // indirect
	github.com/rancher/wrangler v0.2.1-0.20191015042916-f2a6ecca4f20
	github.com/sirupsen/logrus v1.4.1
	github.com/tektoncd/pipeline v0.4.0
	k8s.io/api v0.0.0-20190918195907-bd6ac527cfd2
	k8s.io/apiextensions-apiserver v0.0.0-20190918201827-3de75813f604
	k8s.io/apimachinery v0.0.0-20190817020851-f2f3a405f61d
	k8s.io/client-go v0.0.0-20190918200256-06eb1244587a
	k8s.io/kube-aggregator v0.0.0-20190918201136-c3a845f1fbb2
)
