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

package ggp

import "context"

type Deployment struct {
	Name       string   `json:"name"`
	NameSpace  string   `json:"name_space"`
	Images     string   `json:"images"`
	Replicas   [3]int32 `json:"replicas"`
	CreateTime string   `json:"create_time"`
}

type RuntimeNamespace interface {
	List(ctx context.Context) ([]string, error)
}

type RuntimeDeployment interface {
	List(ctx context.Context, namespace string) ([]*Deployment, error)
	All(ctx context.Context) ([]*Deployment, error)
}
