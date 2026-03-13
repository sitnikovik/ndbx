package event_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestEvent_Equals(t *testing.T) {
	t.Parallel()
	type args struct {
		other event.Event
	}
	type want struct {
		ok bool
	}
	tests := []struct {
		name string
		e    event.Event
		args args
		want want
	}{
		{
			name: "equal events",
			e: event.NewEvent(
				event.NewID("1"),
				event.NewContent(
					"My birthday",
					"The best day of the year",
				),
				event.NewLocation("home"),
				event.NewCreated(
					timex.MustParse(time.RFC3339, "2024-01-01T00:00:00Z"),
					user.NewIdentity(user.NewID("123")),
				),
				event.NewDates(
					timex.MustParse(time.RFC3339, "2024-01-07T00:00:00Z"),
					timex.MustParse(time.RFC3339, "2024-01-07T23:59:59Z"),
				),
			),
			args: args{
				other: event.NewEvent(
					event.NewID("1"),
					event.NewContent(
						"My birthday",
						"The best day of the year",
					),
					event.NewLocation("home"),
					event.NewCreated(
						timex.MustParse(time.RFC3339, "2024-01-01T00:00:00Z"),
						user.NewIdentity(user.NewID("123")),
					),
					event.NewDates(
						timex.MustParse(time.RFC3339, "2024-01-07T00:00:00Z"),
						timex.MustParse(time.RFC3339, "2024-01-07T23:59:59Z"),
					),
				),
			},
			want: want{
				ok: true,
			},
		},
		{
			name: "different events by ID",
			e: event.NewEvent(
				event.NewID("1"),
				event.NewContent(
					"My birthday",
					"The best day of the year",
				),
				event.NewLocation("home"),
				event.NewCreated(
					timex.MustParse(time.RFC3339, "2024-01-01T00:00:00Z"),
					user.NewIdentity(user.NewID("123")),
				),
				event.NewDates(
					timex.MustParse(time.RFC3339, "2024-01-07T00:00:00Z"),
					timex.MustParse(time.RFC3339, "2024-01-07T23:59:59Z"),
				),
			),
			args: args{
				other: event.NewEvent(
					event.NewID("2"),
					event.NewContent(
						"Meeting",
						"Project discussion",
					),
					event.NewLocation("office"),
					event.NewCreated(
						timex.MustParse(time.RFC3339, "2024-01-02T00:00:00Z"),
						user.NewIdentity(user.NewID("456")),
					),
					event.NewDates(
						timex.MustParse(time.RFC3339, "2024-01-08T10:00:00Z"),
						timex.MustParse(time.RFC3339, "2024-01-08T11:00:00Z"),
					),
				),
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "different events by content",
			e: event.NewEvent(
				event.NewID("1"),
				event.NewContent(
					"My birthday party",
					"The best day of the year",
				),
				event.NewLocation("home"),
				event.NewCreated(
					timex.MustParse(time.RFC3339, "2024-01-01T00:00:00Z"),
					user.NewIdentity(user.NewID("123")),
				),
				event.NewDates(
					timex.MustParse(time.RFC3339, "2024-01-07T00:00:00Z"),
					timex.MustParse(time.RFC3339, "2024-01-07T23:59:59Z"),
				),
			),
			args: args{
				other: event.NewEvent(
					event.NewID("1"),
					event.NewContent(
						"My birthday",
						"The best day of the year",
					),
					event.NewLocation("home"),
					event.NewCreated(
						timex.MustParse(time.RFC3339, "2024-01-01T00:00:00Z"),
						user.NewIdentity(user.NewID("123")),
					),
					event.NewDates(
						timex.MustParse(time.RFC3339, "2024-01-07T00:00:00Z"),
						timex.MustParse(time.RFC3339, "2024-01-07T23:59:59Z"),
					),
				),
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "different events by location",
			e: event.NewEvent(
				event.NewID("1"),
				event.NewContent(
					"My birthday",
					"The best day of the year",
				),
				event.NewLocation("home"),
				event.NewCreated(
					timex.MustParse(time.RFC3339, "2024-01-01T00:00:00Z"),
					user.NewIdentity(user.NewID("123")),
				),
				event.NewDates(
					timex.MustParse(time.RFC3339, "2024-01-07T00:00:00Z"),
					timex.MustParse(time.RFC3339, "2024-01-07T23:59:59Z"),
				),
			),
			args: args{
				other: event.NewEvent(
					event.NewID("1"),
					event.NewContent(
						"My birthday",
						"The best day of the year",
					),
					event.NewLocation("office"),
					event.NewCreated(
						timex.MustParse(time.RFC3339, "2024-01-01T00:00:00Z"),
						user.NewIdentity(user.NewID("123")),
					),
					event.NewDates(
						timex.MustParse(time.RFC3339, "2024-01-07T00:00:00Z"),
						timex.MustParse(time.RFC3339, "2024-01-07T23:59:59Z"),
					),
				),
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "different events by created",
			e: event.NewEvent(
				event.NewID("1"),
				event.NewContent(
					"My birthday",
					"The best day of the year",
				),
				event.NewLocation("home"),
				event.NewCreated(
					timex.MustParse(time.RFC3339, "2024-01-01T00:00:00Z"),
					user.NewIdentity(user.NewID("123")),
				),
				event.NewDates(
					timex.MustParse(time.RFC3339, "2024-01-07T00:00:00Z"),
					timex.MustParse(time.RFC3339, "2024-01-07T23:59:59Z"),
				),
			),
			args: args{
				other: event.NewEvent(
					event.NewID("1"),
					event.NewContent(
						"My birthday",
						"The best day of the year",
					),
					event.NewLocation("home"),
					event.NewCreated(
						timex.MustParse(time.RFC3339, "2024-01-02T00:00:00Z"),
						user.NewIdentity(user.NewID("123")),
					),
					event.NewDates(
						timex.MustParse(time.RFC3339, "2024-01-07T00:00:00Z"),
						timex.MustParse(time.RFC3339, "2024-01-07T23:59:59Z"),
					),
				),
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "different events by dates",
			e: event.NewEvent(
				event.NewID("1"),
				event.NewContent(
					"My birthday",
					"The best day of the year",
				),
				event.NewLocation("home"),
				event.NewCreated(
					timex.MustParse(time.RFC3339, "2024-01-01T00:00:00Z"),
					user.NewIdentity(user.NewID("123")),
				),
				event.NewDates(
					timex.MustParse(time.RFC3339, "2024-01-07T00:00:00Z"),
					timex.MustParse(time.RFC3339, "2024-01-07T23:59:59Z"),
				),
			),
			args: args{
				other: event.NewEvent(
					event.NewID("1"),
					event.NewContent(
						"My birthday",
						"The best day of the year",
					),
					event.NewLocation("home"),
					event.NewCreated(
						timex.MustParse(time.RFC3339, "2024-01-01T00:00:00Z"),
						user.NewIdentity(user.NewID("123")),
					),
					event.NewDates(
						timex.MustParse(time.RFC3339, "2024-01-08T00:00:00Z"),
						timex.MustParse(time.RFC3339, "2024-01-08T23:59:59Z"),
					),
				),
			},
			want: want{
				ok: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.ok {
				assert.True(t, tt.e.Equals(tt.args.other))
			} else {
				assert.False(t, tt.e.Equals(tt.args.other))
			}
		})
	}
}
