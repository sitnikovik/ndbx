package body_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/event/post/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestBody_MustBytes(t *testing.T) {
	t.Parallel()
	type want struct {
		val   []byte
		panic bool
	}
	tests := []struct {
		name string
		b    body.Body
		want want
	}{
		{
			name: "full",
			b: body.NewBody(
				event.NewEvent(
					event.NewID("1"),
					event.NewContent("Title", "Description"),
					event.NewLocation("City, Country, Street, 123"),
					event.NewCreated(
						time.Time{},
						user.NewIdentity(
							user.NewID("123"),
						),
					),
					event.NewDates(
						timex.MustParse(time.RFC3339, "2025-02-01T11:00:00Z"),
						timex.MustParse(time.RFC3339, "2025-02-01T13:00:00Z"),
					),
					event.WithQuantity(
						event.NewQuantity(5, 10),
					),
				),
			),
			want: want{
				val: []byte(`{` +
					`"address":"City, Country, Street, 123",` +
					`"description":"Description",` +
					`"finished_at":"2025-02-01T13:00:00Z",` +
					`"max_attendees":10,` +
					`"min_attendees":5,` +
					`"started_at":"2025-02-01T11:00:00Z",` +
					`"title":"Title"` +
					`}`),
				panic: false,
			},
		},
		{
			name: "without optional fields",
			b: body.NewBody(
				event.NewEvent(
					event.NewID("1"),
					event.NewContent("Title", ""),
					event.NewLocation("City, Country, Street, 123"),
					event.NewCreated(
						time.Time{},
						user.NewIdentity(
							user.NewID("123"),
						),
					),
					event.NewDates(
						timex.MustParse(time.RFC3339, "2025-02-01T11:00:00Z"),
						timex.MustParse(time.RFC3339, "2025-02-01T13:00:00Z"),
					),
				),
			),
			want: want{
				val: []byte(`{` +
					`"address":"City, Country, Street, 123",` +
					`"finished_at":"2025-02-01T13:00:00Z",` +
					`"started_at":"2025-02-01T11:00:00Z",` +
					`"title":"Title"` +
					`}`),
				panic: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = tt.b.MustBytes()
				})
				return
			}
			assert.Equal(
				t,
				tt.want.val,
				tt.b.MustBytes(),
			)
		})
	}
}
