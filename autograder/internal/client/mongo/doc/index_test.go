package doc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
)

func TestIndex_HasAllFor(t *testing.T) {
	t.Parallel()
	type args struct {
		keys []string
	}
	type want struct {
		ok bool
	}
	tests := []struct {
		name string
		i    doc.Index
		args args
		want want
	}{
		{
			name: "index has exactly all keys",
			i:    doc.NewIndex("a", "b", "c"),
			args: args{
				keys: []string{"a", "b", "c"},
			},
			want: want{
				ok: true,
			},
		},
		{
			name: "index does not have exactly all keys",
			i:    doc.NewIndex("a", "b", "c"),
			args: args{
				keys: []string{"a", "b"},
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "different len",
			i:    doc.NewIndex("a", "b", "c"),
			args: args{
				keys: []string{"a", "b", "c", "d"},
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "different order",
			i:    doc.NewIndex("a", "b", "c"),
			args: args{
				keys: []string{"c", "b", "a"},
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "empty index and empty keys",
			i:    doc.NewIndex(),
			args: args{
				keys: []string{},
			},
			want: want{
				ok: true,
			},
		},
		{
			name: "empty index and non-empty keys",
			i:    doc.NewIndex(),
			args: args{
				keys: []string{"a"},
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "non-empty index and empty keys",
			i:    doc.NewIndex("a"),
			args: args{
				keys: []string{},
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "index has all keys but also has extra keys",
			i:    doc.NewIndex("a", "b", "c", "d"),
			args: args{
				keys: []string{"a", "b", "c"},
			},
			want: want{
				ok: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.i.HasAllFor(tt.args.keys...)
			if tt.want.ok {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}

func TestIndex_HasAnyOf(t *testing.T) {
	t.Parallel()
	type args struct {
		keys []string
	}
	type want struct {
		ok bool
	}
	tests := []struct {
		name string
		i    doc.Index
		args args
		want want
	}{
		{
			name: "index has all keys",
			i:    doc.NewIndex("a", "b", "c"),
			args: args{
				keys: []string{"a", "b", "c"},
			},
			want: want{
				ok: true,
			},
		},
		{
			name: "index has some keys",
			i:    doc.NewIndex("a", "b", "c"),
			args: args{
				keys: []string{"a", "d"},
			},
			want: want{
				ok: true,
			},
		},
		{
			name: "index has no keys",
			i:    doc.NewIndex("a", "b", "c"),
			args: args{
				keys: []string{"d", "e"},
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "empty index and empty keys",
			i:    doc.NewIndex(),
			args: args{
				keys: []string{},
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "empty index and non-empty keys",
			i:    doc.NewIndex(),
			args: args{
				keys: []string{"a"},
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "non-empty index and empty keys",
			i:    doc.NewIndex("a"),
			args: args{
				keys: []string{},
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "index has some keys and some extra keys",
			i:    doc.NewIndex("a", "b", "c", "d"),
			args: args{
				keys: []string{"a", "e"},
			},
			want: want{
				ok: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.i.HasAnyOf(tt.args.keys...)
			if tt.want.ok {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}

func TestIndex_Unique(t *testing.T) {
	t.Parallel()
	type want struct {
		ok bool
	}
	tests := []struct {
		name string
		i    doc.Index
		want want
	}{
		{
			name: "ok",
			i:    doc.NewUniqueIndex("a", "b", "c"),
			want: want{
				ok: true,
			},
		},
		{
			name: "not unique",
			i:    doc.NewIndex("a", "b", "c"),
			want: want{
				ok: false,
			},
		},
		{
			name: "empty index",
			i:    doc.NewIndex(),
			want: want{
				ok: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.i.Unique()
			if tt.want.ok {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}

func TestIndex_Empty(t *testing.T) {
	t.Parallel()
	type want struct {
		ok bool
	}
	tests := []struct {
		name string
		i    doc.Index
		want want
	}{
		{
			name: "empty index",
			i:    doc.NewIndex(),
			want: want{
				ok: true,
			},
		},
		{
			name: "non-empty index",
			i:    doc.NewIndex("a"),
			want: want{
				ok: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.i.Empty()
			if tt.want.ok {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}
