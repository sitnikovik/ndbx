package variable_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

func TestValues_MustUser(t *testing.T) {
	t.Parallel()
	type want struct {
		val   user.User
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
						variable.User,
						user.NewUser(
							user.NewID("1"),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					return vars
				}(),
			),

			want: want{
				val: user.NewUser(
					user.NewID("1"),
					"sams3p1ol",
					"Sam Sepiol",
				),
				panic: false,
			},
		},
		{
			name: "missing session variable",
			v:    variable.NewValues(step.NewVariables()),
			want: want{
				val:   user.User{},
				panic: true,
			},
		},
		{
			name: "invalid session variable type",
			v: variable.NewValues(
				func() step.Variables {
					vars := step.NewVariables()
					vars.Set(variable.User, "not a user")
					return vars
				}(),
			),
			want: want{
				val:   user.User{},
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
				tt.v.MustUser(),
			)
		})
	}
}
