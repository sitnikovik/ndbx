package creation_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/creation"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	userfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestCreated_At(t *testing.T) {
	t.Parallel()
	type want struct {
		val time.Time
	}
	tests := []struct {
		name string
		c    creation.Created
		want want
	}{
		{
			name: "ok",
			c: creation.NewCreated(
				timex.MustRFC3339("2025-03-01T12:00:00Z"),
				userfx.
					NewAlexSmith().
					Idendity(),
			),
			want: want{
				val: timex.MustRFC3339("2025-03-01T12:00:00Z"),
			},
		},
		{
			name: "default value at field",
			c: creation.NewCreated(
				time.Time{},
				userfx.
					NewAlexSmith().
					Idendity(),
			),
			want: want{
				val: time.Time{},
			},
		},
		{
			name: "default instance value",
			c:    creation.Created{},
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
				tt.c.At(),
			)
		})
	}
}

func TestCreated_By(t *testing.T) {
	t.Parallel()
	type want struct {
		val user.Identity
	}
	tests := []struct {
		name string
		c    creation.Created
		want want
	}{
		{
			name: "ok",
			c: creation.NewCreated(
				timex.MustRFC3339("2025-03-01T12:00:00Z"),
				userfx.
					NewAlexSmith().
					Idendity(),
			),
			want: want{
				userfx.
					NewAlexSmith().
					Idendity(),
			},
		},
		{
			name: "default value by field",
			c: creation.NewCreated(
				timex.MustRFC3339("2025-03-01T12:00:00Z"),
				user.Identity{},
			),
			want: want{
				val: user.Identity{},
			},
		},
		{
			name: "default instance value",
			c:    creation.Created{},
			want: want{
				val: user.Identity{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.c.By(),
			)
		})
	}
}
