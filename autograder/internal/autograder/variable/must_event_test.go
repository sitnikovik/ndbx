package variable_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	eventfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/event"
)

func TestValues_MustEvent(t *testing.T) {
	t.Parallel()
	type want struct {
		val   event.Event
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
					vars.Set(
						variable.Event,
						eventfx.NewTestEvent(),
					)
					return vars
				}(),
			),
			want: want{
				val:   eventfx.NewTestEvent(),
				panic: false,
			},
		},
		{
			name: "got id",
			v: variable.NewValues(
				func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.Event,
						eventfx.NewTestEvent().ID(),
					)
					return vars
				}(),
			),
			want: want{
				val:   event.Event{},
				panic: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = tt.v.MustEvent()
				})
				return
			}
			assert.Equal(
				t,
				tt.want.val,
				tt.v.MustEvent(),
			)
		})
	}
}
