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
	"path/filepath"
)

const (
	DefaultNameSpacePrefix               = "ggp-namespace"
	DefaultIngressSpacePrefix            = "ggp-ingress"
	DefaultPodSpacePrefix                = "ggp-pod"
	DefaultServicePrefix                 = "ggp-service"
	DefaultSecretPrefix                  = "ggp-secret"
	DefaultStatefulSetSpacePrefix        = "ggp-stateful-set"
	DefaultDeploymentPrefix              = "ggp-deployment"
	DefaultConfigMapPrefix               = "ggp-configmap"
	DefaultReplicaSetSpacePrefix         = "ggp-replica-set"
	DefaultStorageClassSpacePrefix       = "ggp-storage-class"
	DefaultPersistentVolumeClaimPrefix   = "ggp-persistentvolume-claim"
	DefaultEndPointsPrefix               = "ggp-endpoints"
	DefaultHorizontalPodAutoscalerPrefix = "ggp-horizontal-pod-autoscaler"
)

type PrefixType int

const (
	NameSpace PrefixType = iota
	Ingress
	Pod
	Service
	Secret
	StatefulSet
	Event
	Deployment
	ConfigMap
	ReplicaSet
	StorageClass
	PersistentVolumeClaim
	Endpoints
	HorizontalPodAutoscaler
)

func (p PrefixType) Strings(s ...string) string {
	switch p {
	case NameSpace:
		return filepath.Join(DefaultNameSpacePrefix, s[0])
	case Ingress:
		return filepath.Join(DefaultIngressSpacePrefix, s[0])
	case Pod:
		return filepath.Join(DefaultPodSpacePrefix, s[0])
	case Service:
		return filepath.Join(DefaultServicePrefix, s[0])
	case Secret:
		return filepath.Join(DefaultSecretPrefix, s[0])
	case StatefulSet:
		return filepath.Join(DefaultStatefulSetSpacePrefix, s[0])
	case Deployment:
		return filepath.Join(DefaultDeploymentPrefix, s[0])
	case ConfigMap:
		return filepath.Join(DefaultConfigMapPrefix, s[0])
	case ReplicaSet:
		return filepath.Join(DefaultReplicaSetSpacePrefix, s[0])
	case StorageClass:
		return filepath.Join(DefaultStorageClassSpacePrefix, s[0])
	case PersistentVolumeClaim:
		return filepath.Join(DefaultPersistentVolumeClaimPrefix, s[0])
	case Endpoints:
		return filepath.Join(DefaultEndPointsPrefix, s[0])
	case HorizontalPodAutoscaler:
		return filepath.Join(DefaultHorizontalPodAutoscalerPrefix, s[0])
	case Event:
		if len(s) != 3 {
			return ""
		}
		return filepath.Join(s[0], s[1], s[2])
	default:
		return ""
	}
}

// Prefix return get the corresponding prefix according to the input type.
func (c *controller) Prefix(t PrefixType, s ...string) string {
	return t.Strings(s...)
}

// NameSpacePrefix Namespace related prefix.
// Deprecated: will be removed in the next commit.
func NameSpacePrefix(name string) string {
	return filepath.Join(DefaultNameSpacePrefix, name)
}

// PodSpacePrefix Pod related prefix.
// Deprecated: will be removed in the next commit.
func PodSpacePrefix(namespace string) string {
	return filepath.Join(DefaultPodSpacePrefix, namespace)
}

// EndpointsSpacePrefix related prefix
// Deprecated: will be removed in the next commit.
func EndpointsSpacePrefix(namespace string) string {
	return filepath.Join(DefaultEndPointsPrefix, namespace)
}
