package stage

import (
	"github.com/wesovilabs/baboon/spec/internal"
	"time"
)



func WithTimeout(d time.Duration) internal.StageOpt {
	return func(p *internal.StageOpts) error {
		p.Timeout = d
		return nil
	}
}

func WithLabel(label string) internal.StageOpt {
	return func(p *internal.StageOpts) error {
		p.Label = label
		return nil
	}
}
