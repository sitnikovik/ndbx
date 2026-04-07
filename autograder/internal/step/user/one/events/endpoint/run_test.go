package endpoint_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/get/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event/reaction"
	"github.com/sitnikovik/ndbx/autograder/internal/app/reaction/count"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/user/one/events/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/step/user/one/events/expect"
	httpfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/httpx"
	userfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/user"
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
		s    *impl.Step
		args args
		want want
	}{
		{
			name: "ok found",
			s: impl.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithGet(
						func(_ string) (*http.Response, error) {
							v := `{` +
								`"events": [` +
								`{` +
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
				"http://localhost",
				userfx.NewAlexSmith(),
				body.NewBody(),
				[]event.Event{
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
				},
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewAlexSmith().Hash(),
						"123",
					)
					return vars
				}(),
			},
			want: want{
				err: nil,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewAlexSmith().Hash(),
						"123",
					)
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "ok found with reactions",
			s: impl.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithGet(
						func(_ string) (*http.Response, error) {
							v := `{` +
								`"events": [` +
								`{` +
								`"id": "1",` +
								`"title": "test title",` +
								`"description": "test description",` +
								`"location": {` +
								`"address": "test location"` +
								`},` +
								`"created_at": "2024-01-01T00:00:00Z",` +
								`"created_by": "test_user",` +
								`"started_at": "2024-01-01T01:00:00Z",` +
								`"finished_at": "2024-01-01T02:00:00Z",` +
								`"reactions": {` +
								`"likes": 24,` +
								`"dislikes": 3` +
								`}` +
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
				"http://localhost",
				userfx.NewAlexSmith(),
				body.NewBody(),
				[]event.Event{
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
				},
				impl.WithExpectations(
					expect.NewExpectations(
						expect.WithReactions(
							[]reaction.Reactions{
								reaction.NewReactions(
									reaction.WithCounts(
										count.NewCounts(
											count.WithLikes(24),
											count.WithDislikes(3),
										),
									),
								),
							},
						),
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewAlexSmith().Hash(),
						"123",
					)
					return vars
				}(),
			},
			want: want{
				err: nil,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewAlexSmith().Hash(),
						"123",
					)
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "ok not found",
			s: impl.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithGet(
						func(_ string) (*http.Response, error) {
							v := `{` +
								`"events": [],` +
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
				"http://localhost",
				userfx.NewAlexSmith(),
				body.NewBody(),
				nil,
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewAlexSmith().Hash(),
						"123",
					)
					return vars
				}(),
			},
			want: want{
				err: nil,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewAlexSmith().Hash(),
						"123",
					)
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "http request failed",
			s: impl.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithGet(
						func(_ string) (*http.Response, error) {
							return nil, assert.AnError
						},
					),
				),
				"http://localhost",
				userfx.NewAlexSmith(),
				body.NewBody(),
				[]event.Event{
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
				},
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewAlexSmith().Hash(),
						"123",
					)
					return vars
				}(),
			},
			want: want{
				err: errs.ErrExternalDependencyFailed,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewAlexSmith().Hash(),
						"123",
					)
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "not found and unexpected http response",
			s: impl.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithGet(
						func(_ string) (*http.Response, error) {
							return &http.Response{
								StatusCode:    http.StatusNotFound,
								Body:          http.NoBody,
								ContentLength: 0,
							}, nil
						},
					),
				),
				"http://localhost",
				userfx.NewAlexSmith(),
				body.NewBody(),
				[]event.Event{
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
				},
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewAlexSmith().Hash(),
						"123",
					)
					return vars
				}(),
			},
			want: want{
				err: errs.ErrExpectationFailed,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewAlexSmith().Hash(),
						"123",
					)
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "more than expected",
			s: impl.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithGet(
						func(_ string) (*http.Response, error) {
							v := `{` +
								`"events": [` +
								`{` +
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
								`},` +
								`{` +
								`"id": "2",` +
								`"title": "test title",` +
								`"description": "test description",` +
								`"location": {` +
								`"address": "test location"` +
								`},` +
								`"created_at": "2024-01-01T00:00:00Z",` +
								`"created_by": "test_user",` +
								`"started_at": "2024-01-01T01:00:00Z",` +
								`"finished_at": "2024-01-01T02:00:00Z"` +
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
				"http://localhost",
				userfx.NewAlexSmith(),
				body.NewBody(),
				[]event.Event{
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
				},
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewAlexSmith().Hash(),
						"123",
					)
					return vars
				}(),
			},
			want: want{
				err: errs.ErrExpectationFailed,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewAlexSmith().Hash(),
						"123",
					)
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "incorrect count field",
			s: impl.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithGet(
						func(_ string) (*http.Response, error) {
							v := `{` +
								`"events": [` +
								`{` +
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
								`}` +
								`],` +
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
				"http://localhost",
				userfx.NewAlexSmith(),
				body.NewBody(),
				[]event.Event{
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
				},
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewAlexSmith().Hash(),
						"123",
					)
					return vars
				}(),
			},
			want: want{
				err: errs.ErrExpectationFailed,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewAlexSmith().Hash(),
						"123",
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
