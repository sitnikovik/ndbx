package refresh

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

const (
	// Name is the name of the job.
	Name = "Refresh Session"
	// Description provides a brief explanation of what the job does.
	Description = "Checking that the application does not recreate a session " +
		"if the user already has a valid session cookie"
)

// Job represents the refresh session job in the autograder process.
type Job struct {
	// steps is a list of step runners that will be executed as part of this job.
	steps []step.Runner
}

// NewJob creates a new Job instance with the provided steps to be executed in this job.
func NewJob(steps ...step.Runner,
) *Job {
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

// Run performs the refresh session operation.
func (j *Job) Run(
	ctx context.Context,
	vars step.Variables,
) error {
	return step.NewList(j.steps...).Run(ctx, vars)
}
