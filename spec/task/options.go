package task

import (
	"github.com/wesovilabs/baboon/spec/internal"
	"time"
)

func WithTimeout(d time.Duration) internal.TaskOpt {
	return func(p *internal.TaskOpts) error {
		p.Timeout = d
		return nil
	}
}

func WithLabel(label string) internal.TaskOpt {
	return func(p *internal.TaskOpts) error {
		p.Label = label
		return nil
	}
}
