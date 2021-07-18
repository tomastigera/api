// Copyright (c) 2021 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
)

// FakeGlobalAlertTemplates implements GlobalAlertTemplateInterface
type FakeGlobalAlertTemplates struct {
	Fake *FakeProjectcalicoV3
}

var globalalerttemplatesResource = schema.GroupVersionResource{Group: "projectcalico.org", Version: "v3", Resource: "globalalerttemplates"}

var globalalerttemplatesKind = schema.GroupVersionKind{Group: "projectcalico.org", Version: "v3", Kind: "GlobalAlertTemplate"}

// Get takes name of the globalAlertTemplate, and returns the corresponding globalAlertTemplate object, and an error if there is any.
func (c *FakeGlobalAlertTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.GlobalAlertTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(globalalerttemplatesResource, name), &v3.GlobalAlertTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.GlobalAlertTemplate), err
}

// List takes label and field selectors, and returns the list of GlobalAlertTemplates that match those selectors.
func (c *FakeGlobalAlertTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v3.GlobalAlertTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(globalalerttemplatesResource, globalalerttemplatesKind, opts), &v3.GlobalAlertTemplateList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.GlobalAlertTemplateList{ListMeta: obj.(*v3.GlobalAlertTemplateList).ListMeta}
	for _, item := range obj.(*v3.GlobalAlertTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested globalAlertTemplates.
func (c *FakeGlobalAlertTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(globalalerttemplatesResource, opts))
}

// Create takes the representation of a globalAlertTemplate and creates it.  Returns the server's representation of the globalAlertTemplate, and an error, if there is any.
func (c *FakeGlobalAlertTemplates) Create(ctx context.Context, globalAlertTemplate *v3.GlobalAlertTemplate, opts v1.CreateOptions) (result *v3.GlobalAlertTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(globalalerttemplatesResource, globalAlertTemplate), &v3.GlobalAlertTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.GlobalAlertTemplate), err
}

// Update takes the representation of a globalAlertTemplate and updates it. Returns the server's representation of the globalAlertTemplate, and an error, if there is any.
func (c *FakeGlobalAlertTemplates) Update(ctx context.Context, globalAlertTemplate *v3.GlobalAlertTemplate, opts v1.UpdateOptions) (result *v3.GlobalAlertTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(globalalerttemplatesResource, globalAlertTemplate), &v3.GlobalAlertTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.GlobalAlertTemplate), err
}

// Delete takes name of the globalAlertTemplate and deletes it. Returns an error if one occurs.
func (c *FakeGlobalAlertTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(globalalerttemplatesResource, name), &v3.GlobalAlertTemplate{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGlobalAlertTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(globalalerttemplatesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v3.GlobalAlertTemplateList{})
	return err
}

// Patch applies the patch and returns the patched globalAlertTemplate.
func (c *FakeGlobalAlertTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.GlobalAlertTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(globalalerttemplatesResource, name, pt, data, subresources...), &v3.GlobalAlertTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.GlobalAlertTemplate), err
}
