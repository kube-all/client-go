/*
Copyright 2022 The kubeall.com Authors.

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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/kube-all/api/cluster/v1alpha1"
	scheme "github.com/kube-all/client-go/cluster/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// WorkspaceResourceQuotasGetter has a method to return a WorkspaceResourceQuotaInterface.
// A group's client should implement this interface.
type WorkspaceResourceQuotasGetter interface {
	WorkspaceResourceQuotas() WorkspaceResourceQuotaInterface
}

// WorkspaceResourceQuotaInterface has methods to work with WorkspaceResourceQuota resources.
type WorkspaceResourceQuotaInterface interface {
	Create(ctx context.Context, workspaceResourceQuota *v1alpha1.WorkspaceResourceQuota, opts v1.CreateOptions) (*v1alpha1.WorkspaceResourceQuota, error)
	Update(ctx context.Context, workspaceResourceQuota *v1alpha1.WorkspaceResourceQuota, opts v1.UpdateOptions) (*v1alpha1.WorkspaceResourceQuota, error)
	UpdateStatus(ctx context.Context, workspaceResourceQuota *v1alpha1.WorkspaceResourceQuota, opts v1.UpdateOptions) (*v1alpha1.WorkspaceResourceQuota, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.WorkspaceResourceQuota, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.WorkspaceResourceQuotaList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WorkspaceResourceQuota, err error)
	WorkspaceResourceQuotaExpansion
}

// workspaceResourceQuotas implements WorkspaceResourceQuotaInterface
type workspaceResourceQuotas struct {
	client rest.Interface
}

// newWorkspaceResourceQuotas returns a WorkspaceResourceQuotas
func newWorkspaceResourceQuotas(c *ClusterV1alpha1Client) *workspaceResourceQuotas {
	return &workspaceResourceQuotas{
		client: c.RESTClient(),
	}
}

// Get takes name of the workspaceResourceQuota, and returns the corresponding workspaceResourceQuota object, and an error if there is any.
func (c *workspaceResourceQuotas) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.WorkspaceResourceQuota, err error) {
	result = &v1alpha1.WorkspaceResourceQuota{}
	err = c.client.Get().
		Resource("workspaceresourcequotas").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WorkspaceResourceQuotas that match those selectors.
func (c *workspaceResourceQuotas) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.WorkspaceResourceQuotaList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.WorkspaceResourceQuotaList{}
	err = c.client.Get().
		Resource("workspaceresourcequotas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested workspaceResourceQuotas.
func (c *workspaceResourceQuotas) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("workspaceresourcequotas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a workspaceResourceQuota and creates it.  Returns the server's representation of the workspaceResourceQuota, and an error, if there is any.
func (c *workspaceResourceQuotas) Create(ctx context.Context, workspaceResourceQuota *v1alpha1.WorkspaceResourceQuota, opts v1.CreateOptions) (result *v1alpha1.WorkspaceResourceQuota, err error) {
	result = &v1alpha1.WorkspaceResourceQuota{}
	err = c.client.Post().
		Resource("workspaceresourcequotas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(workspaceResourceQuota).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a workspaceResourceQuota and updates it. Returns the server's representation of the workspaceResourceQuota, and an error, if there is any.
func (c *workspaceResourceQuotas) Update(ctx context.Context, workspaceResourceQuota *v1alpha1.WorkspaceResourceQuota, opts v1.UpdateOptions) (result *v1alpha1.WorkspaceResourceQuota, err error) {
	result = &v1alpha1.WorkspaceResourceQuota{}
	err = c.client.Put().
		Resource("workspaceresourcequotas").
		Name(workspaceResourceQuota.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(workspaceResourceQuota).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *workspaceResourceQuotas) UpdateStatus(ctx context.Context, workspaceResourceQuota *v1alpha1.WorkspaceResourceQuota, opts v1.UpdateOptions) (result *v1alpha1.WorkspaceResourceQuota, err error) {
	result = &v1alpha1.WorkspaceResourceQuota{}
	err = c.client.Put().
		Resource("workspaceresourcequotas").
		Name(workspaceResourceQuota.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(workspaceResourceQuota).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the workspaceResourceQuota and deletes it. Returns an error if one occurs.
func (c *workspaceResourceQuotas) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("workspaceresourcequotas").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *workspaceResourceQuotas) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("workspaceresourcequotas").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched workspaceResourceQuota.
func (c *workspaceResourceQuotas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WorkspaceResourceQuota, err error) {
	result = &v1alpha1.WorkspaceResourceQuota{}
	err = c.client.Patch(pt).
		Resource("workspaceresourcequotas").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
