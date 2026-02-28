package step

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

// List represents a list of steps to be executed as part of a job.
type List struct {
	// steps defines the steps to be executed sequentially.
	steps []Runner
}

// NewList creates a new List instance with the provided steps to run.
func NewList(steps ...Runner) *List {
	return &List{
		steps: steps,
	}
}

// Run executes each step in the list sequentially.
//
// If any step returns an error, the execution will stop immediately and the error will be returned.
// If all steps execute successfully, nil will be returned.
// If the list is empty, an error indicating that there is nothing to run will be returned.
func (l *List) Run(ctx context.Context, vars Variables) error {
	if len(l.steps) == 0 {
		return errs.ErrNothingToRun
	}
	for _, step := range l.steps {
		if err := step.Run(ctx, vars); err != nil {
			return err
		}
	}
	return nil
}
