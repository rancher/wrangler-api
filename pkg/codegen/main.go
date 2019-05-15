package main

import (
	"github.com/rancher/wrangler-api/pkg/apis/some.api.group/v1"
	"github.com/rancher/wrangler/pkg/controller-gen"
	"github.com/rancher/wrangler/pkg/controller-gen/args"
)

func main() {
	controllergen.Run(args.Options{
		OutputPackage: "github.com/rancher/wrangler-api/pkg/generated",
		Boilerplate:   "hack/boilerplate.go.txt",
		Groups: map[string]args.Group{
			"some.api.group": {
				Types: []interface{}{
					v1.Foo{},
				},
				GenerateTypes: true,
			},
		},
	})
}
