package endpoint_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cookie/session"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/events/create/ok/endpoint"
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
					httpxfk.WithPostJSON(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusCreated,
								Body: func() io.ReadCloser {
									v := `{"id":"123"}`
									return io.NopCloser(strings.NewReader(v))
								}(),
								Header: http.Header{
									"Set-Cookie": []string{
										"X-Session-Id=0123456789abcdef0123456789abcdef; HttpOnly; Max-Age=3600; Secure=true",
									},
								},
								ContentLength: int64(len(`{"id":"123"}`)),
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
					vars.Set(session.Name, "0123456789abcdef0123456789abcdef")
					vars.Set(variable.SessionTTL, 3600*time.Second)
					return vars
				}(),
			},
			want: want{
				err: nil,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(session.Name, "0123456789abcdef0123456789abcdef")
					vars.Set(variable.SessionTTL, 3600*time.Second)
					vars.Set(variable.EventID, "123")
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "http failed",
			s: endpoint.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPostJSON(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return nil, assert.AnError
						},
					),
				),
				"http://localhost",
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
			name: "unexpected response",
			s: endpoint.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPostJSON(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusInternalServerError,
								Body: func() io.ReadCloser {
									v := `{"id":"123"}`
									return io.NopCloser(strings.NewReader(v))
								}(),
								ContentLength: int64(len(`{"id":"123"}`)),
							}, nil
						},
					),
				),
				"http://localhost",
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
			name: "created a new session cookie",
			s: endpoint.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPostJSON(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusCreated,
								Body: func() io.ReadCloser {
									v := `{"id":"123"}`
									return io.NopCloser(strings.NewReader(v))
								}(),
								ContentLength: int64(len(`{"id":"123"}`)),
								Header: http.Header{
									"Set-Cookie": []string{
										"X-Session-Id=1123456789abcdef0123456789abcdef; HttpOnly; Max-Age=3600; Secure=true",
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
					vars.Set(session.Name, "0123456789abcdef0123456789abcdef")
					vars.Set(variable.SessionTTL, 3600*time.Second)
					return vars
				}(),
			},
			want: want{
				err: errs.ErrExpectationFailed,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(session.Name, "0123456789abcdef0123456789abcdef")
					vars.Set(variable.SessionTTL, 3600*time.Second)
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "invalid session in cookie",
			s: endpoint.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPostJSON(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusCreated,
								Body: func() io.ReadCloser {
									v := `{"id":"123"}`
									return io.NopCloser(strings.NewReader(v))
								}(),
								ContentLength: int64(len(`{"id":"123"}`)),
								Header: http.Header{
									"Set-Cookie": []string{
										"X-Session-Id=invalidsessionvalue; HttpOnly; Max-Age=3600; Secure=true",
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
					vars.Set(session.Name, "0123456789abcdef0123456789abcdef")
					vars.Set(variable.SessionTTL, 3600*time.Second)
					return vars
				}(),
			},
			want: want{
				err: errs.ErrExpectationFailed,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(session.Name, "0123456789abcdef0123456789abcdef")
					vars.Set(variable.SessionTTL, 3600*time.Second)
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "unexpected session in cookie",
			s: endpoint.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPostJSON(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusCreated,
								Body: func() io.ReadCloser {
									v := `{"id":"123"}`
									return io.NopCloser(strings.NewReader(v))
								}(),
								ContentLength: int64(len(`{"id":"123"}`)),
								Header: http.Header{
									"Set-Cookie": []string{
										"X-Session-Id=1123456789abcdef0123456789abcdef; HttpOnly; Max-Age=3600; Secure=true",
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
					vars.Set(session.Name, "0123456789abcdef0123456789abcdef")
					vars.Set(variable.SessionTTL, 3600*time.Second)
					return vars
				}(),
			},
			want: want{
				err: errs.ErrExpectationFailed,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(session.Name, "0123456789abcdef0123456789abcdef")
					vars.Set(variable.SessionTTL, 3600*time.Second)
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "got unexpected json body",
			s: endpoint.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPostJSON(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusCreated,
								Body: func() io.ReadCloser {
									v := `{"status":"ok"}`
									return io.NopCloser(strings.NewReader(v))
								}(),
								ContentLength: int64(len(`{"status":"ok"}`)),
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
					vars.Set(session.Name, "0123456789abcdef0123456789abcdef")
					vars.Set(variable.SessionTTL, 3600*time.Second)
					return vars
				}(),
			},
			want: want{
				err: errs.ErrExpectationFailed,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(session.Name, "0123456789abcdef0123456789abcdef")
					vars.Set(variable.SessionTTL, 3600*time.Second)
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
			assert.Equal(
				t,
				tt.want.vars,
				tt.args.vars,
			)
		})
	}
}
