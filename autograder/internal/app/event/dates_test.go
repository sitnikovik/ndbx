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

func TestDates_Equals(t *testing.T) {
	t.Parallel()
	type args struct {
		other event.Dates
	}
	type want struct {
		value bool
	}
	tests := []struct {
		name string
		d    event.Dates
		args args
		want want
	}{
		{
			name: "same",
			d: event.NewDates(
				timex.MustRFC3339("2025-03-15T12:00:00Z"),
				timex.MustRFC3339("2025-03-15T13:00:00Z"),
			),
			args: args{
				other: event.NewDates(
					timex.MustRFC3339("2025-03-15T12:00:00Z"),
					timex.MustRFC3339("2025-03-15T13:00:00Z"),
				),
			},
			want: want{
				value: true,
			},
		},
		{
			name: "same but different timezones",
			d: event.NewDates(
				timex.MustRFC3339("2025-03-15T15:00:00+03:00"),
				timex.MustRFC3339("2025-03-15T16:00:00+03:00"),
			),
			args: args{
				other: event.NewDates(
					timex.MustRFC3339("2025-03-15T12:00:00Z"),
					timex.MustRFC3339("2025-03-15T13:00:00Z"),
				),
			},
			want: want{
				value: true,
			},
		},
		{
			name: "different start dates",
			d: event.NewDates(
				timex.MustRFC3339("2025-03-15T12:00:00Z"),
				timex.MustRFC3339("2025-03-15T15:00:00Z"),
			),
			args: args{
				other: event.NewDates(
					timex.MustRFC3339("2025-03-15T14:00:00Z"),
					timex.MustRFC3339("2025-03-15T15:00:00Z"),
				),
			},
			want: want{
				value: false,
			},
		},
		{
			name: "different finish dates",
			d: event.NewDates(
				timex.MustRFC3339("2025-03-15T12:00:00Z"),
				timex.MustRFC3339("2025-03-15T14:00:00Z"),
			),
			args: args{
				other: event.NewDates(
					timex.MustRFC3339("2025-03-15T12:00:00Z"),
					timex.MustRFC3339("2025-03-15T15:00:00Z"),
				),
			},
			want: want{
				value: false,
			},
		},
		{
			name: "zero dates",
			d: event.NewDates(
				time.Time{},
				time.Time{},
			),
			args: args{
				other: event.NewDates(
					time.Time{},
					time.Time{},
				),
			},
			want: want{
				value: true,
			},
		},
		{
			name: "default value",
			d:    event.Dates{},
			args: args{
				other: event.NewDates(
					timex.MustRFC3339("2025-03-15T12:00:00Z"),
					timex.MustRFC3339("2025-03-15T15:00:00Z"),
				),
			},
			want: want{
				value: false,
			},
		},
		{
			name: "default arg",
			d: event.NewDates(
				timex.MustRFC3339("2025-03-15T12:00:00Z"),
				timex.MustRFC3339("2025-03-15T15:00:00Z"),
			),
			args: args{
				other: event.Dates{},
			},
			want: want{
				value: false,
			},
		},
		{
			name: "default arg and value",
			d:    event.Dates{},
			args: args{
				other: event.Dates{},
			},
			want: want{
				value: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.d.Equals(tt.args.other)
			if tt.want.value {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}
