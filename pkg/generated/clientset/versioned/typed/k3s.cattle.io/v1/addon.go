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

package v1

import (
	"context"

	v1 "github.com/k3s-io/k3s/pkg/apis/k3s.cattle.io/v1"
	scheme "github.com/k3s-io/k3s/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// AddonsGetter has a method to return a AddonInterface.
// A group's client should implement this interface.
type AddonsGetter interface {
	Addons(namespace string) AddonInterface
}

// AddonInterface has methods to work with Addon resources.
type AddonInterface interface {
	Create(ctx context.Context, addon *v1.Addon, opts metav1.CreateOptions) (*v1.Addon, error)
	Update(ctx context.Context, addon *v1.Addon, opts metav1.UpdateOptions) (*v1.Addon, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Addon, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.AddonList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Addon, err error)
	AddonExpansion
}

// addons implements AddonInterface
type addons struct {
	*gentype.ClientWithList[*v1.Addon, *v1.AddonList]
}

// newAddons returns a Addons
func newAddons(c *K3sV1Client, namespace string) *addons {
	return &addons{
		gentype.NewClientWithList[*v1.Addon, *v1.AddonList](
			"addons",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.Addon { return &v1.Addon{} },
			func() *v1.AddonList { return &v1.AddonList{} }),
	}
}