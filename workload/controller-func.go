/*
Copyright 2021 The Gridsum Authors.

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

package workload

import (
	"errors"
	corev1 "k8s.io/api/core/v1"
	"reflect"
)

// loadPodPrefix distinguish prefixes are k 8 s resources of pods.
func (c *controller) loadPodPrefix(namespace string) (interface{}, bool) {
	return c.cachesMap.Load(PodSpacePrefix(namespace))
}

// GetPod return get the specified pod resource based on the namespace and pod name.
func (c *controller) GetPod(namespace, name string) *corev1.Pod {
	if list, ok := c.loadPodPrefix(namespace); ok {
		for _, pod := range list.([]*corev1.Pod) {
			if pod.Name == name {
				return pod
			}
		}
	}
	return nil
}

// GetPodByNameSpace return get all pods under this namespace.
func (c *controller) GetPodByNameSpace(namespace string) ([]*corev1.Pod, error) {
	if list, ok := c.loadPodPrefix(namespace); ok {
		return list.([]*corev1.Pod), nil
	}
	return nil, errors.New("pod not found")
}

// GetPodByLabel return get exact matching pods based on namespace and label.
func (c *controller) GetPodByLabel(namespace string, labels map[string]string) ([]*corev1.Pod, error) {
	ret := make([]*corev1.Pod, 0)
	if list, ok := c.loadPodPrefix(namespace); ok {
		for _, pod := range list.([]*corev1.Pod) {
			for _, label := range labels {
				if reflect.DeepEqual(pod.Labels, label) {
					ret = append(ret, pod)
				}
			}
		}
		return ret, nil
	}
	return nil, errors.New("pod not found")
}

// GetPodEventMessage return used to save events, only the latest one is saved. make sure it's unique.
func (c *controller) GetPodEventMessage(namespace, kind, name string) string {
	key := PodEventMessagePrefix(name, kind, name)
	if v, ok := c.cachesMap.Load(key); ok {
		return v.(*corev1.Event).Message
	}
	return ""
}
