package endpoint_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab4/job/events/one/not-found/endpoint"
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
		s    *endpoint.Step
		args args
		want want
	}{
		{
			name: "not found",
			s: endpoint.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithGet(
						func(_ string) (*http.Response, error) {
							v := `{"message":"Not found"}`
							return &http.Response{
								StatusCode: http.StatusNotFound,
								Body: func() io.ReadCloser {
									return io.NopCloser(strings.NewReader(v))
								}(),
								ContentLength: int64(len(v)),
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
				vars:  step.NewVariables(),
				err:   nil,
				panic: false,
			},
		},
		{
			name: "found",
			s: endpoint.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithGet(
						func(_ string) (*http.Response, error) {
							v := `{` +
								`"id": "2ip3ue32-9ejnojkdsdp932u8eji",` +
								`"title": "test title",` +
								`"description": "test description",` +
								`"location": {` +
								`"address": "test location"` +
								`},` +
								`"created_at": "2024-01-01T00:00:00Z",` +
								`"created_by": "test_user",` +
								`"started_at": "2024-01-01T01:00:00Z",` +
								`"finished_at": "2024-01-01T02:00:00Z"` +
								`}`
							return &http.Response{
								StatusCode: http.StatusOK,
								Body: func() io.ReadCloser {
									return io.NopCloser(
										strings.NewReader(v),
									)
								}(),
								ContentLength: int64(len(v)),
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
				vars:  step.NewVariables(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "req failed",
			s: endpoint.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithGet(
						func(_ string) (*http.Response, error) {
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
				vars:  step.NewVariables(),
				err:   errs.ErrHTTPFailed,
				panic: false,
			},
		},
		{
			name: "not found but unexpected response",
			s: endpoint.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithGet(
						func(_ string) (*http.Response, error) {
							v := `{"message": "not found"}`
							return &http.Response{
								StatusCode: http.StatusOK,
								Body: func() io.ReadCloser {
									return io.NopCloser(
										strings.NewReader(v),
									)
								}(),
								ContentLength: int64(len(v)),
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
				vars:  step.NewVariables(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "not found but unexpected http status code",
			s: endpoint.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithGet(
						func(_ string) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusTeapot,
								Body:       http.NoBody,
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
				vars:  step.NewVariables(),
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
