package wait_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/step"
	"github.com/sitnikovik/ndbx/autograder/internal/step/wait"
)

func TestWaitingStep_Run(t *testing.T) {
	t.Parallel()
	t.Run("after sleep duration", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		err := wait.
			NewWaitingStep(100*time.Millisecond).
			Run(
				ctx,
				step.NewVariables(),
			)
		assert.NoError(t, err)
	})
	t.Run("before sleep duration", func(t *testing.T) {
		t.Parallel()
		ctx, cancel := context.WithTimeout(
			context.Background(),
			50*time.Millisecond,
		)
		defer cancel()
		err := wait.
			NewWaitingStep(100*time.Millisecond).
			Run(
				ctx,
				step.NewVariables(),
			)
		assert.ErrorIs(t, err, context.DeadlineExceeded)
	})
	t.Run("context cancelled", func(t *testing.T) {
		t.Parallel()
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		err := wait.
			NewWaitingStep(100*time.Millisecond).
			Run(
				ctx,
				step.NewVariables(),
			)
		assert.ErrorIs(t, err, context.Canceled)
	})
}
