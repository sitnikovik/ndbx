package endpoint_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/users/list/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/user/list/by/endpoint"
	httpxfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/httpx"
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
		s    *impl.Step
		args args
		want want
	}{
		{
			name: "found",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithGet(
						func(_ string) (*http.Response, error) {
							v := `{` +
								`"users": [` +
								`{` +
								`"id": "1",` +
								`"username": "sams3p1ol",` +
								`"full_name": "Sam Sepiol"` +
								`}` +
								`],` +
								`"count": 1` +
								`}`
							return &http.Response{
								StatusCode: http.StatusOK,
								Body: func() io.ReadCloser {
									return io.NopCloser(strings.NewReader(v))
								}(),
								ContentLength: int64(len(v)),
							}, nil
						},
					),
				),
				"http://localhost:8000",
				body.NewBody(
					body.WithFullName("sepiol"),
				),
				[]user.User{
					user.NewUser(
						user.NewID("1"),
						"sams3p1ol",
						"Sam Sepiol",
					),
				},
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
			name: "not found but expected to be",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithGet(
						func(_ string) (*http.Response, error) {
							v := `{` +
								`"users": [],` +
								`"count": 0` +
								`}`
							return &http.Response{
								StatusCode: http.StatusOK,
								Body: func() io.ReadCloser {
									return io.NopCloser(strings.NewReader(v))
								}(),
								ContentLength: int64(len(v)),
							}, nil
						},
					),
				),
				"http://localhost:8000",
				body.NewBody(
					body.WithFullName("sepiol"),
				),
				[]user.User{
					user.NewUser(
						user.NewID("1"),
						"sams3p1ol",
						"Sam Sepiol",
					),
				},
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
			name: "found but count is mismatch",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithGet(
						func(_ string) (*http.Response, error) {
							v := `{` +
								`"users": [],` +
								`"count": 1` +
								`}`
							return &http.Response{
								StatusCode: http.StatusOK,
								Body: func() io.ReadCloser {
									return io.NopCloser(strings.NewReader(v))
								}(),
								ContentLength: int64(len(v)),
							}, nil
						},
					),
				),
				"http://localhost:8000",
				body.NewBody(
					body.WithFullName("sepiol"),
				),
				[]user.User{},
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
			name: "http failed",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithGet(
						func(_ string) (*http.Response, error) {
							return nil, assert.AnError
						},
					),
				),
				"http://localhost:8000",
				body.NewBody(
					body.WithFullName("sepiol"),
				),
				[]user.User{},
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
			name: "not found and got 404",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithGet(
						func(_ string) (*http.Response, error) {
							v := `{` +
								`"users": [],` +
								`"count": 0` +
								`}`
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
				"http://localhost:8000",
				body.NewBody(
					body.WithFullName("sepiol"),
				),
				[]user.User{},
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
			name: "not found and got empty content",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithGet(
						func(_ string) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusOK,
								Body:       http.NoBody,
							}, nil
						},
					),
				),
				"http://localhost:8000",
				body.NewBody(
					body.WithFullName("sepiol"),
				),
				[]user.User{},
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
				assert.Equal(
					t,
					tt.want.vars,
					tt.args.vars,
				)
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
