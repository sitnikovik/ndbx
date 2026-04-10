package count_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/reaction/count"
)

func TestCounts_Empty(t *testing.T) {
	t.Parallel()
	type want struct {
		val bool
	}
	tests := []struct {
		name string
		c    count.Counts
		want want
	}{
		{
			name: "all set",
			c: count.NewCounts(
				count.WithLikes(24),
				count.WithDislikes(3),
			),
			want: want{
				val: false,
			},
		},
		{
			name: "only likes set",
			c: count.NewCounts(
				count.WithLikes(24),
			),
			want: want{
				val: false,
			},
		},
		{
			name: "only dislikes set",
			c: count.NewCounts(
				count.WithDislikes(3),
			),
			want: want{
				val: false,
			},
		},
		{
			name: "zeros",
			c: count.NewCounts(
				count.WithLikes(0),
				count.WithDislikes(0),
			),
			want: want{
				val: true,
			},
		},
		{
			name: "default value",
			c:    count.Counts{},
			want: want{
				val: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.c.Empty()
			if tt.want.val {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}

func TestCounts_Likes(t *testing.T) {
	t.Parallel()
	type want struct {
		val uint64
	}
	tests := []struct {
		name string
		c    count.Counts
		want want
	}{
		{
			name: "all set",
			c: count.NewCounts(
				count.WithLikes(24),
				count.WithDislikes(3),
			),
			want: want{
				val: 24,
			},
		},
		{
			name: "only likes set",
			c: count.NewCounts(
				count.WithLikes(24),
			),
			want: want{
				val: 24,
			},
		},
		{
			name: "only dislikes set",
			c: count.NewCounts(
				count.WithDislikes(3),
			),
			want: want{
				val: 0,
			},
		},
		{
			name: "zeros",
			c: count.NewCounts(
				count.WithLikes(0),
				count.WithDislikes(0),
			),
			want: want{
				val: 0,
			},
		},
		{
			name: "default value",
			c:    count.Counts{},
			want: want{
				val: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.c.Likes(),
			)
		})
	}
}

func TestCounts_Dislikes(t *testing.T) {
	t.Parallel()
	type want struct {
		val uint64
	}
	tests := []struct {
		name string
		c    count.Counts
		want want
	}{
		{
			name: "all set",
			c: count.NewCounts(
				count.WithLikes(24),
				count.WithDislikes(3),
			),
			want: want{
				val: 3,
			},
		},
		{
			name: "only likes set",
			c: count.NewCounts(
				count.WithLikes(24),
			),
			want: want{
				val: 0,
			},
		},
		{
			name: "only dislikes set",
			c: count.NewCounts(
				count.WithDislikes(3),
			),
			want: want{
				val: 3,
			},
		},
		{
			name: "zeros",
			c: count.NewCounts(
				count.WithLikes(0),
				count.WithDislikes(0),
			),
			want: want{
				val: 0,
			},
		},
		{
			name: "default value",
			c:    count.Counts{},
			want: want{
				val: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.c.Dislikes(),
			)
		})
	}
}

func TestCounts_With(t *testing.T) {
	t.Parallel()
	type args struct {
		opts []count.Option
	}
	type want struct {
		val count.Counts
	}
	tests := []struct {
		name string
		c    count.Counts
		args args
		want want
	}{
		{
			name: "with new likes",
			c: count.NewCounts(
				count.WithLikes(24),
				count.WithDislikes(3),
			),
			args: args{
				opts: []count.Option{
					count.WithLikes(1321),
				},
			},
			want: want{
				val: count.NewCounts(
					count.WithLikes(1321),
					count.WithDislikes(3),
				),
			},
		},
		{
			name: "with new dislikes",
			c: count.NewCounts(
				count.WithLikes(24),
				count.WithDislikes(3),
			),
			args: args{
				opts: []count.Option{
					count.WithDislikes(1321),
				},
			},
			want: want{
				val: count.NewCounts(
					count.WithLikes(24),
					count.WithDislikes(1321),
				),
			},
		},
		{
			name: "without opts",
			c: count.NewCounts(
				count.WithLikes(24),
				count.WithDislikes(3),
			),
			args: args{
				opts: []count.Option{},
			},
			want: want{
				val: count.NewCounts(
					count.WithLikes(24),
					count.WithDislikes(3),
				),
			},
		},
		{
			name: "with opts but to default value",
			c:    count.Counts{},
			args: args{
				opts: []count.Option{
					count.WithLikes(1321),
					count.WithDislikes(3),
				},
			},
			want: want{
				val: count.NewCounts(
					count.WithLikes(1321),
					count.WithDislikes(3),
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.c.With(
				tt.args.opts...,
			)
			assert.Equal(
				t,
				tt.want.val,
				got,
			)
		})
	}
}
