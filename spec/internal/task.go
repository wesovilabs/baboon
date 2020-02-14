package internal

import (

	"time"
)

type TaskOpt func(p *TaskOpts) error

var defTaskOpts = &TaskOpts{
	Timeout: 30 * time.Second,
}

type TaskOpts struct {
	Label   string
	Timeout time.Duration
}

type Task struct {
	Description string
	Label       string
}

func buildTaskOpts(def ...TaskOpt) (*TaskOpts, error) {
	var opts TaskOpts
	opts.Timeout = defTaskOpts.Timeout
	for _, opt := range def {
		err := opt(&opts)
		if err != nil {
			return nil, err
		}
	}
	return &opts, nil
}
