package wait_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/step/wait"
)

func TestWaitingStep_Name(t *testing.T) {
	t.Parallel()
	assert.Equal(
		t,
		wait.Name,
		wait.
			NewWaitingStep(time.Second).
			Name(),
	)
}

func TestWaitingStep_Description(t *testing.T) {
	t.Parallel()
	assert.Equal(
		t,
		wait.Description,
		wait.
			NewWaitingStep(time.Second).
			Description(),
	)
}
