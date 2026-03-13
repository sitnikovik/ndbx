package setup

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

const (
	// Name is the name of the setup job for Lab 3.
	Name = "Setup"
	// Description provides a brief explanation of what the Lab 3 setup job does.
	Description = "Prepares the environment for Lab 3 by ensuring all necessary services are running and configured correctly."
)

// Job represents the setup job for Lab 3 in the autograder process.
type Job struct {
	steps []step.Runner
}

// NewJob creates a new Job instance with the provided steps.
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

// Run performs the expire operation for the job.
func (j *Job) Run(
	ctx context.Context,
	vars step.Variables,
) error {
	return step.NewList(j.steps...).Run(ctx, vars)
}
