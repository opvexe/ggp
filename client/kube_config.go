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

package client

import "k8s.io/apimachinery/pkg/runtime"

const (
	// DefaultKubeQPS return default kubernetes qps.
	DefaultKubeQPS = 100
	// DefaultKubeBurst return default kubernetes burst.
	DefaultKubeBurst = 200
)

// KubeAPIConfig indicates the configuration for interacting with k8s server.
type KubeAPIConfig struct {
	// Master indicates the address of the Kubernetes API server (overrides any value in KubeConfig)
	// such as https://127.0.0.1:8443
	// default ""
	// Note: Can not use "omitempty" option,  It will affect the output of the default configuration file
	Master string `json:"master"`
	// ContentType indicates the ContentType of message transmission when interacting with k8s
	// default "application/vnd.kubernetes.protobuf"
	ContentType string `json:"contentType,omitempty"`
	// QPS to while talking with kubernetes apiserve
	// default 100
	QPS int32 `json:"qps,omitempty"`
	// Burst to use while talking with kubernetes apiserver
	// default 200
	Burst int32 `json:"burst,omitempty"`
	// KubeConfig indicates the path to kubeConfig file with authorization and master location information.
	// default "/root/.kube/config"
	// +Required
	KubeConfig string `json:"kubeConfig"`
}

// NewKubeAPIConfig return default kubernetes api config.
func NewKubeAPIConfig() *KubeAPIConfig {
	return &KubeAPIConfig{
		Master:      "",
		ContentType: runtime.ContentTypeJSON, // often use runtime.ContentTypeProtobuf.
		QPS:         DefaultKubeQPS,
		Burst:       DefaultKubeBurst,
		KubeConfig:  "",
	}
}
