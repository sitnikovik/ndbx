package step

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// FakeRunner is a test implementation of the step.Runner interface that allows.
type FakeRunner struct {
	fn func(ctx context.Context, vars step.Variables) error
}

// FakeRunnerOption defines a functional option for configuring the FakeRunner.
type FakeRunnerOption func(*FakeRunner)

// WithRunFunc sets the function that will be executed when the FakeRunner's Run method is called.
func WithRunFunc(fn func(ctx context.Context, vars step.Variables) error) FakeRunnerOption {
	return func(r *FakeRunner) {
		r.fn = fn
	}
}

// WithOkRun sets the function to return nil when the FakeRunner's Run method is called.
func WithOkRun() FakeRunnerOption {
	return WithRunFunc(func(_ context.Context, _ step.Variables) error {
		return nil
	})
}

// WithErrRun sets the function to return the specified error when the FakeRunner's Run method is called.
func WithErrRun(err error) FakeRunnerOption {
	return WithRunFunc(func(_ context.Context, _ step.Variables) error {
		return err
	})
}

// NewFakeRunner creates a new FakeRunner with the provided options.
func NewFakeRunner(opts ...FakeRunnerOption) step.Runner {
	r := &FakeRunner{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// Run executes the fake runner's function if it is set,
// otherwise it does nothing and returns nil.
func (r FakeRunner) Run(
	ctx context.Context,
	vars step.Variables,
) error {
	if r.fn != nil {
		return r.fn(ctx, vars)
	}
	return nil
}

// Name returns a default name for the fake runner.
func (r FakeRunner) Name() string {
	return "Fake Runner"
}

// Description returns a default description for the fake runner.
func (r FakeRunner) Description() string {
	return "A fake runner used for testing purposes."
}
