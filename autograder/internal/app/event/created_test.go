package event_test

import (
	"testing"

	impl "github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
	"github.com/stretchr/testify/assert"
)

func TestCreated_Equals(t *testing.T) {
	t.Parallel()
	type want struct {
		value bool
	}
	tests := []struct {
		name  string
		c     impl.Created
		other impl.Created
		want  want
	}{
		{
			name: "same time and user",
			c: impl.NewCreated(
				timex.MustRFC3339("2025-03-15T12:00:00Z"),
				user.NewIdentity(
					user.NewID("1234"),
				),
			),
			other: impl.NewCreated(
				timex.MustRFC3339("2025-03-15T12:00:00Z"),
				user.NewIdentity(
					user.NewID("1234"),
				),
			),
			want: want{
				value: true,
			},
		},
		{
			name: "user not eq",
			c: impl.NewCreated(
				timex.MustRFC3339("2025-03-15T12:00:00Z"),
				user.NewIdentity(
					user.NewID("1234"),
				),
			),
			other: impl.NewCreated(
				timex.MustRFC3339("2025-03-15T12:00:00Z"),
				user.NewIdentity(
					user.NewID("1313"),
				),
			),
			want: want{
				value: false,
			},
		},
		{
			name: "time not eq",
			c: impl.NewCreated(
				timex.MustRFC3339("2025-03-15T13:00:00Z"),
				user.NewIdentity(
					user.NewID("1234"),
				),
			),
			other: impl.NewCreated(
				timex.MustRFC3339("2025-03-15T12:00:00Z"),
				user.NewIdentity(
					user.NewID("1234"),
				),
			),
			want: want{
				value: false,
			},
		},
		{
			name: "time in other timezone",
			c: impl.NewCreated(
				timex.MustRFC3339("2025-03-15T15:00:00+03:00"),
				user.NewIdentity(
					user.NewID("1234"),
				),
			),
			other: impl.NewCreated(
				timex.MustRFC3339("2025-03-15T12:00:00Z"),
				user.NewIdentity(
					user.NewID("1234"),
				),
			),
			want: want{
				value: true,
			},
		},
		{
			name: "empty user id",
			c: impl.NewCreated(
				timex.MustRFC3339("2025-03-15T12:00:00Z"),
				user.NewIdentity(
					user.NewID(""),
				),
			),
			other: impl.NewCreated(
				timex.MustRFC3339("2025-03-15T12:00:00Z"),
				user.NewIdentity(
					user.NewID(""),
				),
			),
			want: want{
				value: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.c.Equals(tt.other)
			if tt.want.value {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}
