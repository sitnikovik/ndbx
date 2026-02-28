package autograder

import (
	"context"

	style "github.com/sitnikovik/paints/style/text"

	"github.com/sitnikovik/ndbx/autograder/internal/console"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run executes all the jobs in the autograder sequentially.
//
// It returns an error if any job fails or if there are no jobs to run.
func (a *Autograder) Run(
	ctx context.Context,
	vars step.Variables,
) error {
	if len(a.jobs) == 0 {
		return errs.ErrNothingToRun
	}
	for _, j := range a.jobs {
		console.Log(
			"Starting job %s\n%s",
			style.Bold(j.Name()),
			style.Dim(j.Description()),
		)
		err := j.Run(ctx, vars)
		if err != nil {
			return err
		}
	}
	return nil
}
