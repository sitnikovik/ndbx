package session_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user/session"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestSession_Updated(t *testing.T) {
	t.Parallel()
	type want struct {
		ok bool
	}
	tests := []struct {
		name string
		s    session.Session
		want want
	}{
		{
			name: "updated_at is after created_at",
			s: session.NewSession(
				session.NewID("1"),
				session.NewDates(
					timex.MustParse(time.RFC3339, "2024-06-01T00:00:00Z"),
					timex.MustParse(time.RFC3339, "2024-06-02T00:00:00Z"),
				),
			),
			want: want{
				ok: true,
			},
		},
		{
			name: "updated_at is equal to created_at",
			s: session.NewSession(
				session.NewID("1"),
				session.NewDates(
					timex.MustParse(time.RFC3339, "2024-06-01T00:00:00Z"),
					timex.MustParse(time.RFC3339, "2024-06-01T00:00:00Z"),
				),
			),
			want: want{
				ok: false,
			},
		},
		{
			name: "updated_at is before created_at",
			s: session.NewSession(
				session.NewID("1"),
				session.NewDates(
					timex.MustParse(time.RFC3339, "2024-06-02T00:00:00Z"),
					timex.MustParse(time.RFC3339, "2024-06-01T00:00:00Z"),
				),
			),
			want: want{
				ok: false,
			},
		},
		{
			name: "updated_at is zero",
			s: session.NewSession(
				session.NewID("1"),
				session.NewDates(
					timex.MustParse(time.RFC3339, "2024-06-01T00:00:00Z"),
					time.Time{},
				),
			),
			want: want{
				ok: false,
			},
		},
		{
			name: "updated_at is in another timezone",
			s: session.NewSession(
				session.NewID("1"),
				session.NewDates(
					timex.MustParse(time.RFC3339, "2024-06-01T00:00:00Z"),
					timex.MustParse(time.RFC3339, "2024-06-01T00:00:00+03:00"),
				),
			),
			want: want{
				ok: false,
			},
		},
		{
			name: "updated_at is in another timezone but before created_at",
			s: session.NewSession(
				session.NewID("1"),
				session.NewDates(
					timex.MustParse(time.RFC3339, "2024-06-01T00:00:00Z"),
					timex.MustParse(time.RFC3339, "2024-05-31T21:00:00-03:00"),
				),
			),
			want: want{
				ok: false,
			},
		},
		{
			name: "updated_at is in another timezone but after created_at",
			s: session.NewSession(
				session.NewID("1"),
				session.NewDates(
					timex.MustParse(time.RFC3339, "2024-06-01T00:00:00Z"),
					timex.MustParse(time.RFC3339, "2024-06-01T03:00:00+03:00"),
				),
			),
			want: want{
				ok: false,
			},
		},
		{
			name: "default value",
			s:    session.Session{},
			want: want{
				ok: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.s.Updated()
			if tt.want.ok {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}
