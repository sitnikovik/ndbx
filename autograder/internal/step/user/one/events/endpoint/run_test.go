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
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/user/one/events/endpoint"
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
				user.NewID("1"),
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
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				err:   nil,
				vars:  step.NewVariables(),
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
				user.NewID("1"),
				body.NewBody(),
				nil,
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				err:   nil,
				vars:  step.NewVariables(),
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
				user.NewID("1"),
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
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				err:   errs.ErrExternalDependencyFailed,
				vars:  step.NewVariables(),
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
				user.NewID("1"),
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
				user.NewID("1"),
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
				user.NewID("1"),
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
			name: "got unexpected event",
			s: impl.NewStep(
				httpfk.NewFakeClient(
					httpfk.WithGet(
						func(_ string) (*http.Response, error) {
							v := `{` +
								`"events": [` +
								`{` +
								`"id": "2",` +
								`"title": "That's not your's",` +
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
				user.NewID("1"),
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
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				err:   errs.ErrExpectationFailed,
				vars:  step.NewVariables(),
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
