package step

import (
	"context"
)

// Runner is an interface that defines the Run method for executing a step.
type Runner interface {
	// Run executes the step with the given context and variables and returns an error if the step fails.
	//
	// The vars argument contains the current variables of the step,
	// which can be used to store and retrieve values across steps
	// and these ones will be passed to the next steps so keep in mind
	// that any changes to the vars will affect the next steps.
	Run(
		ctx context.Context,
		vars Variables,
	) error
	// Name returns the name of the step.
	Name() string
	// Description returns a brief explanation of what the step does.
	Description() string
}
