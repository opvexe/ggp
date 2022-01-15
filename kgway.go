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

package kgway

import "context"

type (
	Deployment struct {
		Name       string
		NameSpace  string
		Images     string
		Replicas   [3]int32
		CreateTime string
	}

	DeploymentStore interface {
		All(ctx context.Context) ([]*Deployment, error)
		List(ctx context.Context, namespace string) ([]*Deployment, error)
	}

	Istio struct {
		GW Gateway
		DS DestinationRule
		VS VirtualService
	}

	Gateway struct {
		Name      string
		Namespace string
		HostPort  uint32
	}

	Gateways []Gateway

	DestinationRule struct {
		Name      string
		Namespace string
		Host      string
	}

	VirtualService struct {
		Name            string
		Namespace       string
		Gateways        Gateways
		DestinationHost string
		DestinationPort int
	}

	IstioStore interface {
		Add(ctx context.Context, istio Istio) error
		Delete(ctx context.Context, name, namespace string) error
		Update(ctx context.Context) error
		List(ctx context.Context, namespace string) ([]*Gateway, error)
	}
)
