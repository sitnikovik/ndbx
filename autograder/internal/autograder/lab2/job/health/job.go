package health

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

const (
	// Name is the name of the health check job.
	Name = "Health Check"
	// Description is a brief explanation of what the health check job does.
	Description = "Checks the health of the application by verifying connectivity to Redis" +
		" and making an HTTP request to the application."
)

// Job represents the health check job.
type Job struct {
	// steps is a slice of steps
	// that defines the steps to be executed in this job.
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

// Run performs the health check operation for the job.
func (j *Job) Run(
	ctx context.Context,
	vars step.Variables,
) error {
	return step.NewList(j.steps...).Run(ctx, vars)
}
