package endpoint_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	consts "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/consts/cookie"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/session/endpoint"
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
	tests := []struct {
		name      string
		step      *endpoint.Step
		args      args
		wantVars  step.Variables
		wantErr   error
		wantPanic bool
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
			wantVars: func() step.Variables {
				vars := step.NewVariables()
				vars.Set(
					consts.SessionName,
					"0123456789abcdef0123456789abcdef",
				)
				return vars
			}(),
			wantErr:   nil,
			wantPanic: false,
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
			wantVars:  step.NewVariables(),
			wantErr:   errs.ErrHTTPFailed,
			wantPanic: false,
		},
		{
			name: "HTTP status code 200",
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
			wantVars:  step.NewVariables(),
			wantErr:   errs.ErrInvalidHTTPStatus,
			wantPanic: false,
		},
		{
			name: "not empty response body",
			step: endpoint.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithPostJSON(
						func(
							_ string,
							_ io.Reader,
						) (*http.Response, error) {
							return &http.Response{
								StatusCode:    http.StatusCreated,
								Body:          io.NopCloser(strings.NewReader("{}")),
								ContentLength: 2,
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
			wantVars:  step.NewVariables(),
			wantErr:   errs.ErrExpectationFailed,
			wantPanic: false,
		},
		{
			name: "empty session in cookie",
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
			wantVars:  step.NewVariables(),
			wantErr:   errs.ErrMissedCookie,
			wantPanic: false,
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
								StatusCode: http.StatusCreated,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										"X-Session-Id=0123456789abcdef0123456789abcdef",
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
			wantVars:  step.NewVariables(),
			wantErr:   errs.ErrMissedCookie,
			wantPanic: false,
		},
		{
			name: "invalid session cookie",
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
			wantVars:  step.NewVariables(),
			wantErr:   errs.ErrExpectationFailed,
			wantPanic: false,
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
								StatusCode: http.StatusCreated,
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
			wantVars:  step.NewVariables(),
			wantErr:   errs.ErrMissedCookie,
			wantPanic: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.wantPanic {
				assert.Panics(
					t,
					func() {
						_ = tt.step.Run(
							tt.args.ctx,
							tt.args.vars,
						)
					},
				)
			} else {
				assert.ErrorIs(
					t,
					tt.step.Run(
						tt.args.ctx,
						tt.args.vars,
					),
					tt.wantErr,
				)
			}
			assert.Equal(
				t,
				tt.wantVars,
				tt.args.vars,
			)
		})
	}
}
