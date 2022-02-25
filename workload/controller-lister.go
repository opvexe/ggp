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
	istio "istio.io/client-go/pkg/listers/networking/v1alpha3"
	appsv1 "k8s.io/client-go/listers/apps/v1"
	autoscalingv2 "k8s.io/client-go/listers/autoscaling/v2beta2"
	corev1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/listers/extensions/v1beta1"
	storagev1 "k8s.io/client-go/listers/storage/v1"
)

func (c *controller) IngressLister() v1beta1.IngressLister {
	return c.listers.Ingress
}

func (c *controller) ServiceLister() corev1.ServiceLister {
	return c.listers.Service
}

func (c *controller) SecretLister() corev1.SecretLister {
	return c.listers.Secret
}

func (c *controller) StatefulSetLister() appsv1.StatefulSetLister {
	return c.listers.StatefulSet
}

func (c *controller) DeploymentLister() appsv1.DeploymentLister {
	return c.listers.Deployment
}

// 	PodLister The following two ways are the same.
//  pods, err := c.PodLister().Pods(v1.NamespaceAll).List(labels.NewSelector())
//  pods,err := kubectl.CoreV1().Pods(v1.NamespaceAll).List(metav1.ListOptions{
func (c *controller) PodLister() corev1.PodLister {
	return c.listers.Pod
}

func (c *controller) ConfigMapLister() corev1.ConfigMapLister {
	return c.listers.ConfigMap
}

func (c *controller) EndpointsLister() corev1.EndpointsLister {
	return c.listers.Endpoints
}

func (c *controller) NodeLister() corev1.NodeLister {
	return c.listers.Nodes
}

func (c *controller) StorageClassLister() storagev1.StorageClassLister {
	return c.listers.StorageClass
}

func (c *controller) PersistentVolumeClaimLister() corev1.PersistentVolumeClaimLister {
	return c.listers.Claims
}

func (c *controller) HorizontalPodAutoscalerLister() autoscalingv2.HorizontalPodAutoscalerLister {
	return c.listers.HorizontalPodAutoscaler
}

func (c *controller) GatewayLister() istio.GatewayLister {
	return c.listers.Gateways
}

func (c *controller) VirtualServiceLister() istio.VirtualServiceLister {
	return c.listers.VirtualService
}

func (c *controller) DestinationRuleLister() istio.DestinationRuleLister {
	return c.listers.DestinationRule
}
