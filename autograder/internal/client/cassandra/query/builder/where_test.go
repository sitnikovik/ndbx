package builder_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/sitnikovik/ndbx/autograder/internal/client/cassandra/query/builder"
)

func TestWhere_String(t *testing.T) {
	t.Parallel()
	type want struct {
		val   string
		panic bool
	}
	tests := []struct {
		name string
		w    *impl.Where
		want want
	}{
		{
			name: "several fields",
			w: func() *impl.Where {
				w := impl.NewWhere()
				w.Add("id", "1")
				w.Add("name", "John Doe")
				return w
			}(),
			want: want{
				val:   "WHERE id = ? AND name = ?",
				panic: false,
			},
		},
		{
			name: "one field",
			w: func() *impl.Where {
				w := impl.NewWhere()
				w.Add("id", "1")
				return w
			}(),
			want: want{
				val:   "WHERE id = ?",
				panic: false,
			},
		},
		{
			name: "no fields",
			w:    impl.NewWhere(),
			want: want{
				val:   "",
				panic: false,
			},
		},
		{
			name: "default value",
			w:    nil,
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
					_ = tt.w.String()
				})
				return
			}
			assert.Equal(
				t,
				tt.want.val,
				tt.w.String(),
			)
		})
	}
}

func TestWhere_Args(t *testing.T) {
	t.Parallel()
	type want struct {
		val   []any
		panic bool
	}
	tests := []struct {
		name string
		w    *impl.Where
		want want
	}{
		{
			name: "several fields",
			w: func() *impl.Where {
				w := impl.NewWhere()
				w.Add("id", "1")
				w.Add("name", "John Doe")
				return w
			}(),
			want: want{
				val:   []any{"1", "John Doe"},
				panic: false,
			},
		},
		{
			name: "one field",
			w: func() *impl.Where {
				w := impl.NewWhere()
				w.Add("id", "1")
				return w
			}(),
			want: want{
				val:   []any{"1"},
				panic: false,
			},
		},
		{
			name: "no fields",
			w:    impl.NewWhere(),
			want: want{
				val:   []any{},
				panic: false,
			},
		},
		{
			name: "default value",
			w:    nil,
			want: want{
				val:   nil,
				panic: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = tt.w.Args()
				})
				return
			}
			assert.Equal(
				t,
				tt.want.val,
				tt.w.Args(),
			)
		})
	}
}

func TestWhere_Count(t *testing.T) {
	t.Parallel()
	type want struct {
		val   int
		panic bool
	}
	tests := []struct {
		name string
		w    *impl.Where
		want want
	}{
		{
			name: "several fields",
			w: func() *impl.Where {
				w := impl.NewWhere()
				w.Add("id", "1")
				w.Add("name", "John Doe")
				return w
			}(),
			want: want{
				val:   2,
				panic: false,
			},
		},
		{
			name: "one field",
			w: func() *impl.Where {
				w := impl.NewWhere()
				w.Add("id", "1")
				return w
			}(),
			want: want{
				val:   1,
				panic: false,
			},
		},
		{
			name: "no fields",
			w:    impl.NewWhere(),
			want: want{
				val:   0,
				panic: false,
			},
		},
		{
			name: "default value",
			w:    nil,
			want: want{
				val:   0,
				panic: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = tt.w.Count()
				})
				return
			}
			assert.Equal(
				t,
				tt.want.val,
				tt.w.Count(),
			)
		})
	}
}
