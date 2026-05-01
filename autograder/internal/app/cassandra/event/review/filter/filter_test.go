package filter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/sitnikovik/ndbx/autograder/internal/app/cassandra/event/review/filter"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	qb "github.com/sitnikovik/ndbx/autograder/internal/client/cassandra/query/builder"
)

func TestFilter_Empty(t *testing.T) {
	t.Parallel()
	type want struct {
		value bool
		panic bool
	}
	tests := []struct {
		name string
		f    *impl.Filter
		want want
	}{
		{
			name: "all set",
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
				value: false,
				panic: false,
			},
		},
		{
			name: "only created by",
			f: impl.NewFilter(
				qb.NewWhere(),
				impl.WithCreatedBy(
					user.NewID("654"),
				),
			),
			want: want{
				value: false,
				panic: false,
			},
		},
		{
			name: "only event id",
			f: impl.NewFilter(
				qb.NewWhere(),
				impl.WithEventID(
					event.NewID("123"),
				),
			),
			want: want{
				value: false,
				panic: false,
			},
		},
		{
			name: "empty",
			f: impl.NewFilter(
				qb.NewWhere(),
			),
			want: want{
				value: true,
				panic: false,
			},
		},
		{
			name: "default value",
			f:    nil,
			want: want{
				value: false,
				panic: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = tt.f.Empty()
				})
				return
			}
			got := tt.f.Empty()
			if tt.want.value {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}

func TestFilter_Where(t *testing.T) {
	t.Parallel()
	type want struct {
		value string
		panic bool
	}
	tests := []struct {
		name string
		f    *impl.Filter
		want want
	}{
		{
			name: "all set",
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
				value: "WHERE event_id = ? AND created_by = ?",
				panic: false,
			},
		},
		{
			name: "only created by",
			f: impl.NewFilter(
				qb.NewWhere(),
				impl.WithCreatedBy(
					user.NewID("654"),
				),
			),
			want: want{
				value: "WHERE created_by = ?",
				panic: false,
			},
		},
		{
			name: "only event id",
			f: impl.NewFilter(
				qb.NewWhere(),
				impl.WithEventID(
					event.NewID("123"),
				),
			),
			want: want{
				value: "WHERE event_id = ?",
				panic: false,
			},
		},
		{
			name: "empty",
			f: impl.NewFilter(
				qb.NewWhere(),
			),
			want: want{
				value: "",
				panic: false,
			},
		},
		{
			name: "default value",
			f:    nil,
			want: want{
				value: "",
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
				tt.want.value,
				tt.f.Where(),
			)
		})
	}
}

func TestFilter_Args(t *testing.T) {
	t.Parallel()
	t.Run("where called", func(t *testing.T) {
		t.Parallel()
		f := impl.NewFilter(
			qb.NewWhere(),
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
			[]any{"123", "654"},
			f.Args(),
		)
	})
	t.Run("where not called", func(t *testing.T) {
		t.Parallel()
		f := impl.NewFilter(
			qb.NewWhere(),
			impl.WithEventID(
				event.NewID("123"),
			),
			impl.WithCreatedBy(
				user.NewID("654"),
			),
		)
		assert.Equal(
			t,
			[]any{"123", "654"},
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
