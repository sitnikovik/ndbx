package endpoint_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/consts/cookie"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/health/endpoint"
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
					httpfk.WithGet(
						func(_ string) (*http.Response, error) {
							return &http.Response{
								StatusCode: 200,
								Body:       http.NoBody,
							}, nil
						},
					),
				),
				"/localhost:8080",
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			wantVars:  step.NewVariables(),
			wantErr:   nil,
			wantPanic: false,
		},
		{
			name: "http failed",
			step: endpoint.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithGet(
						func(_ string) (*http.Response, error) {
							return nil, assert.AnError
						},
					),
				),
				"/localhost:8080",
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
			name: "unexpected http status code",
			step: endpoint.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithGet(
						func(_ string) (*http.Response, error) {
							return &http.Response{
								StatusCode: 500,
								Body:       http.NoBody,
							}, nil
						},
					),
				),
				"/localhost:8080",
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
			name: "has session cookie",
			step: endpoint.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithGet(
						func(_ string) (*http.Response, error) {
							return &http.Response{
								StatusCode: 200,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										cookie.SessionName + "=test-session-123; HttpOnly",
									},
								},
							}, nil
						},
					),
				),
				"/localhost:8080",
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			wantVars:  step.NewVariables(),
			wantErr:   errs.ErrExpectationFailed,
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
