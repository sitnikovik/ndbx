package variable_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user/session"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestValues_MustSession(t *testing.T) {
	t.Parallel()
	type want struct {
		val   session.Session
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
						variable.Session,
						session.NewSession(
							session.NewID("1"),
							session.NewDates(
								timex.MustParse(time.RFC3339, "2026-01-03T11:00:00Z"),
								timex.MustParse(time.RFC3339, "2026-01-03T12:00:00Z"),
							),
						),
					)
					return vars
				}(),
			),

			want: want{
				val: session.NewSession(
					session.NewID("1"),
					session.NewDates(
						timex.MustParse(time.RFC3339, "2026-01-03T11:00:00Z"),
						timex.MustParse(time.RFC3339, "2026-01-03T12:00:00Z"),
					),
				),
				panic: false,
			},
		},
		{
			name: "missing session variable",
			v:    variable.NewValues(step.NewVariables()),
			want: want{
				val:   session.Session{},
				panic: true,
			},
		},
		{
			name: "invalid session variable type",
			v: variable.NewValues(
				func() step.Variables {
					vars := step.NewVariables()
					vars.Set(variable.Session, "not a session")
					return vars
				}(),
			),
			want: want{
				val:   session.Session{},
				panic: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(
					t,
					func() {
						_ = tt.v.MustSession()
					},
				)
				return
			}
			assert.Equal(
				t,
				tt.want.val,
				tt.v.MustSession(),
			)
		})
	}
}
