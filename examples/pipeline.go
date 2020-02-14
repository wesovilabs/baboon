package main

import (
	"context"
	"fmt"
	. "github.com/wesovilabs/baboon/spec"
	"github.com/wesovilabs/baboon/spec/stage"
	"github.com/wesovilabs/baboon/spec/task"

	"github.com/wesovilabs/baboon/task/git"
	"time"
)

var stgSetUp = Stage("setup project", []TaskSpec{
	git.CloneRepository,
	git.PushChanges,
})

var stgChecks = Stage("Verify code format", []TaskSpec{

	Task("format code", func(ctx context.Context) {
		fmt.Println("running go fmt...")
	}),

	Task("analyse code", func(ctx context.Context) {
		fmt.Println("running linters...")
	}),

}, stage.WithLabel("checks"))

var stgTests = Stage("Run tests", []TaskSpec{

	Task("test unit", func(ctx context.Context) {
		fmt.Println("running unit tests...")
	}, task.WithLabel("test-unit")),

	Task("test int", func(ctx context.Context) {
		fmt.Println("running integration tests...")
	}, task.WithLabel("test-int")),

}, stage.WithLabel("tests"))

var pipeline = Pipeline("Development pipeline", []StageSpec{
	Stage("setup project", []TaskSpec{
		git.PullChanges("pull-changes"),
		git.CloneRepository,
		git.PushChanges,
	}, stage.WithTimeout(10*time.Second)),
	stgChecks,
	stgTests,
})

func main() {
	pipeline.Run()
}
