package numbers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
)

func TestAssertEmpty(t *testing.T) {
	t.Parallel()
	t.Run("int zero", func(t *testing.T) {
		t.Parallel()
		assert.NoError(t, numbers.AssertEmpty(0))
	})
	t.Run("int not zero", func(t *testing.T) {
		t.Parallel()
		assert.Error(t, numbers.AssertEmpty(1))
	})
	t.Run("float zero", func(t *testing.T) {
		t.Parallel()
		assert.NoError(t, numbers.AssertEmpty(0.0))
	})
	t.Run("float not zero", func(t *testing.T) {
		t.Parallel()
		assert.Error(t, numbers.AssertEmpty(1.5))
	})
}
