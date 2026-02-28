package autograder

import "github.com/sitnikovik/ndbx/autograder/internal/step"

// Runner is an interface that defines the Run method for executing the autograder jobs.
type Runner interface {
	step.Runner
}
