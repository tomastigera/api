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

// GlobalReportsGetter has a method to return a GlobalReportInterface.
// A group's client should implement this interface.
type GlobalReportsGetter interface {
	GlobalReports() GlobalReportInterface
}

// GlobalReportInterface has methods to work with GlobalReport resources.
type GlobalReportInterface interface {
	Create(ctx context.Context, globalReport *v3.GlobalReport, opts v1.CreateOptions) (*v3.GlobalReport, error)
	Update(ctx context.Context, globalReport *v3.GlobalReport, opts v1.UpdateOptions) (*v3.GlobalReport, error)
	UpdateStatus(ctx context.Context, globalReport *v3.GlobalReport, opts v1.UpdateOptions) (*v3.GlobalReport, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.GlobalReport, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.GlobalReportList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.GlobalReport, err error)
	GlobalReportExpansion
}

// globalReports implements GlobalReportInterface
type globalReports struct {
	client rest.Interface
}

// newGlobalReports returns a GlobalReports
func newGlobalReports(c *ProjectcalicoV3Client) *globalReports {
	return &globalReports{
		client: c.RESTClient(),
	}
}

// Get takes name of the globalReport, and returns the corresponding globalReport object, and an error if there is any.
func (c *globalReports) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.GlobalReport, err error) {
	result = &v3.GlobalReport{}
	err = c.client.Get().
		Resource("globalreports").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GlobalReports that match those selectors.
func (c *globalReports) List(ctx context.Context, opts v1.ListOptions) (result *v3.GlobalReportList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.GlobalReportList{}
	err = c.client.Get().
		Resource("globalreports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested globalReports.
func (c *globalReports) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("globalreports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a globalReport and creates it.  Returns the server's representation of the globalReport, and an error, if there is any.
func (c *globalReports) Create(ctx context.Context, globalReport *v3.GlobalReport, opts v1.CreateOptions) (result *v3.GlobalReport, err error) {
	result = &v3.GlobalReport{}
	err = c.client.Post().
		Resource("globalreports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(globalReport).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a globalReport and updates it. Returns the server's representation of the globalReport, and an error, if there is any.
func (c *globalReports) Update(ctx context.Context, globalReport *v3.GlobalReport, opts v1.UpdateOptions) (result *v3.GlobalReport, err error) {
	result = &v3.GlobalReport{}
	err = c.client.Put().
		Resource("globalreports").
		Name(globalReport.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(globalReport).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *globalReports) UpdateStatus(ctx context.Context, globalReport *v3.GlobalReport, opts v1.UpdateOptions) (result *v3.GlobalReport, err error) {
	result = &v3.GlobalReport{}
	err = c.client.Put().
		Resource("globalreports").
		Name(globalReport.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(globalReport).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the globalReport and deletes it. Returns an error if one occurs.
func (c *globalReports) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("globalreports").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *globalReports) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("globalreports").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched globalReport.
func (c *globalReports) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.GlobalReport, err error) {
	result = &v3.GlobalReport{}
	err = c.client.Patch(pt).
		Resource("globalreports").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
