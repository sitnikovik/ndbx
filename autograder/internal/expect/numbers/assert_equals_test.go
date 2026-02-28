package numbers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
)

func TestAssertEquals(t *testing.T) {
	t.Parallel()
	t.Run("int ok", func(t *testing.T) {
		t.Parallel()
		assert.NoError(t, numbers.AssertEquals(1, 1))
	})
	t.Run("int not equal", func(t *testing.T) {
		t.Parallel()
		assert.ErrorIs(
			t,
			numbers.AssertEquals(1, 2),
			errs.ErrExpectationFailed,
		)
	})
	t.Run("zeros ok", func(t *testing.T) {
		t.Parallel()
		assert.NoError(t, numbers.AssertEquals(0, 0))
	})
	t.Run("float ok", func(t *testing.T) {
		t.Parallel()
		assert.NoError(t, numbers.AssertEquals(1.5, 1.5))
	})
	t.Run("float not equal", func(t *testing.T) {
		t.Parallel()
		assert.ErrorIs(
			t,
			numbers.AssertEquals(1.5, 2.5),
			errs.ErrExpectationFailed,
		)
	})
}
