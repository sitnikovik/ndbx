package reaction_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/sitnikovik/ndbx/autograder/internal/app/cassandra/event/reaction"
	"github.com/sitnikovik/ndbx/autograder/internal/app/creation"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	common "github.com/sitnikovik/ndbx/autograder/internal/app/reaction"
	eventreaction "github.com/sitnikovik/ndbx/autograder/internal/app/reaction/event"
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

func TestToLikes(t *testing.T) {
	t.Parallel()
	type args struct {
		iter scanner
	}
	type want struct {
		val []eventreaction.Like
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
					cassandrafk.NewRow(
						[]string{
							"event_id",
							"like",
							"created_at",
							"created_by",
						},
						[]any{
							"123",
							int8(0),
							timex.MustRFC3339("2025-03-01T12:12:00Z"),
							"wdp723uyeh",
						},
					),
					cassandrafk.NewRow(
						[]string{
							"event_id",
							"like",
							"created_at",
							"created_by",
						},
						[]any{
							"123",
							int8(16),
							timex.MustRFC3339("2025-03-01T12:12:00Z"),
							"wdp723uyeh",
						},
					),
					cassandrafk.NewRow(
						[]string{
							"event_id",
							"like",
							"created_at",
							"created_by",
						},
						[]any{
							"123",
							int8(-1),
							timex.MustRFC3339("2025-03-01T13:12:00Z"),
							"2u2397183wi9",
						},
					),
				),
			},
			want: want{
				val: []eventreaction.Like{
					eventreaction.NewLike(
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
						eventreaction.NewEvent(
							event.NewID("123"),
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
				[]eventreaction.Like{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				impl.ToLikes(tt.args.iter),
			)
		})
	}
}
