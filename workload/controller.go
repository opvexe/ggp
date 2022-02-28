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
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv2 "k8s.io/api/autoscaling/v2beta2"
	corev1 "k8s.io/api/core/v1"
	extensions "k8s.io/api/extensions/v1beta1"
	storagev1 "k8s.io/api/storage/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"sync"
	"time"
	"x6t.io/ggp"
)

const (
	// DefaultResyncPeriod is default k8s synchronised time.
	DefaultResyncPeriod = time.Second * 10
	// DefaultStorageClassResyncPeriod is default k8s StorageClass synchronised time.
	DefaultStorageClassResyncPeriod = time.Second * 300
)

type controller struct {
	// client is k8s client.
	client kubernetes.Interface
	// informers is k8s informer.
	informers *Informer
	// listers is k8s lister.
	listers *Lister
	// stopCh is stop channel.
	stopCh <-chan struct{}
	// cachesMap is k8s caches cr.
	cachesMap sync.Map
}

// NewController stopCh is context.Done.
func NewController(clientset kubernetes.Interface, stopCh <-chan struct{}) ggp.ControllerService {
	c := &controller{
		client:    clientset,
		informers: &Informer{},
		listers:   &Lister{},
		stopCh:    stopCh,
		cachesMap: sync.Map{},
	}

	// create informers factory, enable and assign required informers
	infoFactory := informers.NewSharedInformerFactory(clientset, DefaultResyncPeriod)
	c.informers.Namespace = infoFactory.Core().V1().Namespaces().Informer()
	// informer ingress
	c.informers.Ingress = infoFactory.Extensions().V1beta1().Ingresses().Informer()
	c.listers.Ingress = infoFactory.Extensions().V1beta1().Ingresses().Lister()
	// informer service
	c.informers.Service = infoFactory.Core().V1().Services().Informer()
	c.listers.Service = infoFactory.Core().V1().Services().Lister()
	// informer secret
	c.informers.Secret = infoFactory.Core().V1().Secrets().Informer()
	c.listers.Secret = infoFactory.Core().V1().Secrets().Lister()
	// informer StatefulSet
	c.informers.StatefulSet = infoFactory.Apps().V1().StatefulSets().Informer()
	c.listers.StatefulSet = infoFactory.Apps().V1().StatefulSets().Lister()
	// informer deployment
	c.informers.Deployment = infoFactory.Apps().V1().Deployments().Informer()
	c.listers.Deployment = infoFactory.Apps().V1().Deployments().Lister()
	// informer pod
	c.informers.Pod = infoFactory.Core().V1().Pods().Informer()
	c.listers.Pod = infoFactory.Core().V1().Pods().Lister()
	// informer ConfigMap
	c.informers.ConfigMap = infoFactory.Core().V1().ConfigMaps().Informer()
	c.listers.ConfigMap = infoFactory.Core().V1().ConfigMaps().Lister()
	// informer ReplicaSet
	c.informers.ReplicaSet = infoFactory.Apps().V1().ReplicaSets().Informer()
	// informer Endpoints
	c.informers.Endpoints = infoFactory.Core().V1().Endpoints().Informer()
	c.listers.Endpoints = infoFactory.Core().V1().Endpoints().Lister()
	// informer Nodes
	c.informers.Nodes = infoFactory.Core().V1().Nodes().Informer()
	c.listers.Nodes = infoFactory.Core().V1().Nodes().Lister()
	// informer StorageClass
	c.informers.StorageClass = infoFactory.Storage().V1().StorageClasses().Informer()
	c.listers.StorageClass = infoFactory.Storage().V1().StorageClasses().Lister()
	// informer Claims
	c.informers.Claims = infoFactory.Core().V1().PersistentVolumeClaims().Informer()
	c.listers.Claims = infoFactory.Core().V1().PersistentVolumeClaims().Lister()
	// informer Events
	c.informers.Events = infoFactory.Core().V1().Events().Informer()
	// informer HorizontalPodAutoscaler
	c.informers.HorizontalPodAutoscaler = infoFactory.Autoscaling().V2beta2().HorizontalPodAutoscalers().Informer()
	c.listers.HorizontalPodAutoscaler = infoFactory.Autoscaling().V2beta2().HorizontalPodAutoscalers().Lister()

	// add event handler
	c.informers.Namespace.AddEventHandler(c.AddNameSpaceEventHandler())
	c.informers.Ingress.AddEventHandlerWithResyncPeriod(c, DefaultResyncPeriod)
	c.informers.Service.AddEventHandlerWithResyncPeriod(c, DefaultResyncPeriod)
	c.informers.Secret.AddEventHandlerWithResyncPeriod(c, DefaultResyncPeriod)
	c.informers.StatefulSet.AddEventHandlerWithResyncPeriod(c, DefaultResyncPeriod)
	c.informers.Deployment.AddEventHandlerWithResyncPeriod(c, DefaultResyncPeriod)
	c.informers.Pod.AddEventHandlerWithResyncPeriod(c.AddPodEventHandler(), DefaultResyncPeriod)
	c.informers.ConfigMap.AddEventHandlerWithResyncPeriod(c, DefaultResyncPeriod)
	c.informers.ReplicaSet.AddEventHandlerWithResyncPeriod(c, DefaultResyncPeriod)
	c.informers.Endpoints.AddEventHandlerWithResyncPeriod(c.AddEndpointsEventHandler(), DefaultResyncPeriod)
	c.informers.Nodes.AddEventHandlerWithResyncPeriod(c, DefaultResyncPeriod)
	c.informers.StorageClass.AddEventHandlerWithResyncPeriod(c, DefaultStorageClassResyncPeriod)
	c.informers.Claims.AddEventHandlerWithResyncPeriod(c, DefaultResyncPeriod)
	c.informers.Events.AddEventHandlerWithResyncPeriod(c, DefaultResyncPeriod)
	c.informers.HorizontalPodAutoscaler.AddEventHandlerWithResyncPeriod(c, DefaultResyncPeriod)
	return c
}

