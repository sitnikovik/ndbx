package endpoint_test

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab4/job/events/update/ok/endpoint"
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
		s    *endpoint.Step
		args args
		want want
	}{
		{
			name: "ok",
			s: endpoint.NewStep(
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
			s: endpoint.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPatch(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusInternalServerError,
								Body:       http.NoBody,
							}, assert.AnError
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
			name: "unexpected response",
			s: endpoint.NewStep(
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
			name: "response without session in cookie",
			s: endpoint.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPatch(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusNoContent,
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
				panic: true,
			},
		},
		{
			name: "invalid session found in cookie",
			s: endpoint.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPatch(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusNoContent,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										"X-Session-Id=; HttpOnly; Max-Age=3600; Secure=true",
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
				err: errs.ErrExpectationFailed,
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
