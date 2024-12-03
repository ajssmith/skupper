/*
Copyright 2021 The Skupper Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v2alpha1 "github.com/skupperproject/skupper/pkg/apis/skupper/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRouterAccesses implements RouterAccessInterface
type FakeRouterAccesses struct {
	Fake *FakeSkupperV2alpha1
	ns   string
}

var routeraccessesResource = v2alpha1.SchemeGroupVersion.WithResource("routeraccesses")

var routeraccessesKind = v2alpha1.SchemeGroupVersion.WithKind("RouterAccess")

// Get takes name of the routerAccess, and returns the corresponding routerAccess object, and an error if there is any.
func (c *FakeRouterAccesses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.RouterAccess, err error) {
	emptyResult := &v2alpha1.RouterAccess{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(routeraccessesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v2alpha1.RouterAccess), err
}

// List takes label and field selectors, and returns the list of RouterAccesses that match those selectors.
func (c *FakeRouterAccesses) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.RouterAccessList, err error) {
	emptyResult := &v2alpha1.RouterAccessList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(routeraccessesResource, routeraccessesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2alpha1.RouterAccessList{ListMeta: obj.(*v2alpha1.RouterAccessList).ListMeta}
	for _, item := range obj.(*v2alpha1.RouterAccessList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested routerAccesses.
func (c *FakeRouterAccesses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(routeraccessesResource, c.ns, opts))

}

// Create takes the representation of a routerAccess and creates it.  Returns the server's representation of the routerAccess, and an error, if there is any.
func (c *FakeRouterAccesses) Create(ctx context.Context, routerAccess *v2alpha1.RouterAccess, opts v1.CreateOptions) (result *v2alpha1.RouterAccess, err error) {
	emptyResult := &v2alpha1.RouterAccess{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(routeraccessesResource, c.ns, routerAccess, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v2alpha1.RouterAccess), err
}

// Update takes the representation of a routerAccess and updates it. Returns the server's representation of the routerAccess, and an error, if there is any.
func (c *FakeRouterAccesses) Update(ctx context.Context, routerAccess *v2alpha1.RouterAccess, opts v1.UpdateOptions) (result *v2alpha1.RouterAccess, err error) {
	emptyResult := &v2alpha1.RouterAccess{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(routeraccessesResource, c.ns, routerAccess, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v2alpha1.RouterAccess), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRouterAccesses) UpdateStatus(ctx context.Context, routerAccess *v2alpha1.RouterAccess, opts v1.UpdateOptions) (result *v2alpha1.RouterAccess, err error) {
	emptyResult := &v2alpha1.RouterAccess{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(routeraccessesResource, "status", c.ns, routerAccess, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v2alpha1.RouterAccess), err
}

// Delete takes name of the routerAccess and deletes it. Returns an error if one occurs.
func (c *FakeRouterAccesses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(routeraccessesResource, c.ns, name, opts), &v2alpha1.RouterAccess{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRouterAccesses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(routeraccessesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v2alpha1.RouterAccessList{})
	return err
}

// Patch applies the patch and returns the patched routerAccess.
func (c *FakeRouterAccesses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.RouterAccess, err error) {
	emptyResult := &v2alpha1.RouterAccess{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(routeraccessesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v2alpha1.RouterAccess), err
}