func (c *controller) Ready() bool {
	return c.informers.Ready()
}

func (c *controller) Start() error {
	c.informers.Start(c.stopCh)
	if !c.Ready() {
		// keep blocking if not ready.
	}
	return nil
}

func (c *controller) OnAdd(obj interface{}) {
	// ingress
	if ingress, ok := obj.(*extensions.Ingress); ok {
		if list, ok := c.cachesMap.Load(c.Prefix(Ingress, ingress.Namespace)); ok {
			list = append(list.([]*extensions.Ingress), ingress)
			c.cachesMap.Store(c.Prefix(Ingress, ingress.Namespace), list)
		} else {
			c.cachesMap.Store(c.Prefix(Ingress, ingress.Namespace), []*extensions.Ingress{ingress})
		}
	}

	// service
	if service, ok := obj.(*corev1.Service); ok {
		if list, ok := c.cachesMap.Load(c.Prefix(Service, service.Namespace)); ok {
			list = append(list.([]*corev1.Service), service)
			c.cachesMap.Store(c.Prefix(Service, service.Namespace), list)
		} else {
			c.cachesMap.Store(c.Prefix(Service, service.Namespace), []*corev1.Service{service})
		}
	}

	// secret
	if secret, ok := obj.(*corev1.Secret); ok {
		if list, ok := c.cachesMap.Load(c.Prefix(Secret, secret.Namespace)); ok {
			list = append(list.([]*corev1.Secret), secret)
			c.cachesMap.Store(c.Prefix(Secret, secret.Namespace), list)
		} else {
			c.cachesMap.Store(c.Prefix(Secret, secret.Namespace), []*corev1.Secret{secret})
		}
	}

	// StatefulSet
	if statefulset, ok := obj.(*appsv1.StatefulSet); ok {
		if list, ok := c.cachesMap.Load(c.Prefix(StatefulSet, statefulset.Namespace)); ok {
			list = append(list.([]*appsv1.StatefulSet), statefulset)
			c.cachesMap.Store(c.Prefix(StatefulSet, statefulset.Namespace), list)
		} else {
			c.cachesMap.Store(c.Prefix(StatefulSet, statefulset.Namespace), []*appsv1.StatefulSet{statefulset})
		}
	}

	// deployment
	if deployment, ok := obj.(*appsv1.Deployment); ok {
		if list, ok := c.cachesMap.Load(c.Prefix(Deployment, deployment.Namespace)); ok {
			list = append(list.([]*appsv1.Deployment), deployment)
			c.cachesMap.Store(c.Prefix(Deployment, deployment.Namespace), list)
		} else {
			c.cachesMap.Store(c.Prefix(Deployment, deployment.Namespace), []*appsv1.Deployment{deployment})
		}
	}

	// configMap
	if configmap, ok := obj.(*corev1.ConfigMap); ok {
		if list, ok := c.cachesMap.Load(c.Prefix(ConfigMap, configmap.Namespace)); ok {
			list = append(list.([]*corev1.ConfigMap), configmap)
			c.cachesMap.Store(c.Prefix(ConfigMap, configmap.Namespace), list)
		} else {
			c.cachesMap.Store(c.Prefix(ConfigMap, configmap.Namespace), []*corev1.ConfigMap{configmap})
		}
	}

	// ReplicaSet
	if replicaset, ok := obj.(*appsv1.ReplicaSet); ok {
		if list, ok := c.cachesMap.Load(c.Prefix(ReplicaSet, replicaset.Namespace)); ok {
			list = append(list.([]*appsv1.ReplicaSet), replicaset)
			c.cachesMap.Store(c.Prefix(ReplicaSet, replicaset.Namespace), list)
		} else {
			c.cachesMap.Store(c.Prefix(ReplicaSet, replicaset.Namespace), []*appsv1.ReplicaSet{replicaset})
		}
	}

	// StorageClass
	if sc, ok := obj.(*storagev1.StorageClass); ok {
		if list, ok := c.cachesMap.Load(c.Prefix(StorageClass, sc.Namespace)); ok {
			list = append(list.([]*storagev1.StorageClass), sc)
			c.cachesMap.Store(c.Prefix(StorageClass, sc.Namespace), list)
		} else {
			c.cachesMap.Store(c.Prefix(StorageClass, sc.Namespace), []*storagev1.StorageClass{sc})
		}
	}

	// Claims
	if claim, ok := obj.(*corev1.PersistentVolumeClaim); ok {
		if list, ok := c.cachesMap.Load(c.Prefix(PersistentVolumeClaim, claim.Namespace)); ok {
			list = append(list.([]*corev1.PersistentVolumeClaim), claim)
			c.cachesMap.Store(c.Prefix(PersistentVolumeClaim, claim.Namespace), list)
		} else {
			c.cachesMap.Store(c.Prefix(PersistentVolumeClaim, claim.Namespace), []*corev1.PersistentVolumeClaim{claim})
		}
	}

	// event
	if event, ok := obj.(*corev1.Event); ok {
		if list, ok := c.cachesMap.Load(c.Prefix(Event, event.Namespace, event.InvolvedObject.Kind, event.InvolvedObject.Name)); ok {
			list = append(list.([]*corev1.Event), event)
			c.cachesMap.Store(c.Prefix(Event, event.Namespace, event.InvolvedObject.Kind, event.InvolvedObject.Name), list)
		} else {
			c.cachesMap.Store(c.Prefix(Event, event.Namespace, event.InvolvedObject.Kind, event.InvolvedObject.Name), []*corev1.Event{event})
		}
	}

	// HorizontalPodAutoscaler
	if hpa, ok := obj.(*autoscalingv2.HorizontalPodAutoscaler); ok {
		if list, ok := c.cachesMap.Load(c.Prefix(HorizontalPodAutoscaler, hpa.Namespace)); ok {
			list = append(list.([]*autoscalingv2.HorizontalPodAutoscaler), hpa)
			c.cachesMap.Store(c.Prefix(HorizontalPodAutoscaler, hpa.Namespace), list)
		} else {
			c.cachesMap.Store(c.Prefix(HorizontalPodAutoscaler, hpa.Namespace), []*autoscalingv2.HorizontalPodAutoscaler{hpa})
		}
	}
}

