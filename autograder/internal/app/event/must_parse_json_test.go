package event_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestMustParseJSON(t *testing.T) {
	t.Parallel()
	type args struct {
		bb []byte
	}
	type want struct {
		val   event.Event
		panic bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "all fields",
			args: args{
				bb: []byte(`{` +
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
			},
			want: want{
				val: event.NewEvent(
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
				panic: false,
			},
		},
		{
			name: "only title and desc",
			args: args{
				bb: []byte(`{` +
					`"title": "test title",` +
					`"description": "test description"` +
					`}`,
				),
			},
			want: want{
				val: event.NewEvent(
					event.NewID(""),
					event.NewContent(
						"test title",
						"test description",
					),
					event.NewLocation(""),
					event.NewCreated(
						time.Time{},
						user.NewIdentity(""),
					),
					event.NewDates(
						time.Time{},
						time.Time{},
					),
				),
				panic: false,
			},
		},
		{
			name: "invalid time format",
			args: args{
				bb: []byte(`{` +
					`"id": "1",` +
					`"title": "test title",` +
					`"description": "test description",` +
					`"location": {` +
					`"address": "test location"` +
					`},` +
					`"created_at": "2024-01-01 00:00:00",` +
					`"created_by": "test_user",` +
					`"started_at": "2024-01-01 01:00:00",` +
					`"finished_at": "2024-01-01 02:00:00"` +
					`}`,
				),
			},
			want: want{
				val:   event.Event{},
				panic: true,
			},
		},
		{
			name: "invalid json",
			args: args{
				bb: []byte(`not json`),
			},
			want: want{
				val:   event.Event{},
				panic: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = event.MustParseJSON(tt.args.bb)
				})
				return
			}
			assert.Equal(
				t,
				tt.want.val,
				event.MustParseJSON(tt.args.bb),
			)
		})
	}
}
