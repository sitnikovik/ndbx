package session_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user/session"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestSession_String(t *testing.T) {
	t.Parallel()
	type want struct {
		val string
	}
	tests := []struct {
		name string
		s    session.Session
		want want
	}{
		{
			name: "ok",
			s: session.NewSession(
				session.NewID("1"),
				session.NewDates(
					timex.MustParse(time.RFC3339, "2024-06-01T00:00:00Z"),
					timex.MustParse(time.RFC3339, "2024-06-02T00:00:00Z"),
				),
			),
			want: want{
				val: "1",
			},
		},
		{
			name: "empty sid",
			s: session.NewSession(
				session.NewID(""),
				session.NewDates(
					timex.MustParse(time.RFC3339, "2024-06-01T00:00:00Z"),
					timex.MustParse(time.RFC3339, "2024-06-02T00:00:00Z"),
				),
			),
			want: want{
				val: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.s.String(),
			)
		})
	}
}

func TestSession_User(t *testing.T) {
	t.Parallel()
	type want struct {
		val session.User
	}
	tests := []struct {
		name string
		s    session.Session
		want want
	}{
		{
			name: "ok",
			s: session.NewSession(
				session.NewID("1"),
				session.NewDates(
					timex.MustParse(time.RFC3339, "2024-06-01T00:00:00Z"),
					timex.MustParse(time.RFC3339, "2024-06-02T00:00:00Z"),
				),
				session.WithUser(
					session.NewUser(
						user.NewID("2"),
					),
				),
			),
			want: want{
				val: session.NewUser(
					user.NewID("2"),
				),
			},
		},
		{
			name: "no user",
			s: session.NewSession(
				session.NewID("1"),
				session.NewDates(
					timex.MustParse(time.RFC3339, "2024-06-01T00:00:00Z"),
					timex.MustParse(time.RFC3339, "2024-06-02T00:00:00Z"),
				),
			),
			want: want{
				val: session.User{},
			},
		},
		{
			name: "default value",
			s:    session.Session{},
			want: want{
				val: session.User{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.s.User(),
			)
		})
	}
}
