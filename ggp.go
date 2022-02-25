/*
Copyright 2021 The Beijing Gridsum Technology Co., Ltd Authors.

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

package ggp

import (
	corev2 "k8s.io/api/core/v1"
	corev1 "k8s.io/client-go/listers/core/v1"
)

type ControllerService interface {
	// Ready k8s Informer ready status.
	Ready() bool
	// Start start k8s Informer.
	Start() error
	// PodLister is k8s pod lister.
	PodLister() corev1.PodLister
	// GetPod return get the specified pod resource based on the namespace and pod name.
	GetPod(namespace,name string) *corev2.Pod
	// GetPodByNameSpace return get all pods under this namespace.
	GetPodByNameSpace(namespace string) ([]*corev2.Pod,error)
	// GetPodByLabel return get exact matching pods based on namespace and label.
	GetPodByLabel(namespace string,labels map[string]string) ([]*corev2.Pod,error)
	// GetPodEventMessage return used to save events, only the latest one is saved. make sure it's unique.
	GetPodEventMessage(namespace,kind,name string) string
}
