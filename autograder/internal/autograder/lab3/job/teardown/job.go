package teardown

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

const (
	// Name is the name of the teardown job for Lab 3.
	Name = "Teardown"
	// Description provides a brief explanation of what the Lab 3 teardown job does.
	Description = "Cleans up the environment for Lab 3"
)

// Job represents the setup job for Lab 3 in the autograder process.
type Job struct {
	// steps is a list of step runners that will be executed as part of the teardown job.
	steps []step.Runner
}

// NewJob creates a new Job instance with the provided step runners.
func NewJob(steps ...step.Runner) *Job {
	return &Job{
		steps: steps,
	}
}

// Name returns the name of the job.
func (j *Job) Name() string {
	return Name
}

// Description returns a brief explanation of what the job does.
func (j *Job) Description() string {
	return Description
}

// Run performs the expire operation for the job by executing all the steps in order.
func (j *Job) Run(
	ctx context.Context,
	vars step.Variables,
) error {
	return step.NewList(j.steps...).Run(ctx, vars)
}