func (c *controller) OnUpdate(oldObj, newObj interface{}) {

}

func (c *controller) OnDelete(obj interface{}) {

}

// AddNameSpaceEventHandler return namespace event handler.
func (c *controller) AddNameSpaceEventHandler() cache.ResourceEventHandlerFuncs {
	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			ns := obj.(*corev1.Namespace)
			c.cachesMap.Store(NameSpacePrefix(ns.Name), ns)
		},
		DeleteFunc: func(obj interface{}) {
			ns := obj.(*corev1.Namespace)
			c.cachesMap.Delete(NameSpacePrefix(ns.Name))
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			ns := newObj.(*corev1.Namespace)
			c.cachesMap.Store(NameSpacePrefix(ns.Name), ns)
		},
	}
}

// AddPodEventHandler return pod event handler.
func (c *controller) AddPodEventHandler() cache.ResourceEventHandlerFuncs {
	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			pod := obj.(*corev1.Pod)
			if list, ok := c.cachesMap.Load(c.Prefix(Pod, pod.Namespace)); ok {
				list = append(list.([]*corev1.Pod), pod)
				c.cachesMap.Store(c.Prefix(Pod, pod.Namespace), list)
			} else {
				c.cachesMap.Store(c.Prefix(Pod, pod.Namespace), []*corev1.Pod{pod})
			}
		},
		DeleteFunc: func(obj interface{}) {
			pod := obj.(*corev1.Pod)
			if list, ok := c.cachesMap.Load(PodSpacePrefix(pod.Namespace)); ok {
				for i, pods := range list.([]*corev1.Pod) {
					if pods.Name == pod.Name {
						newlist := append(list.([]*corev1.Pod)[:i], list.([]*corev1.Pod)[i+1:]...)
						c.cachesMap.Store(PodSpacePrefix(pod.Namespace), newlist)
						break
					}
				}
			}

		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			pod := newObj.(*corev1.Pod)
			if list, ok := c.cachesMap.Load(PodSpacePrefix(pod.Namespace)); ok {
				for i, oldPod := range list.([]*corev1.Pod) {
					if oldPod.Name == pod.Name {
						list.([]*corev1.Pod)[i] = pod
					}
				}
			}
		},
	}
}

// AddEndpointsEventHandler return endpoints event handler.
func (c *controller) AddEndpointsEventHandler() cache.ResourceEventHandlerFuncs {
	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			ep := obj.(*corev1.Endpoints)
			c.cachesMap.Store(EndpointsSpacePrefix(ep.Namespace), ep)
		},
		DeleteFunc: func(obj interface{}) {
			ep := obj.(*corev1.Endpoints)
			c.cachesMap.Delete(EndpointsSpacePrefix(ep.Namespace))
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			ep := newObj.(*corev1.Endpoints)
			c.cachesMap.Store(EndpointsSpacePrefix(ep.Namespace), ep)
		},
	}
}
