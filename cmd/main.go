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

package main

import (
	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.

	"os"

	"github.com/LavenderQAQ/louvain-scheduler/cmd/operator"
	"github.com/LavenderQAQ/louvain-scheduler/cmd/scheduler"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/klog/v2"
)

var (
	setupLog = klog.New(klog.Background().GetSink())
)

func main() {
	// var metricsAddr string
	// var enableLeaderElection bool
	// var probeAddr string
	// flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	// flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	// flag.BoolVar(&enableLeaderElection, "leader-elect", false,
	// 	"Enable leader election for controller manager. "+
	// 		"Enabling this will ensure there is only one active controller manager.")

	go func() {
		if err := operator.Start(setupLog, ":8080", false, ":8081"); err != nil {
			setupLog.Error(err, "unable to run manager")
			os.Exit(1)
		}
	}()

	if err := scheduler.Start(); err != nil {
		setupLog.Error(err, "unable to run scheduler")
		os.Exit(1)
	}
}
