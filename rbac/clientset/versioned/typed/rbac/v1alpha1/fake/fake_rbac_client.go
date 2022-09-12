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
	v1alpha1 "github.com/kube-all/client-go/rbac/clientset/versioned/typed/rbac/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeRbacV1alpha1 struct {
	*testing.Fake
}

func (c *FakeRbacV1alpha1) KubeUsers() v1alpha1.KubeUserInterface {
	return &FakeKubeUsers{c}
}

func (c *FakeRbacV1alpha1) KubeUserAPIKeys() v1alpha1.KubeUserAPIKeyInterface {
	return &FakeKubeUserAPIKeys{c}
}

func (c *FakeRbacV1alpha1) UserKubeConfigs() v1alpha1.UserKubeConfigInterface {
	return &FakeUserKubeConfigs{c}
}

func (c *FakeRbacV1alpha1) WorkspaceRoles() v1alpha1.WorkspaceRoleInterface {
	return &FakeWorkspaceRoles{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeRbacV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
