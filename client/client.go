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

import (
	istio "istio.io/client-go/pkg/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type ManagerClient struct {
	kubeClient  kubernetes.Interface
	istioClient istio.Interface
}

func NewManagerClient(config *Config) (*ManagerClient, error) {
	kubeClient, err := NewKubeClient(config.KubeAPIConfig)
	if err != nil {
		return nil, err
	}
	istioClient, err := NewIstioClient(config.KubeAPIConfig)
	if err != nil {
		return nil, err
	}
	return &ManagerClient{
		kubeClient:  kubeClient,
		istioClient: istioClient,
	}, nil
}

func NewKubeClient(config *KubeAPIConfig) (kubernetes.Interface, error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags(config.Master, config.KubeConfig)
	if err != nil {
		return nil, err
	}
	kubeConfig.QPS = float32(config.QPS)
	kubeConfig.Burst = int(config.Burst)
	kubeConfig.ContentType = config.ContentType
	return kubernetes.NewForConfigOrDie(kubeConfig), nil
}

func NewIstioClient(config *KubeAPIConfig) (istio.Interface, error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags(config.Master, config.KubeConfig)
	if err != nil {
		return nil, err
	}
	kubeConfig.QPS = float32(config.QPS)
	kubeConfig.Burst = int(config.Burst)
	kubeConfig.ContentType = config.ContentType
	return istio.NewForConfigOrDie(kubeConfig), nil
}

func (c *ManagerClient) KubeClient() kubernetes.Interface {
	return c.kubeClient
}

func (c *ManagerClient) IstioClient() istio.Interface {
	return c.istioClient
}
