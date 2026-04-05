package endpoint_test

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cookie/session"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/user/auth/login"
	httpxfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/httpx"
	userfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/user"
	cookiefx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/cookie/session"
	sidfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/cookie/session/id"
)

var (
	userFixture = userfx.NewAlexSmith()
)

func TestStep_Run(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx  context.Context
		vars step.Variables
	}
	type want struct {
		err   error
		vars  step.Variables
		panic bool
	}
	tests := []struct {
		name string
		s    *impl.Step
		args args
		want want
	}{
		{
			name: "ok",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPostJSON(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusNoContent,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										cookiefx.NewOKSession(),
									},
								},
							}, nil
						},
					),
				),
				"http://localhost",
				userFixture,
				"qwerty1234",
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				err: nil,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						session.Name,
						sidfx.OK,
					)
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "http request fails",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPostJSON(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return nil, assert.AnError
						},
					),
				),
				"http://localhost",
				userFixture,
				"qwerty1234",
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				err:   errs.ErrHTTPFailed,
				vars:  step.NewVariables(),
				panic: false,
			},
		},
		{
			name: "got non-empty response",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPostJSON(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode:    http.StatusNoContent,
								Body:          http.NoBody,
								ContentLength: 10,
							}, nil
						},
					),
				),
				"http://localhost",
				userFixture,
				"qwerty1234",
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				err:   errs.ErrExpectationFailed,
				vars:  step.NewVariables(),
				panic: false,
			},
		},
		{
			name: "got invalid session cookie",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPostJSON(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusNoContent,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										"X-Session-Id=invalid-session-id; HttpOnly; Max-Age=3600; Secure=true",
									},
								},
							}, nil
						},
					),
				),
				"http://localhost",
				userFixture,
				"qwerty1234",
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				err:   errs.ErrExpectationFailed,
				vars:  step.NewVariables(),
				panic: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = tt.s.Run(
						tt.args.ctx,
						tt.args.vars,
					)
				})
				return
			}
			assert.ErrorIs(
				t,
				tt.s.Run(
					tt.args.ctx,
					tt.args.vars,
				),
				tt.want.err,
			)
			assert.Equal(
				t,
				tt.want.vars,
				tt.args.vars,
			)
		})
	}
}
