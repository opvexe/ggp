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

import "path/filepath"

const (
	DefaultNameSpacePrefix = "ggp-namespace"
	DefaultPodSpacePrefix  = "ggp-pod"
)

func NameSpacePrefix(name string) string {
	return filepath.Join(DefaultNameSpacePrefix, name)
}

func PodSpacePrefix(name string) string {
	return filepath.Join(DefaultPodSpacePrefix, name)
}

func PodEventMessagePrefix(namespace,kind,name string) string {
	return filepath.Join(namespace,kind,name)
}

