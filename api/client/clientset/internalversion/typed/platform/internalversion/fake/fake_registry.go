/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	platform "tkestack.io/tke/api/platform"
)

// FakeRegistries implements RegistryInterface
type FakeRegistries struct {
	Fake *FakePlatform
}

var registriesResource = schema.GroupVersionResource{Group: "platform.tkestack.io", Version: "", Resource: "registries"}

var registriesKind = schema.GroupVersionKind{Group: "platform.tkestack.io", Version: "", Kind: "Registry"}

// Get takes name of the registry, and returns the corresponding registry object, and an error if there is any.
func (c *FakeRegistries) Get(name string, options v1.GetOptions) (result *platform.Registry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(registriesResource, name), &platform.Registry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platform.Registry), err
}

// List takes label and field selectors, and returns the list of Registries that match those selectors.
func (c *FakeRegistries) List(opts v1.ListOptions) (result *platform.RegistryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(registriesResource, registriesKind, opts), &platform.RegistryList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &platform.RegistryList{ListMeta: obj.(*platform.RegistryList).ListMeta}
	for _, item := range obj.(*platform.RegistryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested registries.
func (c *FakeRegistries) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(registriesResource, opts))
}

// Create takes the representation of a registry and creates it.  Returns the server's representation of the registry, and an error, if there is any.
func (c *FakeRegistries) Create(registry *platform.Registry) (result *platform.Registry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(registriesResource, registry), &platform.Registry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platform.Registry), err
}

// Update takes the representation of a registry and updates it. Returns the server's representation of the registry, and an error, if there is any.
func (c *FakeRegistries) Update(registry *platform.Registry) (result *platform.Registry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(registriesResource, registry), &platform.Registry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platform.Registry), err
}

// Delete takes name of the registry and deletes it. Returns an error if one occurs.
func (c *FakeRegistries) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(registriesResource, name), &platform.Registry{})
	return err
}

// Patch applies the patch and returns the patched registry.
func (c *FakeRegistries) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *platform.Registry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(registriesResource, name, pt, data, subresources...), &platform.Registry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platform.Registry), err
}