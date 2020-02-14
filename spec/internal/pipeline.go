package internal

import (
	"fmt"
	"time"
)

type PipelineOpt func(p *PipelineOpts) error

var defPipelineOpts = &PipelineOpts{
	Timeout:             30 * time.Second,
	CronExpr:            make([]string, 0),
	MaxExecutionsAtTime: 1,
}

type PipelineOpts struct {
	Timeout             time.Duration
	CronExpr            []string
	MaxExecutionsAtTime int
	Triggers            []string
}

type Pipeline struct {
	Description string
	Stages      []*Stage
	options     *PipelineOpts
}

func (p *Pipeline) WithOptions(options ...PipelineOpt) *Pipeline {
	opts, err := buildPipelineOpts(options...)
	if err != nil {

	}
	p.options = opts
	return p
}

func buildPipelineOpts(def ...PipelineOpt) (*PipelineOpts, error) {
	var opts PipelineOpts
	opts.Timeout = defPipelineOpts.Timeout
	for _, opt := range def {
		err := opt(&opts)
		if err != nil {
			return nil, err
		}
	}
	return &opts, nil
}

func (p *Pipeline) Run() {
	fmt.Printf("\n%s\n\n", "starting pipeline...")
	for _, stg := range p.Stages {
		fmt.Printf("[%s]\n", stg.Label)
		for _, tsk := range stg.Tasks {
			fmt.Printf("   - %s: %s\n", tsk.Label, tsk.Description)
		}
	}
	fmt.Printf("\n%s\n", "pipeline completed!")
}
