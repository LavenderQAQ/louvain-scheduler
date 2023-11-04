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
