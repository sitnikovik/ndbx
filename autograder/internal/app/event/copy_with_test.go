package event_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestEvent_CopyWith(t *testing.T) {
	t.Parallel()
	type args struct {
		opts []event.Option
	}
	type want struct {
		val event.Event
	}
	tests := []struct {
		name string
		e    event.Event
		args args
		want want
	}{
		{
			name: "with id",
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
				opts: []event.Option{
					event.WithID("123"),
				},
			},
			want: want{
				val: event.NewEvent(
					event.NewID("123"),
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
		},
		{
			name: "with user identity",
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
				opts: []event.Option{
					event.WithCreatedBy(
						user.NewIdentity(user.NewID("21823091i")),
					),
				},
			},
			want: want{
				val: event.NewEvent(
					event.NewID("1"),
					event.NewContent(
						"My birthday",
						"The best day of the year",
					),
					event.NewLocation("home"),
					event.NewCreated(
						timex.MustParse(time.RFC3339, "2024-01-01T00:00:00Z"),
						user.NewIdentity(user.NewID("21823091i")),
					),
					event.NewDates(
						timex.MustParse(time.RFC3339, "2024-01-07T00:00:00Z"),
						timex.MustParse(time.RFC3339, "2024-01-07T23:59:59Z"),
					),
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.e.CopyWith(tt.args.opts...),
			)
		})
	}
}
