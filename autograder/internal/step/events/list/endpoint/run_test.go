package endpoint_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	rq "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/get/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event/reaction"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event/review"
	"github.com/sitnikovik/ndbx/autograder/internal/app/reaction/count"
	reviewCounts "github.com/sitnikovik/ndbx/autograder/internal/app/review/count"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/events/list/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/step/events/list/endpoint/expect"
	httpfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/httpx"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

var (
	fxEvent = event.NewEvent(
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
	)
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
			name: "not found but expected",
			s: impl.NewStep(
				step.NewDesc(
					"Test Step",
					"Test description",
				),
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
				rq.NewBody(),
				expect.NewExpectations(
					expect.WithEvents(
						fxEvent,
					),
				),
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
			name: "not found and not expected",
			s: impl.NewStep(
				step.NewDesc(
					"Test Step",
					"Test description",
				),
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
				rq.NewBody(),
				expect.NewExpectations(
					expect.WithEvents(),
				),
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
				step.NewDesc(
					"Test Step",
					"Test description",
				),
				httpfk.NewFakeClient(
					httpfk.WithGet(
						func(_ string) (*http.Response, error) {
							return nil, assert.AnError
						},
					),
				),
				"http://localhost",
				rq.NewBody(),
				expect.NewExpectations(
					expect.WithEvents(),
				),
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
			name: "unexpected http response",
			s: impl.NewStep(
				step.NewDesc(
					"Test Step",
					"Test description",
				),
				httpfk.NewFakeClient(
					httpfk.WithGet(
						func(_ string) (*http.Response, error) {
							return &http.Response{
								StatusCode:    http.StatusOK,
								Body:          http.NoBody,
								ContentLength: 0,
							}, nil
						},
					),
				),
				"http://localhost",
				rq.NewBody(),
				expect.NewExpectations(
					expect.WithEvents(),
				),
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
			name: "got events",
			s: impl.NewStep(
				step.NewDesc(
					"Test Step",
					"Test description",
				),
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
				rq.NewBody(),
				expect.NewExpectations(
					expect.WithEvents(
						fxEvent,
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				err: nil,
				vars: func() step.Variables {
					vars := step.NewVariables()
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "got events but count field is wrong",
			s: impl.NewStep(
				step.NewDesc(
					"Test Step",
					"Test description",
				),
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
								`"count": 2` +
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
				rq.NewBody(),
				expect.NewExpectations(
					expect.WithEvents(
						fxEvent,
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				err: errs.ErrExpectationFailed,
				vars: func() step.Variables {
					vars := step.NewVariables()
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "got events more than expected",
			s: impl.NewStep(
				step.NewDesc(
					"Test Step",
					"Test description",
				),
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
								`"count": 2` +
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
				rq.NewBody(),
				expect.NewExpectations(
					expect.WithEvents(
						fxEvent,
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				err: errs.ErrExpectationFailed,
				vars: func() step.Variables {
					vars := step.NewVariables()
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "got events with reactions",
			s: impl.NewStep(
				step.NewDesc(
					"Test Step",
					"Test description",
				),
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
				rq.NewBody(),
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
							event.WithLikes(24),
							event.WithDislikes(3),
						),
					),
					expect.WithReactions(
						reaction.NewReactions(
							reaction.WithCounts(
								count.NewCounts(
									count.WithLikes(24),
									count.WithDislikes(3),
								),
							),
						),
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				err: nil,
				vars: func() step.Variables {
					vars := step.NewVariables()
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "got events with reactions but fails expectations",
			s: impl.NewStep(
				step.NewDesc(
					"Test Step",
					"Test description",
				),
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
								`"likes": 12323,` +
								`"dislikes": 23` +
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
				rq.NewBody(),
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
							event.WithLikes(24),
							event.WithDislikes(3),
						),
					),
					expect.WithReactions(
						reaction.NewReactions(
							reaction.WithCounts(
								count.NewCounts(
									count.WithLikes(24),
									count.WithDislikes(3),
								),
							),
						),
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				err: errs.ErrExpectationFailed,
				vars: func() step.Variables {
					vars := step.NewVariables()
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "got events with reactions but expectations are set incorrectly",
			s: impl.NewStep(
				step.NewDesc(
					"Test Step",
					"Test description",
				),
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
								`"finished_at": "2024-01-01T02:00:00Z",` +
								`"reactions": {` +
								`"likes": 24,` +
								`"dislikes": 3` +
								`}` +
								`}` +
								`],` +
								`"count": 2` +
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
				rq.NewBody(),
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
							event.WithLikes(24),
							event.WithDislikes(3),
						),
						event.NewEvent(
							event.NewID("2"),
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
							event.WithLikes(24),
							event.WithDislikes(3),
						),
					),
					expect.WithReactions(
						reaction.NewReactions(
							reaction.WithCounts(
								count.NewCounts(
									count.WithLikes(24),
									count.WithDislikes(3),
								),
							),
						),
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				err: nil,
				vars: func() step.Variables {
					vars := step.NewVariables()
					return vars
				}(),
				panic: true,
			},
		},
		{
			name: "got events with reviews",
			s: impl.NewStep(
				step.NewDesc(
					"Test Step",
					"Test description",
				),
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
								`"reviews": {` +
								`"rating": 4.8,` +
								`"count": 12` +
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
				rq.NewBody(),
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
							event.WithReviews(
								review.NewReviews(
									review.WithCounts(
										reviewCounts.NewCounts(
											reviewCounts.WithRating(4.8),
											reviewCounts.WithCount(12),
										),
									),
								),
							),
						),
					),
					expect.WithReviews(
						review.NewReviews(
							review.WithCounts(
								reviewCounts.NewCounts(
									reviewCounts.WithRating(4.8),
									reviewCounts.WithCount(12),
								),
							),
						),
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				err: nil,
				vars: func() step.Variables {
					vars := step.NewVariables()
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "got events with reviews but fails expectations",
			s: impl.NewStep(
				step.NewDesc(
					"Test Step",
					"Test description",
				),
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
								`"reviews": {` +
								`"rating": 3.8,` +
								`"count": 12` +
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
				rq.NewBody(),
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
							event.WithReviews(
								review.NewReviews(
									review.WithCounts(
										reviewCounts.NewCounts(
											reviewCounts.WithRating(4.8),
											reviewCounts.WithCount(12),
										),
									),
								),
							),
						),
					),
					expect.WithReviews(
						review.NewReviews(
							review.WithCounts(
								reviewCounts.NewCounts(
									reviewCounts.WithRating(4.8),
									reviewCounts.WithCount(12),
								),
							),
						),
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				err: errs.ErrExpectationFailed,
				vars: func() step.Variables {
					vars := step.NewVariables()
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "got events with reviews but expectations are set incorrectly",
			s: impl.NewStep(
				step.NewDesc(
					"Test Step",
					"Test description",
				),
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
								`"reviews": {` +
								`"rating": 3.8,` +
								`"count": 12` +
								`}` +
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
								`"finished_at": "2024-01-01T02:00:00Z",` +
								`"reviews": {` +
								`"rating": 3.8,` +
								`"count": 12` +
								`}` +
								`}` +
								`],` +
								`"count": 2` +
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
				rq.NewBody(),
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
							event.WithReviews(
								review.NewReviews(
									review.WithCounts(
										reviewCounts.NewCounts(
											reviewCounts.WithRating(4.8),
											reviewCounts.WithCount(12),
										),
									),
								),
							),
						),
						event.NewEvent(
							event.NewID("2"),
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
							event.WithReviews(
								review.NewReviews(
									review.WithCounts(
										reviewCounts.NewCounts(
											reviewCounts.WithRating(4.8),
											reviewCounts.WithCount(12),
										),
									),
								),
							),
						),
					),
					expect.WithReviews(
						review.NewReviews(
							review.WithCounts(
								reviewCounts.NewCounts(
									reviewCounts.WithRating(4.8),
									reviewCounts.WithCount(12),
								),
							),
						),
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				err: nil,
				vars: func() step.Variables {
					vars := step.NewVariables()
					return vars
				}(),
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
