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

type Lister struct {
	Ingress                 v1beta1.IngressLister
	Service                 corev1.ServiceLister
	Secret                  corev1.SecretLister
	StatefulSet             appsv1.StatefulSetLister
	Deployment              appsv1.DeploymentLister
	Pod                     corev1.PodLister
	ConfigMap               corev1.ConfigMapLister
	Endpoints               corev1.EndpointsLister
	Nodes                   corev1.NodeLister
	StorageClass            storagev1.StorageClassLister
	Claims                  corev1.PersistentVolumeClaimLister
	HorizontalPodAutoscaler autoscalingv2.HorizontalPodAutoscalerLister
	// TODO v1alpha3
	Gateways        istio.GatewayLister
	VirtualService  istio.VirtualServiceLister
	DestinationRule istio.DestinationRuleLister
}
