package event_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestDates_StartedAt(t *testing.T) {
	t.Parallel()
	type want struct {
		val time.Time
	}
	tests := []struct {
		name string
		d    event.Dates
		want want
	}{
		{
			name: "ok",
			d: event.NewDates(
				timex.MustParse(time.DateTime, "2025-02-01 11:00:00"),
				timex.MustParse(time.DateTime, "2025-02-01 13:00:00"),
			),
			want: want{
				val: timex.MustParse(time.DateTime, "2025-02-01 11:00:00"),
			},
		},
		{
			name: "on zero",
			d: event.NewDates(
				timex.MustParse(time.DateTime, "0001-01-01 00:00:00"),
				timex.MustParse(time.DateTime, "2025-02-01 13:00:00"),
			),
			want: want{
				val: timex.MustParse(time.DateTime, "0001-01-01 00:00:00"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.d.StartedAt(),
			)
		})
	}
}
func TestDates_FinishedAt(t *testing.T) {
	t.Parallel()
	type want struct {
		val time.Time
	}
	tests := []struct {
		name string
		d    event.Dates
		want want
	}{
		{
			name: "ok",
			d: event.NewDates(
				timex.MustParse(time.DateTime, "2025-02-01 11:00:00"),
				timex.MustParse(time.DateTime, "2025-02-01 13:00:00"),
			),
			want: want{
				val: timex.MustParse(time.DateTime, "2025-02-01 13:00:00"),
			},
		},
		{
			name: "on zero",
			d: event.NewDates(
				timex.MustParse(time.DateTime, "2025-02-01 13:00:00"),
				timex.MustParse(time.DateTime, "0001-01-01 00:00:00"),
			),
			want: want{
				val: timex.MustParse(time.DateTime, "0001-01-01 00:00:00"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.d.FinishedAt(),
			)
		})
	}
}
