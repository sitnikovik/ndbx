package numbers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
)

func TestAssertEqualOrGreater(t *testing.T) {
	t.Parallel()
	t.Run("int equal", func(t *testing.T) {
		t.Parallel()
		assert.NoError(t, numbers.AssertEqualOrGreater(1, 1))
	})
	t.Run("int greater", func(t *testing.T) {
		t.Parallel()
		assert.NoError(t, numbers.AssertEqualOrGreater(1, 2))
	})
	t.Run("int lower", func(t *testing.T) {
		t.Parallel()
		assert.Error(t, numbers.AssertEqualOrGreater(1, 0))
	})
}
