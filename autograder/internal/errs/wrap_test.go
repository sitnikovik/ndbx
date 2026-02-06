package errs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

func TestWrap(t *testing.T) {
	t.Parallel()
	t.Run("nil error", func(t *testing.T) {
		t.Parallel()
		err := errs.Wrap(nil, "test")
		assert.Nil(t, err)
	})
	t.Run("nil error with format args", func(t *testing.T) {
		t.Parallel()
		err := errs.Wrap(nil, "additional context: %s", "details")
		assert.Nil(t, err)
	})
	t.Run("non-nil error", func(t *testing.T) {
		t.Parallel()
		err := errs.Wrap(assert.AnError, "additional context: %s", "details")
		assert.Error(t, err)
		assert.ErrorIs(t, err, assert.AnError)
		assert.Contains(t, err.Error(), "additional context: details")
	})
}
