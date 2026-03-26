package body_test

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/get/one/resp/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestMustParseBody(t *testing.T) {
	t.Parallel()
	type args struct {
		body io.ReadCloser
	}
	type want struct {
		val   body.Body
		panic bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "ok",
			args: args{
				body: io.NopCloser(
					strings.NewReader(`{` +
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
						`}`,
					),
				),
			},
			want: want{
				val: body.NewBody(
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
				panic: false,
			},
		},
		{
			name: "invalid json",
			args: args{
				body: io.NopCloser(
					strings.NewReader(`not json`),
				),
			},
			want: want{
				val:   body.NewBody(event.Event{}),
				panic: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = body.MustParseBody(tt.args.body)
				})
				return
			}
			assert.Equal(
				t,
				tt.want.val,
				body.MustParseBody(tt.args.body),
			)
		})
	}
}
