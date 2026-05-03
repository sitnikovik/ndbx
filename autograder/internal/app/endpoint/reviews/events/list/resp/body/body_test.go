package body_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/creation"
	impl "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/reviews/events/list/resp/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/rating"
	review "github.com/sitnikovik/ndbx/autograder/internal/app/review/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestBody_Reviews(t *testing.T) {
	t.Parallel()
	type want struct {
		value []review.Review
	}
	tests := []struct {
		name string
		r    impl.Body
		want want
	}{
		{
			name: "ok",
			r: impl.NewBody(
				[]review.Review{
					review.NewReview(
						"56e2c0b3a2b4c1a5e6f7f8b3",
						creation.NewStamp(
							creation.NewCreated(
								timex.MustRFC3339("2026-03-14T14:59:32+03:00"),
								user.NewIdentity(
									user.NewID("65e9c0b1a2b3c4d5e6f7a8b9"),
								),
							),
						),
						review.NewEvent(
							event.NewID("12e9c0b1a2b3c3d5e6f7a8b7"),
						),
						"Great!",
						rating.NewRating(5),
						review.WithUpdatedAt(
							timex.MustRFC3339("2026-03-14T14:59:32+03:00"),
						),
					),
				},
				1,
			),
			want: want{
				value: []review.Review{
					review.NewReview(
						"56e2c0b3a2b4c1a5e6f7f8b3",
						creation.NewStamp(
							creation.NewCreated(
								timex.MustRFC3339("2026-03-14T14:59:32+03:00"),
								user.NewIdentity(
									user.NewID("65e9c0b1a2b3c4d5e6f7a8b9"),
								),
							),
						),
						review.NewEvent(
							event.NewID("12e9c0b1a2b3c3d5e6f7a8b7"),
						),
						"Great!",
						rating.NewRating(5),
						review.WithUpdatedAt(
							timex.MustRFC3339("2026-03-14T14:59:32+03:00"),
						),
					),
				},
			},
		},
		{
			name: "empty",
			r: impl.NewBody(
				[]review.Review{},
				1,
			),
			want: want{
				value: []review.Review{},
			},
		},
		{
			name: "default value",
			r:    impl.Body{},
			want: want{
				value: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.r.Reviews(),
			)
		})
	}
}

func TestBody_Count(t *testing.T) {
	t.Parallel()
	type want struct {
		value int
	}
	tests := []struct {
		name string
		r    impl.Body
		want want
	}{
		{
			name: "ok",
			r: impl.NewBody(
				[]review.Review{
					review.NewReview(
						"56e2c0b3a2b4c1a5e6f7f8b3",
						creation.NewStamp(
							creation.NewCreated(
								timex.MustRFC3339("2026-03-14T14:59:32+03:00"),
								user.NewIdentity(
									user.NewID("65e9c0b1a2b3c4d5e6f7a8b9"),
								),
							),
						),
						review.NewEvent(
							event.NewID("12e9c0b1a2b3c3d5e6f7a8b7"),
						),
						"Great!",
						rating.NewRating(5),
						review.WithUpdatedAt(
							timex.MustRFC3339("2026-03-14T14:59:32+03:00"),
						),
					),
				},
				1,
			),
			want: want{
				value: 1,
			},
		},
		{
			name: "empty",
			r: impl.NewBody(
				[]review.Review{},
				0,
			),
			want: want{
				value: 0,
			},
		},
		{
			name: "default value",
			r:    impl.Body{},
			want: want{
				value: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.r.Count(),
			)
		})
	}
}
