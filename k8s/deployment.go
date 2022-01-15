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
	"fmt"

	core "k8s.io/api/core/v1"
	meta1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"x6t.io/kgway"
)

type DeploymentService struct {
	kubernetes.Interface
}

func NewDeploymentService(kube kubernetes.Interface) *DeploymentService {
	return &DeploymentService{
		Interface: kube,
	}
}

func (s *DeploymentService) List(ctx context.Context, namespace string) ([]*kgway.Deployment, error) {
	opts := meta1.ListOptions{}
	list, err := s.AppsV1().Deployments(namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}

	items := make([]*kgway.Deployment, len(list.Items))
	for i, item := range list.Items {
		items[i] = &kgway.Deployment{
			Name:      item.Name,
			NameSpace: item.Namespace,
			Images:    s.Images(item.Spec.Template.Spec.Containers),
			Replicas: [3]int32{
				item.Status.Replicas,
				item.Status.AvailableReplicas,
				item.Status.UnavailableReplicas,
			},
			CreateTime: item.CreationTimestamp.Format("2006-01-02 15:03:04"),
		}
	}
	return items, nil
}

func (s *DeploymentService) All(ctx context.Context) ([]*kgway.Deployment, error) {
	// TODO

	return nil, nil
}

func (s *DeploymentService) Images(containers []core.Container) string {
	images := containers[0].Image
	if length := len(containers); length > 1 {
		images += fmt.Sprintf(" + 其他%d个镜像", length-1)
	}
	return images
}
