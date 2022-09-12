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

package fake

import (
	"context"

	v1alpha1 "github.com/kube-all/api/cluster/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeWorkspaceResourceQuotas implements WorkspaceResourceQuotaInterface
type FakeWorkspaceResourceQuotas struct {
	Fake *FakeClusterV1alpha1
}

var workspaceresourcequotasResource = schema.GroupVersionResource{Group: "cluster.kubeall.com", Version: "v1alpha1", Resource: "workspaceresourcequotas"}

var workspaceresourcequotasKind = schema.GroupVersionKind{Group: "cluster.kubeall.com", Version: "v1alpha1", Kind: "WorkspaceResourceQuota"}

// Get takes name of the workspaceResourceQuota, and returns the corresponding workspaceResourceQuota object, and an error if there is any.
func (c *FakeWorkspaceResourceQuotas) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.WorkspaceResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(workspaceresourcequotasResource, name), &v1alpha1.WorkspaceResourceQuota{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkspaceResourceQuota), err
}

// List takes label and field selectors, and returns the list of WorkspaceResourceQuotas that match those selectors.
func (c *FakeWorkspaceResourceQuotas) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.WorkspaceResourceQuotaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(workspaceresourcequotasResource, workspaceresourcequotasKind, opts), &v1alpha1.WorkspaceResourceQuotaList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.WorkspaceResourceQuotaList{ListMeta: obj.(*v1alpha1.WorkspaceResourceQuotaList).ListMeta}
	for _, item := range obj.(*v1alpha1.WorkspaceResourceQuotaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested workspaceResourceQuotas.
func (c *FakeWorkspaceResourceQuotas) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(workspaceresourcequotasResource, opts))
}

// Create takes the representation of a workspaceResourceQuota and creates it.  Returns the server's representation of the workspaceResourceQuota, and an error, if there is any.
func (c *FakeWorkspaceResourceQuotas) Create(ctx context.Context, workspaceResourceQuota *v1alpha1.WorkspaceResourceQuota, opts v1.CreateOptions) (result *v1alpha1.WorkspaceResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(workspaceresourcequotasResource, workspaceResourceQuota), &v1alpha1.WorkspaceResourceQuota{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkspaceResourceQuota), err
}

// Update takes the representation of a workspaceResourceQuota and updates it. Returns the server's representation of the workspaceResourceQuota, and an error, if there is any.
func (c *FakeWorkspaceResourceQuotas) Update(ctx context.Context, workspaceResourceQuota *v1alpha1.WorkspaceResourceQuota, opts v1.UpdateOptions) (result *v1alpha1.WorkspaceResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(workspaceresourcequotasResource, workspaceResourceQuota), &v1alpha1.WorkspaceResourceQuota{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkspaceResourceQuota), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWorkspaceResourceQuotas) UpdateStatus(ctx context.Context, workspaceResourceQuota *v1alpha1.WorkspaceResourceQuota, opts v1.UpdateOptions) (*v1alpha1.WorkspaceResourceQuota, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(workspaceresourcequotasResource, "status", workspaceResourceQuota), &v1alpha1.WorkspaceResourceQuota{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkspaceResourceQuota), err
}

// Delete takes name of the workspaceResourceQuota and deletes it. Returns an error if one occurs.
func (c *FakeWorkspaceResourceQuotas) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(workspaceresourcequotasResource, name, opts), &v1alpha1.WorkspaceResourceQuota{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWorkspaceResourceQuotas) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(workspaceresourcequotasResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.WorkspaceResourceQuotaList{})
	return err
}

// Patch applies the patch and returns the patched workspaceResourceQuota.
func (c *FakeWorkspaceResourceQuotas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WorkspaceResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(workspaceresourcequotasResource, name, pt, data, subresources...), &v1alpha1.WorkspaceResourceQuota{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkspaceResourceQuota), err
}
