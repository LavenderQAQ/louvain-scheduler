package main

import (
	"os"

	"github.com/LavenderQAQ/louvain-scheduler/pkg/plugin"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main() {
	cmd := app.NewSchedulerCommand(
		app.WithPlugin(plugin.Name, plugin.New),
	)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
