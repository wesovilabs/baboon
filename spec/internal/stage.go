package internal

import (

	"time"
)

type StageOpt func(p *StageOpts) error

var defStageOpts = &StageOpts{
	Timeout: 30 * time.Second,
}

type StageOpts struct {
	Label   string
	Timeout time.Duration
}

type Stage struct {
	Description string
	Label       string
	Tasks       []*Task
}

func (s *Stage) WithOptions(options ...StageOpt) *Stage {
	_, err := buildStageOpts(options...)
	if err != nil {

	}
	//s.options = opts
	return s
}

func buildStageOpts(def ...StageOpt) (*StageOpts, error) {
	var opts StageOpts
	opts.Timeout = defStageOpts.Timeout
	for _, opt := range def {
		err := opt(&opts)
		if err != nil {
			return nil, err
		}
	}
	return &opts, nil
}
