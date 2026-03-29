package endpoint_test

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab4/job/events/update/unauth/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	httpxfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/httpx"
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
					httpxfk.WithPatch(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusUnauthorized,
								Body:       http.NoBody,
							}, nil
						},
					),
				),
				"http://localhost",
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.EventID,
						"1",
					)
					return vars
				}(),
			},
			want: want{
				err: nil,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.EventID,
						"1",
					)
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "http failed",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPatch(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return nil, assert.AnError
						},
					),
				),
				"http://localhost",
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.EventID,
						"1",
					)
					return vars
				}(),
			},
			want: want{
				err: errs.ErrHTTPFailed,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.EventID,
						"1",
					)
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "unexpected ok response",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPatch(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusOK,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										"X-Session-Id=0123456789abcdef0123456789abcdef; HttpOnly; Max-Age=3600; Secure=true",
									},
								},
							}, nil
						},
					),
				),
				"http://localhost",
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.EventID,
						"1",
					)
					return vars
				}(),
			},
			want: want{
				err: errs.ErrInvalidHTTPStatus,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.EventID,
						"1",
					)
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "changed",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPatch(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusNoContent,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										"X-Session-Id=0123456789abcdef0123456789abcdef; HttpOnly; Max-Age=3600; Secure=true",
									},
								},
							}, nil
						},
					),
				),
				"http://localhost",
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.EventID,
						"1",
					)
					return vars
				}(),
			},
			want: want{
				err: errs.ErrInvalidHTTPStatus,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.EventID,
						"1",
					)
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "unexpected not found response",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPatch(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusNotFound,
								Body:       http.NoBody,
							}, nil
						},
					),
				),
				"http://localhost",
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.EventID,
						"1",
					)
					return vars
				}(),
			},
			want: want{
				err: errs.ErrInvalidHTTPStatus,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.EventID,
						"1",
					)
					return vars
				}(),
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
		})
	}
}
