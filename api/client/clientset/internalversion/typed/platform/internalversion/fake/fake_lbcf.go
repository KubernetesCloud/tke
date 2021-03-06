/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	platform "tkestack.io/tke/api/platform"
)

// FakeLBCFs implements LBCFInterface
type FakeLBCFs struct {
	Fake *FakePlatform
}

var lbcfsResource = schema.GroupVersionResource{Group: "platform.tkestack.io", Version: "", Resource: "lbcfs"}

var lbcfsKind = schema.GroupVersionKind{Group: "platform.tkestack.io", Version: "", Kind: "LBCF"}

// Get takes name of the lBCF, and returns the corresponding lBCF object, and an error if there is any.
func (c *FakeLBCFs) Get(ctx context.Context, name string, options v1.GetOptions) (result *platform.LBCF, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(lbcfsResource, name), &platform.LBCF{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platform.LBCF), err
}

// List takes label and field selectors, and returns the list of LBCFs that match those selectors.
func (c *FakeLBCFs) List(ctx context.Context, opts v1.ListOptions) (result *platform.LBCFList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(lbcfsResource, lbcfsKind, opts), &platform.LBCFList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &platform.LBCFList{ListMeta: obj.(*platform.LBCFList).ListMeta}
	for _, item := range obj.(*platform.LBCFList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lBCFs.
func (c *FakeLBCFs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(lbcfsResource, opts))
}

// Create takes the representation of a lBCF and creates it.  Returns the server's representation of the lBCF, and an error, if there is any.
func (c *FakeLBCFs) Create(ctx context.Context, lBCF *platform.LBCF, opts v1.CreateOptions) (result *platform.LBCF, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(lbcfsResource, lBCF), &platform.LBCF{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platform.LBCF), err
}

// Update takes the representation of a lBCF and updates it. Returns the server's representation of the lBCF, and an error, if there is any.
func (c *FakeLBCFs) Update(ctx context.Context, lBCF *platform.LBCF, opts v1.UpdateOptions) (result *platform.LBCF, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(lbcfsResource, lBCF), &platform.LBCF{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platform.LBCF), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLBCFs) UpdateStatus(ctx context.Context, lBCF *platform.LBCF, opts v1.UpdateOptions) (*platform.LBCF, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(lbcfsResource, "status", lBCF), &platform.LBCF{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platform.LBCF), err
}

// Delete takes name of the lBCF and deletes it. Returns an error if one occurs.
func (c *FakeLBCFs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(lbcfsResource, name), &platform.LBCF{})
	return err
}

// Patch applies the patch and returns the patched lBCF.
func (c *FakeLBCFs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *platform.LBCF, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(lbcfsResource, name, pt, data, subresources...), &platform.LBCF{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platform.LBCF), err
}
