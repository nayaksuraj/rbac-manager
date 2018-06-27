// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/reactiveops/rbac-manager/pkg/apis/rbacmanager/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRBACDefinitions implements RBACDefinitionInterface
type FakeRBACDefinitions struct {
	Fake *FakeRbacmanagerV1beta1
}

var rbacdefinitionsResource = schema.GroupVersionResource{Group: "rbacmanager.reactiveops.io", Version: "v1beta1", Resource: "rbacdefinitions"}

var rbacdefinitionsKind = schema.GroupVersionKind{Group: "rbacmanager.reactiveops.io", Version: "v1beta1", Kind: "RBACDefinition"}

// Get takes name of the rBACDefinition, and returns the corresponding rBACDefinition object, and an error if there is any.
func (c *FakeRBACDefinitions) Get(name string, options v1.GetOptions) (result *v1beta1.RBACDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(rbacdefinitionsResource, name), &v1beta1.RBACDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RBACDefinition), err
}

// List takes label and field selectors, and returns the list of RBACDefinitions that match those selectors.
func (c *FakeRBACDefinitions) List(opts v1.ListOptions) (result *v1beta1.RBACDefinitionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(rbacdefinitionsResource, rbacdefinitionsKind, opts), &v1beta1.RBACDefinitionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.RBACDefinitionList{}
	for _, item := range obj.(*v1beta1.RBACDefinitionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rBACDefinitions.
func (c *FakeRBACDefinitions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(rbacdefinitionsResource, opts))
}

// Create takes the representation of a rBACDefinition and creates it.  Returns the server's representation of the rBACDefinition, and an error, if there is any.
func (c *FakeRBACDefinitions) Create(rBACDefinition *v1beta1.RBACDefinition) (result *v1beta1.RBACDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(rbacdefinitionsResource, rBACDefinition), &v1beta1.RBACDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RBACDefinition), err
}

// Update takes the representation of a rBACDefinition and updates it. Returns the server's representation of the rBACDefinition, and an error, if there is any.
func (c *FakeRBACDefinitions) Update(rBACDefinition *v1beta1.RBACDefinition) (result *v1beta1.RBACDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(rbacdefinitionsResource, rBACDefinition), &v1beta1.RBACDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RBACDefinition), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRBACDefinitions) UpdateStatus(rBACDefinition *v1beta1.RBACDefinition) (*v1beta1.RBACDefinition, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(rbacdefinitionsResource, "status", rBACDefinition), &v1beta1.RBACDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RBACDefinition), err
}

// Delete takes name of the rBACDefinition and deletes it. Returns an error if one occurs.
func (c *FakeRBACDefinitions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(rbacdefinitionsResource, name), &v1beta1.RBACDefinition{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRBACDefinitions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(rbacdefinitionsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.RBACDefinitionList{})
	return err
}

// Patch applies the patch and returns the patched rBACDefinition.
func (c *FakeRBACDefinitions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.RBACDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(rbacdefinitionsResource, name, data, subresources...), &v1beta1.RBACDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RBACDefinition), err
}
