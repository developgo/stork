/*
Copyright 2018 Openstorage.org

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
	"time"

	v1alpha1 "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	scheme "github.com/libopenstorage/stork/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SchedulePoliciesGetter has a method to return a SchedulePolicyInterface.
// A group's client should implement this interface.
type SchedulePoliciesGetter interface {
	SchedulePolicies() SchedulePolicyInterface
}

// SchedulePolicyInterface has methods to work with SchedulePolicy resources.
type SchedulePolicyInterface interface {
	Create(*v1alpha1.SchedulePolicy) (*v1alpha1.SchedulePolicy, error)
	Update(*v1alpha1.SchedulePolicy) (*v1alpha1.SchedulePolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SchedulePolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.SchedulePolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SchedulePolicy, err error)
	SchedulePolicyExpansion
}

// schedulePolicies implements SchedulePolicyInterface
type schedulePolicies struct {
	client rest.Interface
}

// newSchedulePolicies returns a SchedulePolicies
func newSchedulePolicies(c *StorkV1alpha1Client) *schedulePolicies {
	return &schedulePolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the schedulePolicy, and returns the corresponding schedulePolicy object, and an error if there is any.
func (c *schedulePolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.SchedulePolicy, err error) {
	result = &v1alpha1.SchedulePolicy{}
	err = c.client.Get().
		Resource("schedulepolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SchedulePolicies that match those selectors.
func (c *schedulePolicies) List(opts v1.ListOptions) (result *v1alpha1.SchedulePolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SchedulePolicyList{}
	err = c.client.Get().
		Resource("schedulepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested schedulePolicies.
func (c *schedulePolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("schedulepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a schedulePolicy and creates it.  Returns the server's representation of the schedulePolicy, and an error, if there is any.
func (c *schedulePolicies) Create(schedulePolicy *v1alpha1.SchedulePolicy) (result *v1alpha1.SchedulePolicy, err error) {
	result = &v1alpha1.SchedulePolicy{}
	err = c.client.Post().
		Resource("schedulepolicies").
		Body(schedulePolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a schedulePolicy and updates it. Returns the server's representation of the schedulePolicy, and an error, if there is any.
func (c *schedulePolicies) Update(schedulePolicy *v1alpha1.SchedulePolicy) (result *v1alpha1.SchedulePolicy, err error) {
	result = &v1alpha1.SchedulePolicy{}
	err = c.client.Put().
		Resource("schedulepolicies").
		Name(schedulePolicy.Name).
		Body(schedulePolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the schedulePolicy and deletes it. Returns an error if one occurs.
func (c *schedulePolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("schedulepolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *schedulePolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("schedulepolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched schedulePolicy.
func (c *schedulePolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SchedulePolicy, err error) {
	result = &v1alpha1.SchedulePolicy{}
	err = c.client.Patch(pt).
		Resource("schedulepolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
