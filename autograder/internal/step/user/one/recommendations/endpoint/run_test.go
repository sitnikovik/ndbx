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
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/user/one/recommendations/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/step/user/one/recommendations/expect"
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
			name: "ok",
			s: impl.NewStep(
				NewDescFx(),
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
								`]` +
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
				"http://localhost:8080",
				userfx.NewJohnDoe(),
				expect.NewExpectations(
					expect.WithEvents(
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
					),
					expect.WithResponse(
						NewResponseXpctFx(),
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewJohnDoe().Hash(),
						"123",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewJohnDoe().Hash(),
						"123",
					)
					return vars
				}(),
				err:   nil,
				panic: false,
			},
		},
		{
			name: "no event hash in vars",
			s: impl.NewStep(
				NewDescFx(),
				httpfk.NewFakeClient(),
				"http://localhost:8080",
				userfx.NewJohnDoe(),
				expect.NewExpectations(
					expect.WithNoEvents(),
				),
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
			name: "http failed",
			s: impl.NewStep(
				NewDescFx(),
				httpfk.NewFakeClient(
					httpfk.WithGet(
						func(_ string) (*http.Response, error) {
							return nil, assert.AnError
						},
					),
				),
				"http://localhost:8080",
				userfx.NewJohnDoe(),
				expect.NewExpectations(
					expect.WithNoEvents(),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewJohnDoe().Hash(),
						"123",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewJohnDoe().Hash(),
						"123",
					)
					return vars
				}(),
				err:   errs.ErrHTTPFailed,
				panic: false,
			},
		},
		{
			name: "unexpected resp on not found",
			s: impl.NewStep(
				NewDescFx(),
				httpfk.NewFakeClient(
					httpfk.WithGet(
						func(_ string) (*http.Response, error) {
							v := `{` +
								`"message": "no events"` +
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
				"http://localhost:8080",
				userfx.NewJohnDoe(),
				expect.NewExpectations(
					expect.WithNoEvents(),
					expect.WithResponse(
						NewResponseXpctFx(),
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewJohnDoe().Hash(),
						"123",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewJohnDoe().Hash(),
						"123",
					)
					return vars
				}(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "unexpected events",
			s: impl.NewStep(
				NewDescFx(),
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
								`"title": "test title 2",` +
								`"description": "test description",` +
								`"location": {` +
								`"address": "test location"` +
								`},` +
								`"created_at": "2024-01-01T00:00:00Z",` +
								`"created_by": "test_user",` +
								`"started_at": "2024-01-01T01:00:00Z",` +
								`"finished_at": "2024-01-01T02:00:00Z"` +
								`}` +
								`]` +
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
				"http://localhost:8080",
				userfx.NewJohnDoe(),
				expect.NewExpectations(
					expect.WithEvents(
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
					),
					expect.WithResponse(
						NewResponseXpctFx(),
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewJohnDoe().Hash(),
						"123",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewJohnDoe().Hash(),
						"123",
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
