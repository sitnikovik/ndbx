package reaction_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cassandra"
	"github.com/sitnikovik/ndbx/autograder/internal/app/cassandra/event/reaction"
	"github.com/sitnikovik/ndbx/autograder/internal/app/cassandra/event/reaction/filter"
	"github.com/sitnikovik/ndbx/autograder/internal/app/creation"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	common "github.com/sitnikovik/ndbx/autograder/internal/app/reaction"
	eventReaction "github.com/sitnikovik/ndbx/autograder/internal/app/reaction/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	qb "github.com/sitnikovik/ndbx/autograder/internal/client/cassandra/query/builder"
	cassandrafk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/cassandra"
	dbfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/cassandra/client"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestLikes_Select(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx context.Context
	}
	type want struct {
		val     []eventReaction.Like
		errored bool
	}
	tests := []struct {
		name string
		l    *reaction.Likes
		args args
		want want
	}{
		{
			name: "ok",
			l: reaction.NewLikes(
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
										"event_id",
										"like",
										"created_at",
										"created_by",
									},
									[]any{
										"123",
										int8(1),
										timex.MustRFC3339("2025-03-01T12:00:00Z"),
										"123213213",
									},
								),
							), nil
						},
					),
				),
				reaction.WithFilter(
					filter.NewFilter(
						qb.NewWhere(),
						filter.WithEventID(
							event.NewID("123"),
						),
						filter.WithCreatedBy(
							user.NewID("123213213"),
						),
					),
				),
				reaction.WithLimit(1),
			),
			args: args{
				ctx: context.Background(),
			},
			want: want{
				val: []eventReaction.Like{
					eventReaction.NewLike(
						common.NewLike(
							creation.NewStamp(
								creation.NewCreated(
									timex.MustRFC3339("2025-03-01T12:00:00Z"),
									user.NewIdentity(
										user.NewID("123213213"),
									),
								),
							),
						),
						eventReaction.NewEvent(
							event.NewID("123"),
						),
					),
				},
				errored: false,
			},
		},
		{
			name: "failed to select",
			l: reaction.NewLikes(
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
				reaction.WithFilter(
					filter.NewFilter(
						qb.NewWhere(),
						filter.WithEventID(
							event.NewID("123"),
						),
						filter.WithCreatedBy(
							user.NewID("123213213"),
						),
					),
				),
				reaction.WithLimit(1),
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
