/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/apply"
	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/kv"
	v1beta1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type IngressHandler func(string, *v1beta1.Ingress) (*v1beta1.Ingress, error)

type IngressController interface {
	generic.ControllerMeta
	IngressClient

	OnChange(ctx context.Context, name string, sync IngressHandler)
	OnRemove(ctx context.Context, name string, sync IngressHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() IngressCache
}

type IngressClient interface {
	Create(*v1beta1.Ingress) (*v1beta1.Ingress, error)
	Update(*v1beta1.Ingress) (*v1beta1.Ingress, error)
	UpdateStatus(*v1beta1.Ingress) (*v1beta1.Ingress, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1beta1.Ingress, error)
	List(namespace string, opts metav1.ListOptions) (*v1beta1.IngressList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Ingress, err error)
}

type IngressCache interface {
	Get(namespace, name string) (*v1beta1.Ingress, error)
	List(namespace string, selector labels.Selector) ([]*v1beta1.Ingress, error)

	AddIndexer(indexName string, indexer IngressIndexer)
	GetByIndex(indexName, key string) ([]*v1beta1.Ingress, error)
}

type IngressIndexer func(obj *v1beta1.Ingress) ([]string, error)

type ingressController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewIngressController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) IngressController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &ingressController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromIngressHandlerToHandler(sync IngressHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1beta1.Ingress
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1beta1.Ingress))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *ingressController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1beta1.Ingress))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateIngressDeepCopyOnChange(client IngressClient, obj *v1beta1.Ingress, handler func(obj *v1beta1.Ingress) (*v1beta1.Ingress, error)) (*v1beta1.Ingress, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *ingressController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *ingressController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *ingressController) OnChange(ctx context.Context, name string, sync IngressHandler) {
	c.AddGenericHandler(ctx, name, FromIngressHandlerToHandler(sync))
}

func (c *ingressController) OnRemove(ctx context.Context, name string, sync IngressHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromIngressHandlerToHandler(sync)))
}

func (c *ingressController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *ingressController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *ingressController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *ingressController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *ingressController) Cache() IngressCache {
	return &ingressCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *ingressController) Create(obj *v1beta1.Ingress) (*v1beta1.Ingress, error) {
	result := &v1beta1.Ingress{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *ingressController) Update(obj *v1beta1.Ingress) (*v1beta1.Ingress, error) {
	result := &v1beta1.Ingress{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *ingressController) UpdateStatus(obj *v1beta1.Ingress) (*v1beta1.Ingress, error) {
	result := &v1beta1.Ingress{}
	return result, c.client.UpdateStatus(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *ingressController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *ingressController) Get(namespace, name string, options metav1.GetOptions) (*v1beta1.Ingress, error) {
	result := &v1beta1.Ingress{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *ingressController) List(namespace string, opts metav1.ListOptions) (*v1beta1.IngressList, error) {
	result := &v1beta1.IngressList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *ingressController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *ingressController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1beta1.Ingress, error) {
	result := &v1beta1.Ingress{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type ingressCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *ingressCache) Get(namespace, name string) (*v1beta1.Ingress, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1beta1.Ingress), nil
}

func (c *ingressCache) List(namespace string, selector labels.Selector) (ret []*v1beta1.Ingress, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Ingress))
	})

	return ret, err
}

func (c *ingressCache) AddIndexer(indexName string, indexer IngressIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1beta1.Ingress))
		},
	}))
}

func (c *ingressCache) GetByIndex(indexName, key string) (result []*v1beta1.Ingress, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1beta1.Ingress, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1beta1.Ingress))
	}
	return result, nil
}

type IngressStatusHandler func(obj *v1beta1.Ingress, status v1beta1.IngressStatus) (v1beta1.IngressStatus, error)

type IngressGeneratingHandler func(obj *v1beta1.Ingress, status v1beta1.IngressStatus) ([]runtime.Object, v1beta1.IngressStatus, error)

func RegisterIngressStatusHandler(ctx context.Context, controller IngressController, condition condition.Cond, name string, handler IngressStatusHandler) {
	statusHandler := &ingressStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromIngressHandlerToHandler(statusHandler.sync))
}

func RegisterIngressGeneratingHandler(ctx context.Context, controller IngressController, apply apply.Apply,
	condition condition.Cond, name string, handler IngressGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &ingressGeneratingHandler{
		IngressGeneratingHandler: handler,
		apply:                    apply,
		name:                     name,
		gvk:                      controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterIngressStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type ingressStatusHandler struct {
	client    IngressClient
	condition condition.Cond
	handler   IngressStatusHandler
}

func (a *ingressStatusHandler) sync(key string, obj *v1beta1.Ingress) (*v1beta1.Ingress, error) {
	if obj == nil {
		return obj, nil
	}

	origStatus := obj.Status.DeepCopy()
	obj = obj.DeepCopy()
	newStatus, err := a.handler(obj, obj.Status)
	if err != nil {
		// Revert to old status on error
		newStatus = *origStatus.DeepCopy()
	}

	if a.condition != "" {
		if errors.IsConflict(err) {
			a.condition.SetError(&newStatus, "", nil)
		} else {
			a.condition.SetError(&newStatus, "", err)
		}
	}
	if !equality.Semantic.DeepEqual(origStatus, &newStatus) {
		if a.condition != "" {
			// Since status has changed, update the lastUpdatedTime
			a.condition.LastUpdated(&newStatus, time.Now().UTC().Format(time.RFC3339))
		}

		var newErr error
		obj.Status = newStatus
		newObj, newErr := a.client.UpdateStatus(obj)
		if err == nil {
			err = newErr
		}
		if newErr == nil {
			obj = newObj
		}
	}
	return obj, err
}

type ingressGeneratingHandler struct {
	IngressGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *ingressGeneratingHandler) Remove(key string, obj *v1beta1.Ingress) (*v1beta1.Ingress, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v1beta1.Ingress{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *ingressGeneratingHandler) Handle(obj *v1beta1.Ingress, status v1beta1.IngressStatus) (v1beta1.IngressStatus, error) {
	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.IngressGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
