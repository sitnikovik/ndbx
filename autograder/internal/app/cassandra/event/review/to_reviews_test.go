package review_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/sitnikovik/ndbx/autograder/internal/app/cassandra/event/review"
	"github.com/sitnikovik/ndbx/autograder/internal/app/creation"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/rating"
	eventreview "github.com/sitnikovik/ndbx/autograder/internal/app/review/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	cassandrafk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/cassandra"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

// scanner implements the scanner
// that can be used to scan the likes.
type scanner interface {
	// Scan scans the next like and writes
	// the values to the provided variables.
	Scan(...any) bool
	// Close closes the scanner.
	Close() error
}

func TestToReviews(t *testing.T) {
	t.Parallel()
	type args struct {
		iter scanner
	}
	type want struct {
		val []eventreview.Review
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "ok",
			args: args{
				iter: cassandrafk.NewIter(
					cassandrafk.NewRow(
						[]string{
							"id",
							"event_id",
							"rate",
							"comment",
							"created_at",
							"created_by",
							"updated_at",
						},
						[]any{
							"123",
							"43224",
							int8(4),
							"Great!",
							timex.MustRFC3339("2025-03-01T12:00:00Z"),
							"123213213",
							timex.MustRFC3339("2025-03-03T14:00:00Z"),
						},
					),
				),
			},
			want: want{
				val: []eventreview.Review{
					eventreview.NewReview(
						"123",
						creation.NewStamp(
							creation.NewCreated(
								timex.MustRFC3339("2025-03-01T12:00:00Z"),
								user.NewIdentity(
									user.NewID("123213213"),
								),
							),
						),
						eventreview.NewEvent(
							event.NewID("43224"),
						),
						"Great!",
						rating.NewRating(4),
						eventreview.WithUpdatedAt(
							timex.MustRFC3339("2025-03-03T14:00:00Z"),
						),
					),
				},
			},
		},
		{
			name: "empty rows",
			args: args{
				iter: cassandrafk.NewIter(),
			},
			want: want{
				[]eventreview.Review{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				impl.ToReviews(tt.args.iter),
			)
		})
	}
}
