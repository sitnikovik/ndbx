package endpoint_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab4/job/events/one/ok/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	httpfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/httpx"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
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
			name: "ok",
			s: endpoint.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithGet(
						func(_ string) (*http.Response, error) {
							v := `{` +
								`"id": "1",` +
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
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.Event,
						event.NewEvent(
							event.NewID("1"),
							event.NewContent(
								"test title",
								"test description",
							),
							event.NewLocation("test location"),
							event.NewCreated(
								timex.MustRFC3339("2024-01-01T00:00:00Z"),
								user.NewIdentity("test_user"),
							),
							event.NewDates(
								timex.MustRFC3339("2024-01-01T01:00:00Z"),
								timex.MustRFC3339("2024-01-01T02:00:00Z"),
							),
						),
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.Event,
						event.NewEvent(
							event.NewID("1"),
							event.NewContent(
								"test title",
								"test description",
							),
							event.NewLocation("test location"),
							event.NewCreated(
								timex.MustRFC3339("2024-01-01T00:00:00Z"),
								user.NewIdentity("test_user"),
							),
							event.NewDates(
								timex.MustRFC3339("2024-01-01T01:00:00Z"),
								timex.MustRFC3339("2024-01-01T02:00:00Z"),
							),
						),
					)
					return vars
				}(),
				err:   nil,
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
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.Event,
						event.NewEvent(
							event.NewID("1"),
							event.NewContent(
								"test title",
								"test description",
							),
							event.NewLocation("test location"),
							event.NewCreated(
								timex.MustRFC3339("2024-01-01T00:00:00Z"),
								user.NewIdentity("test_user"),
							),
							event.NewDates(
								timex.MustRFC3339("2024-01-01T01:00:00Z"),
								timex.MustRFC3339("2024-01-01T02:00:00Z"),
							),
						),
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.Event,
						event.NewEvent(
							event.NewID("1"),
							event.NewContent(
								"test title",
								"test description",
							),
							event.NewLocation("test location"),
							event.NewCreated(
								timex.MustRFC3339("2024-01-01T00:00:00Z"),
								user.NewIdentity("test_user"),
							),
							event.NewDates(
								timex.MustRFC3339("2024-01-01T01:00:00Z"),
								timex.MustRFC3339("2024-01-01T02:00:00Z"),
							),
						),
					)
					return vars
				}(),
				err:   errs.ErrHTTPFailed,
				panic: false,
			},
		},
		{
			name: "unexpected http status code",
			s: endpoint.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithGet(
						func(_ string) (*http.Response, error) {
							v := `{"message": "ooops..."}`
							return &http.Response{
								StatusCode: http.StatusTeapot,
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
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.Event,
						event.NewEvent(
							event.NewID("1"),
							event.NewContent(
								"test title",
								"test description",
							),
							event.NewLocation("test location"),
							event.NewCreated(
								timex.MustRFC3339("2024-01-01T00:00:00Z"),
								user.NewIdentity("test_user"),
							),
							event.NewDates(
								timex.MustRFC3339("2024-01-01T01:00:00Z"),
								timex.MustRFC3339("2024-01-01T02:00:00Z"),
							),
						),
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.Event,
						event.NewEvent(
							event.NewID("1"),
							event.NewContent(
								"test title",
								"test description",
							),
							event.NewLocation("test location"),
							event.NewCreated(
								timex.MustRFC3339("2024-01-01T00:00:00Z"),
								user.NewIdentity("test_user"),
							),
							event.NewDates(
								timex.MustRFC3339("2024-01-01T01:00:00Z"),
								timex.MustRFC3339("2024-01-01T02:00:00Z"),
							),
						),
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
