/*
Copyright 2019 The OpenEBS Authors

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
	v1alpha1 "github.com/openebs/zfs-localpv/pkg/apis/openebs.io/zfs/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeZFSVolumes implements ZFSVolumeInterface
type FakeZFSVolumes struct {
	Fake *FakeZfsV1alpha1
	ns   string
}

var zfsvolumesResource = schema.GroupVersionResource{Group: "zfs.openebs.io", Version: "v1alpha1", Resource: "zfsvolumes"}

var zfsvolumesKind = schema.GroupVersionKind{Group: "zfs.openebs.io", Version: "v1alpha1", Kind: "ZFSVolume"}

// Get takes name of the zFSVolume, and returns the corresponding zFSVolume object, and an error if there is any.
func (c *FakeZFSVolumes) Get(name string, options v1.GetOptions) (result *v1alpha1.ZFSVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(zfsvolumesResource, c.ns, name), &v1alpha1.ZFSVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ZFSVolume), err
}

// List takes label and field selectors, and returns the list of ZFSVolumes that match those selectors.
func (c *FakeZFSVolumes) List(opts v1.ListOptions) (result *v1alpha1.ZFSVolumeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(zfsvolumesResource, zfsvolumesKind, c.ns, opts), &v1alpha1.ZFSVolumeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ZFSVolumeList{ListMeta: obj.(*v1alpha1.ZFSVolumeList).ListMeta}
	for _, item := range obj.(*v1alpha1.ZFSVolumeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested zFSVolumes.
func (c *FakeZFSVolumes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(zfsvolumesResource, c.ns, opts))

}

// Create takes the representation of a zFSVolume and creates it.  Returns the server's representation of the zFSVolume, and an error, if there is any.
func (c *FakeZFSVolumes) Create(zFSVolume *v1alpha1.ZFSVolume) (result *v1alpha1.ZFSVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(zfsvolumesResource, c.ns, zFSVolume), &v1alpha1.ZFSVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ZFSVolume), err
}

// Update takes the representation of a zFSVolume and updates it. Returns the server's representation of the zFSVolume, and an error, if there is any.
func (c *FakeZFSVolumes) Update(zFSVolume *v1alpha1.ZFSVolume) (result *v1alpha1.ZFSVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(zfsvolumesResource, c.ns, zFSVolume), &v1alpha1.ZFSVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ZFSVolume), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeZFSVolumes) UpdateStatus(zFSVolume *v1alpha1.ZFSVolume) (*v1alpha1.ZFSVolume, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(zfsvolumesResource, "status", c.ns, zFSVolume), &v1alpha1.ZFSVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ZFSVolume), err
}

// Delete takes name of the zFSVolume and deletes it. Returns an error if one occurs.
func (c *FakeZFSVolumes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(zfsvolumesResource, c.ns, name), &v1alpha1.ZFSVolume{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeZFSVolumes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(zfsvolumesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ZFSVolumeList{})
	return err
}

// Patch applies the patch and returns the patched zFSVolume.
func (c *FakeZFSVolumes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ZFSVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(zfsvolumesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ZFSVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ZFSVolume), err
}
