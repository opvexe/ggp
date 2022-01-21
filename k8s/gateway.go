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

package k8s

import (
	"context"
	networkingv1alpha3 "istio.io/api/networking/v1alpha3"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	istioclient "istio.io/client-go/pkg/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"x6t.io/kgway"
)

type GatewayService struct {
	istioclient.Clientset
}

func NewGatewayService(clientset istioclient.Clientset) *GatewayService {
	return &GatewayService{
		Clientset: clientset,
	}
}

func (s *GatewayService) Add(ctx context.Context, is *kgway.Istio) error {
	opt := metav1.CreateOptions{}
	gw, err := s.Clientset.NetworkingV1alpha3().Gateways(is.GW.Namespace).
		Create(ctx, &v1alpha3.Gateway{
			ObjectMeta: metav1.ObjectMeta{
				Name:      is.GW.Name,
				Namespace: is.GW.Namespace,
			},
			Spec: networkingv1alpha3.Gateway{
				Servers: []*networkingv1alpha3.Server{
					&networkingv1alpha3.Server{
						Port: &networkingv1alpha3.Port{
							Number:   is.GW.HostPort,
							Protocol: "TCP",
							Name:     "tcp-0",
						},
						Hosts: []string{
							"*",
						},
					},
				},
				Selector: map[string]string{
					"kubeedge": "edgemesh-gateway",
				},
			},
		}, opt)
	if err != nil {
		return err
	}

}
