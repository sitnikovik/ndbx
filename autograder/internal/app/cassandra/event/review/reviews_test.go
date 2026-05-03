package review_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cassandra"
	impl "github.com/sitnikovik/ndbx/autograder/internal/app/cassandra/event/review"
	"github.com/sitnikovik/ndbx/autograder/internal/app/cassandra/event/review/filter"
	"github.com/sitnikovik/ndbx/autograder/internal/app/creation"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/rating"
	eventreview "github.com/sitnikovik/ndbx/autograder/internal/app/review/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	qb "github.com/sitnikovik/ndbx/autograder/internal/client/cassandra/query/builder"
	cassandrafk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/cassandra"
	dbfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/cassandra/client"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestReviews_Select(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx context.Context
	}
	type want struct {
		val     []eventreview.Review
		errored bool
	}
	tests := []struct {
		name string
		l    *impl.Reviews
		args args
		want want
	}{
		{
			name: "ok",
			l: impl.NewReviews(
				dbfk.NewClient(
					dbfk.WithSelect(
						func(
							_ context.Context,
							_ string,
							_ ...any,
						) (cassandra.Scanner, error) {
							return cassandrafk.NewIter(
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
							), nil
						},
					),
				),
				impl.WithLimit(1),
				impl.WithFilter(
					filter.NewFilter(
						qb.NewWhere(),
						filter.WithEventID(
							event.NewID("123"),
						),
					),
				),
			),
			args: args{
				ctx: context.Background(),
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
				errored: false,
			},
		},
		{
			name: "failed to select",
			l: impl.NewReviews(
				dbfk.NewClient(
					dbfk.WithSelect(
						func(
							_ context.Context,
							_ string,
							_ ...any,
						) (cassandra.Scanner, error) {
							return nil, assert.AnError
						},
					),
				),
				impl.WithLimit(1),
			),
			args: args{
				ctx: context.Background(),
			},
			want: want{
				val:     nil,
				errored: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.l.Select(tt.args.ctx)
			assert.Equal(
				t,
				tt.want.val,
				got,
			)
			if tt.want.errored {
				assert.Error(
					t,
					err,
					"unexpected error: %v",
					err,
				)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
