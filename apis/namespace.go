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

package apis

import (
	"context"
	meta1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type NamespaceService struct {
	kubernetes.Interface
}

func NewNamespaceService(kube kubernetes.Interface) *NamespaceService {
	return &NamespaceService{Interface: kube}
}

func (s *NamespaceService) List(ctx context.Context) ([]string, error) {
	opts := meta1.ListOptions{}
	ns, err := s.CoreV1().Namespaces().List(ctx, opts)
	if err != nil {
		return nil, err
	}

	namespaces := make([]string, len(ns.Items))
	for _, item := range ns.Items {
		namespaces = append(namespaces, item.Name)
	}
	return namespaces, nil
}
