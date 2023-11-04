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

package scheduler

import (
	"github.com/LavenderQAQ/louvain-scheduler/pkg/plugin"
	_ "k8s.io/component-base/logs/json/register"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func Start() error {
	cmd := app.NewSchedulerCommand(
		app.WithPlugin(plugin.Name, plugin.New),
	)

	return cmd.Execute()
}
