package git

import (
	"context"
	"github.com/wesovilabs/baboon/spec"
	"github.com/wesovilabs/baboon/spec/task"
)

var CloneRepository spec.TaskSpec = spec.Task("Clone Repository", func(ctx context.Context) {

})

var CommitChanges spec.TaskSpec = spec.Task("Commit changes", func(ctx context.Context) {

})

var PushChanges spec.TaskSpec = spec.Task("Push changes into repository", func(ctx context.Context) {

})

var PullChanges = func(label string) spec.TaskSpec {

	return spec.Task("Pull changes from repository", func(ctx context.Context) {

	}, task.WithLabel(label))
}
