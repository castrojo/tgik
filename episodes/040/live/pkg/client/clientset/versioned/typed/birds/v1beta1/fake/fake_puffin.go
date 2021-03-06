// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/heptio/tgik/episodes/040/live/pkg/apis/birds/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePuffins implements PuffinInterface
type FakePuffins struct {
	Fake *FakeBirdsV1beta1
	ns   string
}

var puffinsResource = schema.GroupVersionResource{Group: "birds.fabulous.af", Version: "v1beta1", Resource: "puffins"}

var puffinsKind = schema.GroupVersionKind{Group: "birds.fabulous.af", Version: "v1beta1", Kind: "Puffin"}

// Get takes name of the puffin, and returns the corresponding puffin object, and an error if there is any.
func (c *FakePuffins) Get(name string, options v1.GetOptions) (result *v1beta1.Puffin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(puffinsResource, c.ns, name), &v1beta1.Puffin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Puffin), err
}

// List takes label and field selectors, and returns the list of Puffins that match those selectors.
func (c *FakePuffins) List(opts v1.ListOptions) (result *v1beta1.PuffinList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(puffinsResource, puffinsKind, c.ns, opts), &v1beta1.PuffinList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.PuffinList{}
	for _, item := range obj.(*v1beta1.PuffinList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested puffins.
func (c *FakePuffins) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(puffinsResource, c.ns, opts))

}

// Create takes the representation of a puffin and creates it.  Returns the server's representation of the puffin, and an error, if there is any.
func (c *FakePuffins) Create(puffin *v1beta1.Puffin) (result *v1beta1.Puffin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(puffinsResource, c.ns, puffin), &v1beta1.Puffin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Puffin), err
}

// Update takes the representation of a puffin and updates it. Returns the server's representation of the puffin, and an error, if there is any.
func (c *FakePuffins) Update(puffin *v1beta1.Puffin) (result *v1beta1.Puffin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(puffinsResource, c.ns, puffin), &v1beta1.Puffin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Puffin), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePuffins) UpdateStatus(puffin *v1beta1.Puffin) (*v1beta1.Puffin, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(puffinsResource, "status", c.ns, puffin), &v1beta1.Puffin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Puffin), err
}

// Delete takes name of the puffin and deletes it. Returns an error if one occurs.
func (c *FakePuffins) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(puffinsResource, c.ns, name), &v1beta1.Puffin{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePuffins) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(puffinsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.PuffinList{})
	return err
}

// Patch applies the patch and returns the patched puffin.
func (c *FakePuffins) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Puffin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(puffinsResource, c.ns, name, data, subresources...), &v1beta1.Puffin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Puffin), err
}
