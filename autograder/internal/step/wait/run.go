package wait

import (
	"context"
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/console"
	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run sleeps for the specified duration or until the context is canceled.
func (s *WaitingStep) Run(
	ctx context.Context,
	_ step.Variables,
) error {
	console.Log(
		"😴 waiting for %ss",
		log.Number(s.duration.Seconds()),
	)
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(s.duration):
		return nil
	}
}
