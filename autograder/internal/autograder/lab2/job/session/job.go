package session

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

const (
	// Name is the name of the job.
	Name = "Create Session"
	// Description provides a brief explanation of what the job does.
	Description = "Checking that the application correctly creates a session in Redis" +
		" and sets the appropriate cookie in the response."
)

// Job represents the create session job in the autograder process.
type Job struct {
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

// Run performs the create session operation.
func (j *Job) Run(
	ctx context.Context,
	vars step.Variables,
) error {
	return step.NewList(j.steps...).Run(ctx, vars)
}
