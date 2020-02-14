package pipeline

import (
	"github.com/wesovilabs/baboon/spec/internal"
	"time"
)



func WithTimeout(d time.Duration) internal.PipelineOpt  {
	return func(p *internal.PipelineOpts) error {
		p.Timeout = d
		return nil
	}
}

func WithCron(expr []string) internal.PipelineOpt {
	return func(p *internal.PipelineOpts) error {
		p.CronExpr = expr
		return nil
	}
}

func WithMaxExecutionsAtTime(executions int) internal.PipelineOpt {
	return func(p *internal.PipelineOpts) error {
		p.MaxExecutionsAtTime = executions
		return nil
	}
}

func WithTriggers(triggers ...string) internal.PipelineOpt {
	return func(p *internal.PipelineOpts) error {
		p.Triggers = triggers
		return nil
	}
}
