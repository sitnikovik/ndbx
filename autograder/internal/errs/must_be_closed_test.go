package errs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

func TestMustBeClosed(t *testing.T) {
	t.Parallel()
	t.Run("err", func(t *testing.T) {
		defer func() {
			r := recover()
			assert.NotNil(t, r)
			err, ok := r.(error)
			require.True(t, ok)
			assert.ErrorContains(
				t,
				err,
				"failed to close resource",
			)
		}()
		errs.MustBeClosed(assert.AnError)
	})
	t.Run("no err", func(t *testing.T) {
		assert.NotPanics(
			t,
			func() {
				errs.MustBeClosed(nil)
			},
		)
	})
}
