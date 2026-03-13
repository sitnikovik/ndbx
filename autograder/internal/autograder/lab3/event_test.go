package lab3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestNewTestEvent(t *testing.T) {
	t.Parallel()
	type want struct {
		val event.Event
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "ok",
			want: want{
				val: event.NewEvent(
					event.NewID("000000000000000000000001"),
					event.NewContent("Title", "Description"),
					event.NewLocation("City, Country, Street, 123"),
					event.NewCreated(
						timex.MustRFC3339("2025-02-01T09:00:00Z"),
						user.NewIdentity(
							user.NewID("000000000000000000000123"),
						),
					),
					event.NewDates(
						timex.MustRFC3339("2025-02-01T11:00:00Z"),
						timex.MustRFC3339("2025-02-01T13:00:00Z"),
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
				lab3.NewTestEvent(),
			)
		})
	}
}
