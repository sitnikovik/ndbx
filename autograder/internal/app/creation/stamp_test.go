package creation_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/creation"
	userfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestStamp_Created(t *testing.T) {
	t.Parallel()
	type want struct {
		val creation.Created
	}
	tests := []struct {
		name string
		s    creation.Stamp
		want want
	}{
		{
			name: "ok",
			s: creation.NewStamp(
				creation.NewCreated(
					timex.MustRFC3339("2025-03-01T12:00:00Z"),
					userfx.
						NewAlexSmith().
						Idendity(),
				),
			),
			want: want{
				val: creation.NewCreated(
					timex.MustRFC3339("2025-03-01T12:00:00Z"),
					userfx.
						NewAlexSmith().
						Idendity(),
				),
			},
		},
		{
			name: "default value created field",
			s: creation.NewStamp(
				creation.Created{},
			),
			want: want{
				val: creation.Created{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.s.Created(),
			)
		})
	}
}
