package numbers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
)

func TestAssertPositive(t *testing.T) {
	t.Parallel()
	t.Run("int positive", func(t *testing.T) {
		t.Parallel()
		assert.NoError(t, numbers.AssertPositive(1))
	})
	t.Run("zero", func(t *testing.T) {
		t.Parallel()
		assert.ErrorIs(
			t,
			numbers.AssertPositive(0),
			errs.ErrExpectationFailed,
		)
	})
	t.Run("int negative", func(t *testing.T) {
		t.Parallel()
		assert.ErrorIs(
			t,
			numbers.AssertPositive(-1),
			errs.ErrExpectationFailed,
		)
	})
	t.Run("float positive", func(t *testing.T) {
		t.Parallel()
		assert.NoError(t, numbers.AssertPositive(1.5))
	})
	t.Run("float negative", func(t *testing.T) {
		t.Parallel()
		assert.ErrorIs(
			t,
			numbers.AssertPositive(-1.5),
			errs.ErrExpectationFailed,
		)
	})
}
