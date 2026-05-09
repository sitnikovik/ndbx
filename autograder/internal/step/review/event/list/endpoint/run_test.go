package endpoint_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/review/event/list/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/step/review/event/list/endpoint/expect"
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
			name: "ok",
			s: impl.NewStep(
				descFixture,
				httpxfk.NewFakeClient(
					httpxfk.WithGet(
						func(_ string) (*http.Response, error) {
							v := `{` +
								`"reviews": [
						{` +
								`"id": "56e2c0b3a2b4c1a5e6f7f8b3",` +
								`"event_id": "12e9c0b1a2b3c3d5e6f7a8b7",` +
								`"comment": "Great!",` +
								`"created_at": "2026-03-14T14:59:32+03:00",` +
								`"created_by": "65e9c0b1a2b3c4d5e6f7a8b9",` +
								`"rating": 5,` +
								`"updated_at": "2026-03-14T14:59:32+03:00"` +
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
				eventFixture,
				"/localhost",
				expect.NewExpectations(
					expect.WithCount(1),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: varsFixture,
			},
			want: want{
				vars:  varsFixture,
				err:   nil,
				panic: false,
			},
		},
		{
			name: "got more than expected",
			s: impl.NewStep(
				descFixture,
				httpxfk.NewFakeClient(
					httpxfk.WithGet(
						func(_ string) (*http.Response, error) {
							v := `{` +
								`"reviews": [
						{` +
								`"id": "56e2c0b3a2b4c1a5e6f7f8b3",` +
								`"event_id": "12e9c0b1a2b3c3d5e6f7a8b7",` +
								`"comment": "Great!",` +
								`"created_at": "2026-03-14T14:59:32+03:00",` +
								`"created_by": "65e9c0b1a2b3c4d5e6f7a8b9",` +
								`"rating": 5,` +
								`"updated_at": "2026-03-14T14:59:32+03:00"` +
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
				eventFixture,
				"/localhost",
				expect.NewExpectations(
					expect.WithCount(0),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: varsFixture,
			},
			want: want{
				vars:  varsFixture,
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "not found but expected",
			s: impl.NewStep(
				descFixture,
				httpxfk.NewFakeClient(
					httpxfk.WithGet(
						func(_ string) (*http.Response, error) {
							v := `{` +
								`"reviews": [],` +
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
				eventFixture,
				"/localhost",
				expect.NewExpectations(
					expect.WithCount(0),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: varsFixture,
			},
			want: want{
				vars:  varsFixture,
				err:   nil,
				panic: false,
			},
		},
		{
			name: "count field mismatch ",
			s: impl.NewStep(
				descFixture,
				httpxfk.NewFakeClient(
					httpxfk.WithGet(
						func(_ string) (*http.Response, error) {
							v := `{` +
								`"reviews": [],` +
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
				eventFixture,
				"/localhost",
				expect.NewExpectations(
					expect.WithCount(0),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: varsFixture,
			},
			want: want{
				vars:  varsFixture,
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "not found but not expected",
			s: impl.NewStep(
				descFixture,
				httpxfk.NewFakeClient(
					httpxfk.WithGet(
						func(_ string) (*http.Response, error) {
							v := `{` +
								`"reviews": [],` +
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
				eventFixture,
				"/localhost",
				expect.NewExpectations(
					expect.WithCount(1),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: varsFixture,
			},
			want: want{
				vars:  varsFixture,
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "event id is not found in variables",
			s: impl.NewStep(
				descFixture,
				httpxfk.NewFakeClient(),
				eventFixture,
				"/localhost",
				expect.NewExpectations(
					expect.WithCount(0),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					v := step.NewVariables()
					v.Del(eventFixture.Hash())
					return v
				}(),
			},
			want: want{
				vars: func() step.Variables {
					v := step.NewVariables()
					v.Del(eventFixture.Hash())
					return v
				}(),
				err:   nil,
				panic: true,
			},
		},
		{
			name: "http failed",
			s: impl.NewStep(
				descFixture,
				httpxfk.NewFakeClient(
					httpxfk.WithGet(
						func(_ string) (*http.Response, error) {
							return nil, assert.AnError
						},
					),
				),
				eventFixture,
				"/localhost",
				expect.NewExpectations(
					expect.WithCount(1),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: varsFixture,
			},
			want: want{
				vars:  varsFixture,
				err:   errs.ErrExternalDependencyFailed,
				panic: false,
			},
		},
		{
			name: "got unexpected status code",
			s: impl.NewStep(
				descFixture,
				httpxfk.NewFakeClient(
					httpxfk.WithGet(
						func(_ string) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusContinue,
								Body:       http.NoBody,
							}, nil
						},
					),
				),
				eventFixture,
				"/localhost",
				expect.NewExpectations(
					expect.WithCount(0),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: varsFixture,
			},
			want: want{
				vars:  varsFixture,
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "got wrong body",
			s: impl.NewStep(
				descFixture,
				httpxfk.NewFakeClient(
					httpxfk.WithGet(
						func(_ string) (*http.Response, error) {
							v := `foo`
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
				eventFixture,
				"/localhost",
				expect.NewExpectations(
					expect.WithCount(0),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: varsFixture,
			},
			want: want{
				vars:  varsFixture,
				err:   nil,
				panic: true,
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
