package spec

import (
	"context"
	"github.com/wesovilabs/baboon/spec/internal"
)

type Root interface {
}

func Pipeline(desc string, stages []StageSpec, options ...internal.PipelineOpt) *internal.Pipeline {
	pipeline := &internal.Pipeline{
		Description: desc,
		Stages:      make([]*internal.Stage, len(stages)),
	}
	for i, stg := range stages {
		pipeline.Stages[i] = stg()
	}
	return pipeline.WithOptions(options...)
}

type StageSpec func() *internal.Stage

func Stage(desc string, tasks []TaskSpec, options ...internal.StageOpt) StageSpec {
	return func() *internal.Stage {
		stg:=&internal.Stage{
			Description: desc,
			Tasks: make([]*internal.Task,len(tasks)),
		}
		for i, tsk := range tasks {
			stg.Tasks[i] = tsk()
		}
		return stg.WithOptions(options...)
	}
}

type TaskSpec func() *internal.Task

func Task(desc string, fn func(ctx context.Context), options ...internal.TaskOpt) TaskSpec {
	return func() *internal.Task {
		return &internal.Task{
			Description: desc,
		}
	}
}
