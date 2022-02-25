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

import "k8s.io/client-go/tools/cache"

type Informer struct {
	Namespace               cache.SharedIndexInformer
	Ingress                 cache.SharedIndexInformer
	Service                 cache.SharedIndexInformer
	Secret                  cache.SharedIndexInformer
	StatefulSet             cache.SharedIndexInformer
	Deployment              cache.SharedIndexInformer
	Pod                     cache.SharedIndexInformer
	ConfigMap               cache.SharedIndexInformer
	ReplicaSet              cache.SharedIndexInformer
	Endpoints               cache.SharedIndexInformer
	Nodes                   cache.SharedIndexInformer
	StorageClass            cache.SharedIndexInformer
	Claims                  cache.SharedIndexInformer
	Events                  cache.SharedIndexInformer
	HorizontalPodAutoscaler cache.SharedIndexInformer
	// TODO
	Gateways        cache.SharedIndexInformer
	VirtualService  cache.SharedIndexInformer
	DestinationRule cache.SharedIndexInformer
}

func (i *Informer) Ready() bool {
	if i.Namespace.HasSynced() &&
		i.Ingress.HasSynced() &&
		i.Service.HasSynced() &&
		i.Secret.HasSynced() &&
		i.StatefulSet.HasSynced() &&
		i.Deployment.HasSynced() &&
		i.Pod.HasSynced() &&
		i.ConfigMap.HasSynced() &&
		i.Nodes.HasSynced() &&
		i.Events.HasSynced() &&
		i.HorizontalPodAutoscaler.HasSynced() &&
		i.StorageClass.HasSynced() &&
		i.Claims.HasSynced() &&
		// TODO
		i.Gateways.HasSynced() &&
		i.VirtualService.HasSynced() &&
		i.DestinationRule.HasSynced() {
		return true
	}
	return false
}

func (i *Informer) Start(stop <-chan struct{}) {
	go i.Namespace.Run(stop)
	go i.Ingress.Run(stop)
	go i.Service.Run(stop)
	go i.Secret.Run(stop)
	go i.StatefulSet.Run(stop)
	go i.Deployment.Run(stop)
	go i.Pod.Run(stop)
	go i.ConfigMap.Run(stop)
	go i.ReplicaSet.Run(stop)
	go i.Endpoints.Run(stop)
	go i.Nodes.Run(stop)
	go i.StorageClass.Run(stop)
	go i.Events.Run(stop)
	go i.HorizontalPodAutoscaler.Run(stop)
	go i.Claims.Run(stop)
	// TODO
	go i.Gateways.Run(stop)
	go i.VirtualService.Run(stop)
	go i.DestinationRule.Run(stop)
}
