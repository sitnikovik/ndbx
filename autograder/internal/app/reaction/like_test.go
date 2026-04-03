package reaction_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/creation"
	impl "github.com/sitnikovik/ndbx/autograder/internal/app/reaction"
	userfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestLike_Created(t *testing.T) {
	t.Parallel()
	type want struct {
		val creation.Created
	}
	tests := []struct {
		name string
		l    impl.Like
		want want
	}{
		{
			name: "ok",
			l: impl.NewLike(
				creation.NewStamp(
					creation.NewCreated(
						timex.MustRFC3339("2025-03-01T12:00:00Z"),
						userfx.NewJohnDoe().Idendity(),
					),
				),
			),
			want: want{
				val: creation.NewCreated(
					timex.MustRFC3339("2025-03-01T12:00:00Z"),
					userfx.NewJohnDoe().Idendity(),
				),
			},
		},
		{
			name: "default value",
			l:    impl.Like{},
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
				tt.l.Created(),
			)
		})
	}
}
