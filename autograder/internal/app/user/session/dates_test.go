package session_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user/session"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestDates_CreatedAt(t *testing.T) {
	t.Parallel()
	type want struct {
		val time.Time
	}
	tests := []struct {
		name string
		d    session.Dates
		want want
	}{
		{
			name: "ok",
			d: session.NewDates(
				timex.MustParse(time.RFC3339, "2023-01-01T00:00:00Z"),
				timex.MustParse(time.RFC3339, "2023-01-02T00:00:00Z"),
			),
			want: want{
				val: timex.MustParse(time.RFC3339, "2023-01-01T00:00:00Z"),
			},
		},
		{
			name: "created_at is default value",
			d: session.NewDates(
				time.Time{},
				timex.MustParse(time.RFC3339, "2023-01-02T00:00:00Z"),
			),
			want: want{
				val: time.Time{},
			},
		},
		{
			name: "all default values",
			d: session.NewDates(
				time.Time{},
				time.Time{},
			),
			want: want{
				val: time.Time{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.d.CreatedAt(),
			)
		})
	}
}

func TestDates_UpdatedAt(t *testing.T) {
	t.Parallel()
	type want struct {
		val time.Time
	}
	tests := []struct {
		name string
		d    session.Dates
		want want
	}{
		{
			name: "ok",
			d: session.NewDates(
				timex.MustParse(time.RFC3339, "2023-01-01T00:00:00Z"),
				timex.MustParse(time.RFC3339, "2023-01-02T00:00:00Z"),
			),
			want: want{
				val: timex.MustParse(time.RFC3339, "2023-01-02T00:00:00Z"),
			},
		},
		{
			name: "updated_at is default value",
			d: session.NewDates(
				timex.MustParse(time.RFC3339, "2023-01-01T00:00:00Z"),
				time.Time{},
			),
			want: want{
				val: time.Time{},
			},
		},
		{
			name: "all default values",
			d: session.NewDates(
				time.Time{},
				time.Time{},
			),
			want: want{
				val: time.Time{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.d.UpdatedAt(),
			)
		})
	}
}
