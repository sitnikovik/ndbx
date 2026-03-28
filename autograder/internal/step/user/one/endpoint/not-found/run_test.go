package notfound_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/user/one/endpoint/not-found"
	httpxfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/httpx"
	userfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/user"
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
			name: "ok not found",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithGet(
						func(_ string) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusNotFound,
								Body:       http.NoBody,
							}, nil
						},
					),
				),
				"http://localhost:8000",
				userfx.NewAlexSmith(),
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
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithGet(
						func(_ string) (*http.Response, error) {
							v := `{` +
								`"id": "123",` +
								`"username": "sams3p1ol",` +
								`"full_name": "Sam Sepiol"` +
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
				user.NewUser(
					user.NewID("1"),
					"sams3p1ol",
					"Sam Sepiol",
				),
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
				userfx.NewAlexSmith(),
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
			name: "not found but got non-empty body",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithGet(
						func(_ string) (*http.Response, error) {
							v := `{` +
								`"message": "not found"` +
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
				userfx.NewAlexSmith(),
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
