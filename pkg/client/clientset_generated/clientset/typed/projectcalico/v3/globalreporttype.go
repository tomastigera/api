// Copyright (c) 2021 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package v3

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	scheme "github.com/projectcalico/api/pkg/client/clientset_generated/clientset/scheme"
)

// GlobalReportTypesGetter has a method to return a GlobalReportTypeInterface.
// A group's client should implement this interface.
type GlobalReportTypesGetter interface {
	GlobalReportTypes() GlobalReportTypeInterface
}

// GlobalReportTypeInterface has methods to work with GlobalReportType resources.
type GlobalReportTypeInterface interface {
	Create(ctx context.Context, globalReportType *v3.GlobalReportType, opts v1.CreateOptions) (*v3.GlobalReportType, error)
	Update(ctx context.Context, globalReportType *v3.GlobalReportType, opts v1.UpdateOptions) (*v3.GlobalReportType, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.GlobalReportType, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.GlobalReportTypeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.GlobalReportType, err error)
	GlobalReportTypeExpansion
}

// globalReportTypes implements GlobalReportTypeInterface
type globalReportTypes struct {
	client rest.Interface
}

// newGlobalReportTypes returns a GlobalReportTypes
func newGlobalReportTypes(c *ProjectcalicoV3Client) *globalReportTypes {
	return &globalReportTypes{
		client: c.RESTClient(),
	}
}

// Get takes name of the globalReportType, and returns the corresponding globalReportType object, and an error if there is any.
func (c *globalReportTypes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.GlobalReportType, err error) {
	result = &v3.GlobalReportType{}
	err = c.client.Get().
		Resource("globalreporttypes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GlobalReportTypes that match those selectors.
func (c *globalReportTypes) List(ctx context.Context, opts v1.ListOptions) (result *v3.GlobalReportTypeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.GlobalReportTypeList{}
	err = c.client.Get().
		Resource("globalreporttypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested globalReportTypes.
func (c *globalReportTypes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("globalreporttypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a globalReportType and creates it.  Returns the server's representation of the globalReportType, and an error, if there is any.
func (c *globalReportTypes) Create(ctx context.Context, globalReportType *v3.GlobalReportType, opts v1.CreateOptions) (result *v3.GlobalReportType, err error) {
	result = &v3.GlobalReportType{}
	err = c.client.Post().
		Resource("globalreporttypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(globalReportType).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a globalReportType and updates it. Returns the server's representation of the globalReportType, and an error, if there is any.
func (c *globalReportTypes) Update(ctx context.Context, globalReportType *v3.GlobalReportType, opts v1.UpdateOptions) (result *v3.GlobalReportType, err error) {
	result = &v3.GlobalReportType{}
	err = c.client.Put().
		Resource("globalreporttypes").
		Name(globalReportType.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(globalReportType).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the globalReportType and deletes it. Returns an error if one occurs.
func (c *globalReportTypes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("globalreporttypes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *globalReportTypes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("globalreporttypes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched globalReportType.
func (c *globalReportTypes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.GlobalReportType, err error) {
	result = &v3.GlobalReportType{}
	err = c.client.Patch(pt).
		Resource("globalreporttypes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
