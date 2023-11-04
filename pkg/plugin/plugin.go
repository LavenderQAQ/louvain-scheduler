/*
Copyright 2023.

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

package plugin

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const Name = "louvain-scheduler"

type louvainScheduler struct{}

// var _ framework.FilterPlugin = &louvainScheduler{}
// var _ framework.PreScorePlugin = &louvainScheduler{}

func New(_ runtime.Object, _ framework.Handle) (framework.Plugin, error) {
	return &louvainScheduler{}, nil
}

func (ls *louvainScheduler) Name() string {
	return Name
}
