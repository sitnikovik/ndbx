package variable_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

func TestValues_MustEventID(t *testing.T) {
	t.Parallel()
	type want struct {
		val   string
		panic bool
	}
	tests := []struct {
		name string
		v    variable.Values
		want want
	}{
		{
			name: "ok",
			v: variable.NewValues(
				func() step.Variables {
					vars := step.NewVariables()
					vars.Set(variable.EventID, "event-id-123")
					return vars
				}(),
			),
			want: want{
				val:   "event-id-123",
				panic: false,
			},
		},
		{
			name: "missing EventID variable",
			v:    variable.NewValues(step.NewVariables()),
			want: want{
				val:   "",
				panic: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = tt.v.MustEventID()
				})
				return
			}
			assert.Equal(
				t,
				tt.want.val,
				tt.v.MustEventID(),
			)
		})
	}
}
