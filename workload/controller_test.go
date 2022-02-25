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
	"context"
	"k8s.io/client-go/kubernetes"
	"reflect"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	"testing"
	"x6t.io/ggp"
	"x6t.io/ggp/client"
)

func NewMockKubeClient() (kubernetes.Interface, context.Context) {
	clientset, err := client.NewKubeClient(client.NewConfig().KubeAPIConfig)
	if err != nil {
		panic(err)
	}
	return clientset, signals.SetupSignalHandler()
}

func TestNewController(t *testing.T) {
	// NewMockKubeClient new kube client.
	clientset, ctx := NewMockKubeClient()

	type args struct {
		clientset kubernetes.Interface
		stopCh    <-chan struct{}
	}
	tests := []struct {
		name string
		args args
		want ggp.ControllerService
	}{
		{
			name: "new controller",
			args: args{
				clientset: clientset,
				stopCh:    ctx.Done(),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewController(tt.args.clientset, tt.args.stopCh); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewController() = %v, want %v", got, tt.want)
			}
		})
	}
	<-ctx.Done()
}
