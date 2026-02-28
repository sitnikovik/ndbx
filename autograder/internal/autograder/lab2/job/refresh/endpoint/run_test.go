package endpoint_test

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	consts "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/consts/cookie"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/refresh/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	httpfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/httpx"
)

func TestStep_Run(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx  context.Context
		vars step.Variables
	}
	type want struct {
		vars  step.Variables
		err   error
		panic bool
	}
	tests := []struct {
		name string
		step *endpoint.Step
		args args
		want want
	}{
		{
			name: "ok",
			step: endpoint.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithPostJSON(
						func(
							_ string,
							_ io.Reader,
						) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusOK,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										"X-Session-Id=0123456789abcdef0123456789abcdef; HttpOnly; Max-Age=3600",
									},
								},
							}, nil
						},
					),
				),
				"/localhost",
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						consts.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						consts.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					return vars
				}(),
				err:   nil,
				panic: false,
			},
		},
		{
			name: "HTTP request failed",
			step: endpoint.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithPostJSON(
						func(
							_ string,
							_ io.Reader,
						) (*http.Response, error) {
							return nil, assert.AnError
						},
					),
				),
				"/localhost",
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				vars:  step.NewVariables(),
				err:   assert.AnError,
				panic: false,
			},
		},
		{
			name: "invalid HTTP status",
			step: endpoint.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithPostJSON(
						func(
							_ string,
							_ io.Reader,
						) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusCreated,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										"X-Session-Id=0123456789abcdef0123456789abcdef; HttpOnly; Max-Age=3600",
									},
								},
							}, nil
						},
					),
				),
				"/localhost",
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				vars:  step.NewVariables(),
				err:   errs.ErrInvalidHTTPStatus,
				panic: false,
			},
		},
		{
			name: "response with body",
			step: endpoint.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithPostJSON(
						func(
							_ string,
							_ io.Reader,
						) (*http.Response, error) {
							return &http.Response{
								StatusCode:    http.StatusOK,
								ContentLength: 1,
								Body:          http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										"X-Session-Id=0123456789abcdef0123456789abcdef; HttpOnly; Max-Age=3600",
									},
								},
							}, nil
						},
					),
				),
				"/localhost",
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				vars:  step.NewVariables(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "missing session cookie",
			step: endpoint.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithPostJSON(
						func(
							_ string,
							_ io.Reader,
						) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusOK,
								Body:       http.NoBody,
								Header:     http.Header{},
							}, nil
						},
					),
				),
				"/localhost",
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				vars:  step.NewVariables(),
				err:   nil,
				panic: true,
			},
		},
		{
			name: "session cookie without value",
			step: endpoint.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithPostJSON(
						func(
							_ string,
							_ io.Reader,
						) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusOK,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										"X-Session-Id=; HttpOnly; Max-Age=3600",
									},
								},
							}, nil
						},
					),
				),
				"/localhost",
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				vars:  step.NewVariables(),
				err:   errs.ErrMissedCookie,
				panic: false,
			},
		},
		{
			name: "session cookie without HttpOnly flag",
			step: endpoint.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithPostJSON(
						func(
							_ string,
							_ io.Reader,
						) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusOK,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										"X-Session-Id=0123456789abcdef0123456789abcdef; Max-Age=3600",
									},
								},
							}, nil
						},
					),
				),
				"/localhost",
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				vars:  step.NewVariables(),
				err:   errs.ErrMissedCookie,
				panic: false,
			},
		},
		{
			name: "session cookie without MaxAge flag",
			step: endpoint.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithPostJSON(
						func(
							_ string,
							_ io.Reader,
						) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusOK,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										"X-Session-Id=0123456789abcdef0123456789abcdef; HttpOnly",
									},
								},
							}, nil
						},
					),
				),
				"/localhost",
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				vars:  step.NewVariables(),
				err:   errs.ErrMissedCookie,
				panic: false,
			},
		},
		{
			name: "invalid session ID format",
			step: endpoint.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithPostJSON(
						func(
							_ string,
							_ io.Reader,
						) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusOK,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										"X-Session-Id=invalid-session-id; HttpOnly; Max-Age=3600",
									},
								},
							}, nil
						},
					),
				),
				"/localhost",
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				vars:  step.NewVariables(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "session ID mismatch",
			step: endpoint.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithPostJSON(
						func(
							_ string,
							_ io.Reader,
						) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusOK,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										"X-Session-Id=0123456789abcdef0123456789abcdef; HttpOnly; Max-Age=3600",
									},
								},
							}, nil
						},
					),
				),
				"/localhost",
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						consts.SessionName,
						"fedcba9876543210fedcba9876543210",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						consts.SessionName,
						"fedcba9876543210fedcba9876543210",
					)
					return vars
				}(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = tt.step.Run(
						tt.args.ctx,
						tt.args.vars,
					)
				})
				return
			}
			assert.ErrorIs(
				t,
				tt.step.Run(
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
