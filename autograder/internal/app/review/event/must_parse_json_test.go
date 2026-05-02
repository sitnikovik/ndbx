package event_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/creation"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/rating"
	impl "github.com/sitnikovik/ndbx/autograder/internal/app/review/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestMustParseJSON(t *testing.T) {
	t.Parallel()
	type args struct {
		bb []byte
	}
	type want struct {
		val   impl.Review
		panic bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "all fields",
			args: args{
				bb: []byte(`{` +
					`"id": "56e2c0b3a2b4c1a5e6f7f8b3",` +
					`"event_id": "12e9c0b1a2b3c3d5e6f7a8b7",` +
					`"comment": "Great!",` +
					`"created_at": "2026-03-14T14:59:32+03:00",` +
					`"created_by": "65e9c0b1a2b3c4d5e6f7a8b9",` +
					`"rating": 5,` +
					`"updated_at": "2026-03-14T14:59:32+03:00"` +
					`}`,
				),
			},
			want: want{
				val: impl.NewReview(
					"56e2c0b3a2b4c1a5e6f7f8b3",
					creation.NewStamp(
						creation.NewCreated(
							timex.MustRFC3339("2026-03-14T14:59:32+03:00"),
							user.NewIdentity(
								user.NewID("65e9c0b1a2b3c4d5e6f7a8b9"),
							),
						),
					),
					impl.NewEvent(
						event.NewID("12e9c0b1a2b3c3d5e6f7a8b7"),
					),
					"Great!",
					rating.NewRating(5),
					impl.WithUpdatedAt(
						timex.MustRFC3339("2026-03-14T14:59:32+03:00"),
					),
				),
				panic: false,
			},
		},
		{
			name: "only comment and id",
			args: args{
				bb: []byte(`{` +
					`"id": "56e2c0b3a2b4c1a5e6f7f8b3",` +
					`"comment": "Great!"` +
					`}`,
				),
			},
			want: want{
				val: impl.NewReview(
					"56e2c0b3a2b4c1a5e6f7f8b3",
					creation.NewStamp(
						creation.NewCreated(
							time.Time{},
							user.NewIdentity(
								user.NewID(""),
							),
						),
					),
					impl.NewEvent(
						event.NewID(""),
					),
					"Great!",
					rating.NewRating(0),
				),
				panic: false,
			},
		},
		{
			name: "invalid time format",
			args: args{
				bb: []byte(`{` +
					`"id": "56e2c0b3a2b4c1a5e6f7f8b3",` +
					`"event_id": "12e9c0b1a2b3c3d5e6f7a8b7",` +
					`"comment": "Great!",` +
					`"created_at": "2026-03-14 14:59:32",` +
					`"created_by": "65e9c0b1a2b3c4d5e6f7a8b9",` +
					`"rating": 5,` +
					`"updated_at": "2026-03-14T14:59:32+03:00"` +
					`}`,
				),
			},
			want: want{
				val:   impl.Review{},
				panic: true,
			},
		},
		{
			name: "invalid json",
			args: args{
				bb: []byte(`not json`),
			},
			want: want{
				val:   impl.Review{},
				panic: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = impl.MustParseJSON(tt.args.bb)
				})
				return
			}
			assert.Equal(
				t,
				tt.want.val,
				impl.MustParseJSON(tt.args.bb),
			)
		})
	}
}
