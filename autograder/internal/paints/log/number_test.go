package log_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
)

func TestNumber(t *testing.T) {
	t.Parallel()
	t.Run("int positive", func(t *testing.T) {
		t.Parallel()
		got := log.Number(10)
		want := "\033[36m10\033[0m"
		assert.Equal(t, want, got)
	})
	t.Run("int negative", func(t *testing.T) {
		t.Parallel()
		got := log.Number(-5)
		want := "\033[36m-5\033[0m"
		assert.Equal(t, want, got)
	})
	t.Run("int zero", func(t *testing.T) {
		t.Parallel()
		got := log.Number(0)
		want := "\033[36m0\033[0m"
		assert.Equal(t, want, got)
	})
}
