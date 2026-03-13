package session_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cookie/session"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

func TestSession_MatchVariables(t *testing.T) {
	t.Parallel()
	type args struct {
		vars step.Variables
	}
	type want struct {
		err   error
		panic bool
	}
	tests := []struct {
		name string
		s    session.Session
		args args
		want want
	}{
		{
			name: "ok",
			s: session.NewSession(
				&http.Cookie{
					Name:     session.Name,
					Path:     "/",
					Value:    "123",
					MaxAge:   900,
					HttpOnly: true,
					Secure:   true,
				},
			),
			args: args{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(session.Name, "123")
					vars.Set(variable.SessionTTL, 900*time.Second)
					return vars
				}(),
			},
			want: want{
				err:   nil,
				panic: false,
			},
		},
		{
			name: "session value does not match",
			s: session.NewSession(
				&http.Cookie{
					Name:     session.Name,
					Path:     "/",
					Value:    "123",
					MaxAge:   900,
					HttpOnly: true,
					Secure:   true,
				},
			),
			args: args{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(session.Name, "232")
					vars.Set(variable.SessionTTL, 900*time.Second)
					return vars
				}(),
			},
			want: want{
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "session ttl does not match",
			s: session.NewSession(
				&http.Cookie{
					Name:     session.Name,
					Path:     "/",
					Value:    "123",
					MaxAge:   900,
					HttpOnly: true,
					Secure:   true,
				},
			),
			args: args{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(session.Name, "123")
					vars.Set(variable.SessionTTL, 800*time.Second)
					return vars
				}(),
			},
			want: want{
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "missing session value variable",
			s: session.NewSession(
				&http.Cookie{
					Name:     session.Name,
					Path:     "/",
					Value:    "123",
					MaxAge:   900,
					HttpOnly: true,
					Secure:   true,
				},
			),
			args: args{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(variable.SessionTTL, 900*time.Second)
					return vars
				}(),
			},
			want: want{
				err:   nil,
				panic: true,
			},
		}, {
			name: "missing session ttl variable",
			s: session.NewSession(
				&http.Cookie{
					Name:     session.Name,
					Path:     "/",
					Value:    "123",
					MaxAge:   900,
					HttpOnly: true,
					Secure:   true,
				},
			),
			args: args{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(session.Name, "123")
					return vars
				}(),
			},
			want: want{
				err:   nil,
				panic: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = tt.s.MatchVariables(tt.args.vars)
				})
				return
			}
			assert.ErrorIs(
				t,
				tt.s.MatchVariables(tt.args.vars),
				tt.want.err,
			)
		})
	}
}
