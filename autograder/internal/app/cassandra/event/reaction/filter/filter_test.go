package filter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cassandra/event/reaction/enum/like"
	impl "github.com/sitnikovik/ndbx/autograder/internal/app/cassandra/event/reaction/filter"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	qb "github.com/sitnikovik/ndbx/autograder/internal/client/cassandra/query/builder"
)

func TestFilter_Where(t *testing.T) {
	t.Parallel()
	type want struct {
		val   string
		panic bool
	}
	tests := []struct {
		name string
		f    impl.Filter
		want want
	}{
		{
			name: "all set",
			f: impl.NewFilter(
				qb.NewWhere(),
				impl.WithLike(like.Like),
				impl.WithEventID(
					event.NewID("123"),
				),
				impl.WithCreatedBy(
					user.NewID("654"),
				),
			),
			want: want{
				val:   "WHERE like = ? AND event_id = ? AND created_by = ?",
				panic: false,
			},
		},
		{
			name: "all set but like",
			f: impl.NewFilter(
				qb.NewWhere(),
				impl.WithEventID(
					event.NewID("123"),
				),
				impl.WithCreatedBy(
					user.NewID("654"),
				),
			),
			want: want{
				val:   "WHERE event_id = ? AND created_by = ?",
				panic: false,
			},
		},
		{
			name: "all set but event id",
			f: impl.NewFilter(
				qb.NewWhere(),
				impl.WithLike(like.Dislike),
				impl.WithCreatedBy(
					user.NewID("654"),
				),
			),
			want: want{
				val:   "WHERE like = ? AND created_by = ?",
				panic: false,
			},
		},
		{
			name: "all set but created by",
			f: impl.NewFilter(
				qb.NewWhere(),
				impl.WithLike(like.Unspecified),
				impl.WithEventID(
					event.NewID("123"),
				),
			),
			want: want{
				val:   "WHERE like = ? AND event_id = ?",
				panic: false,
			},
		},
		{
			name: "only like is set",
			f: impl.NewFilter(
				qb.NewWhere(),
				impl.WithLike(like.Unspecified),
			),
			want: want{
				val:   "WHERE like = ?",
				panic: false,
			},
		},
		{
			name: "only event id is set",
			f: impl.NewFilter(
				qb.NewWhere(),
				impl.WithEventID(
					event.NewID("123"),
				),
			),
			want: want{
				val:   "WHERE event_id = ?",
				panic: false,
			},
		},
		{
			name: "only event id is set",
			f: impl.NewFilter(
				qb.NewWhere(),
				impl.WithCreatedBy(
					user.NewID("123"),
				),
			),
			want: want{
				val:   "WHERE created_by = ?",
				panic: false,
			},
		},
		{
			name: "empty",
			f: impl.NewFilter(
				qb.NewWhere(),
			),
			want: want{
				val:   "",
				panic: false,
			},
		},
		{
			name: "default value",
			f:    impl.Filter{},
			want: want{
				val:   "",
				panic: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = tt.f.Where()
				})
				return
			}
			assert.Equal(
				t,
				tt.want.val,
				tt.f.Where(),
			)
		})
	}
}

func TestFilter_Empty(t *testing.T) {
	t.Parallel()
	type want struct {
		val bool
	}
	tests := []struct {
		name string
		f    impl.Filter
		want want
	}{
		{
			name: "all set",
			f: impl.NewFilter(
				qb.NewWhere(),
				impl.WithLike(like.Like),
				impl.WithEventID(
					event.NewID("123"),
				),
				impl.WithCreatedBy(
					user.NewID("654"),
				),
			),
			want: want{
				val: false,
			},
		},
		{
			name: "all set but like",
			f: impl.NewFilter(
				qb.NewWhere(),
				impl.WithEventID(
					event.NewID("123"),
				),
				impl.WithCreatedBy(
					user.NewID("654"),
				),
			),
			want: want{
				val: false,
			},
		},
		{
			name: "all set but event id",
			f: impl.NewFilter(
				qb.NewWhere(),
				impl.WithLike(like.Dislike),
				impl.WithCreatedBy(
					user.NewID("654"),
				),
			),
			want: want{
				val: false,
			},
		},
		{
			name: "all set but created by",
			f: impl.NewFilter(
				qb.NewWhere(),
				impl.WithLike(like.Unspecified),
				impl.WithEventID(
					event.NewID("123"),
				),
			),
			want: want{
				val: false,
			},
		},
		{
			name: "only like is set",
			f: impl.NewFilter(
				qb.NewWhere(),
				impl.WithLike(like.Unspecified),
			),
			want: want{
				val: false,
			},
		},
		{
			name: "only event id is set",
			f: impl.NewFilter(
				qb.NewWhere(),
				impl.WithEventID(
					event.NewID("123"),
				),
			),
			want: want{
				val: false,
			},
		},
		{
			name: "only event id is set",
			f: impl.NewFilter(
				qb.NewWhere(),
				impl.WithCreatedBy(
					user.NewID("123"),
				),
			),
			want: want{
				val: false,
			},
		},
		{
			name: "empty",
			f: impl.NewFilter(
				qb.NewWhere(),
			),
			want: want{
				val: true,
			},
		},
		{
			name: "default value",
			f:    impl.Filter{},
			want: want{
				val: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.f.Empty()
			if tt.want.val {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}

func TestFilter_Args(t *testing.T) {
	t.Parallel()
	t.Run("where called", func(t *testing.T) {
		t.Parallel()
		f := impl.NewFilter(
			qb.NewWhere(),
			impl.WithLike(like.Like),
			impl.WithEventID(
				event.NewID("123"),
			),
			impl.WithCreatedBy(
				user.NewID("654"),
			),
		)
		_ = f.Where()
		assert.Equal(
			t,
			[]any{like.Like, "123", "654"},
			f.Args(),
		)
	})
	t.Run("where not called", func(t *testing.T) {
		t.Parallel()
		f := impl.NewFilter(
			qb.NewWhere(),
			impl.WithLike(like.Like),
			impl.WithEventID(
				event.NewID("123"),
			),
			impl.WithCreatedBy(
				user.NewID("654"),
			),
		)
		assert.Equal(
			t,
			[]any{like.Like, "123", "654"},
			f.Args(),
		)
	})
	t.Run("empty", func(t *testing.T) {
		t.Parallel()
		f := impl.NewFilter(
			qb.NewWhere(),
		)
		assert.Equal(
			t,
			[]any{},
			f.Args(),
		)
	})
	t.Run("default value", func(t *testing.T) {
		t.Parallel()
		f := impl.Filter{}
		assert.Panics(t, func() {
			_ = f.Args()
		})
	})
}
