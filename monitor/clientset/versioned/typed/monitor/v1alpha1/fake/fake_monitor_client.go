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
	v1alpha1 "github.com/kube-all/client-go/monitor/clientset/versioned/typed/monitor/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeMonitorV1alpha1 struct {
	*testing.Fake
}

func (c *FakeMonitorV1alpha1) ClusterDashboards() v1alpha1.ClusterDashboardInterface {
	return &FakeClusterDashboards{c}
}

func (c *FakeMonitorV1alpha1) ProjectArchitectures(namespace string) v1alpha1.ProjectArchitectureInterface {
	return &FakeProjectArchitectures{c, namespace}
}

func (c *FakeMonitorV1alpha1) Workloads(namespace string) v1alpha1.WorkloadInterface {
	return &FakeWorkloads{c, namespace}
}

func (c *FakeMonitorV1alpha1) WorkspaceDashboards() v1alpha1.WorkspaceDashboardInterface {
	return &FakeWorkspaceDashboards{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeMonitorV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
