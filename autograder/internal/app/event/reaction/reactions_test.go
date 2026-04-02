package reaction_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event/reaction"
	common "github.com/sitnikovik/ndbx/autograder/internal/app/reaction/count"
)

func TestReactions_Counts(t *testing.T) {
	t.Parallel()

	type want struct {
		val common.Counts
	}
	tests := []struct {
		name string
		r    reaction.Reactions
		want want
	}{
		{
			name: "with likes and dislikes",
			r: reaction.NewReactions(
				reaction.WithCounts(
					common.NewCounts(
						common.WithLikes(24),
						common.WithDislikes(3),
					),
				),
			),
			want: want{
				val: common.NewCounts(
					common.WithLikes(24),
					common.WithDislikes(3),
				),
			},
		},
		{
			name: "with likes only",
			r: reaction.NewReactions(
				reaction.WithCounts(
					common.NewCounts(
						common.WithLikes(24),
					),
				),
			),
			want: want{
				val: common.NewCounts(
					common.WithLikes(24),
				),
			},
		},
		{
			name: "with dislikes only",
			r: reaction.NewReactions(
				reaction.WithCounts(
					common.NewCounts(
						common.WithDislikes(3),
					),
				),
			),
			want: want{
				val: common.NewCounts(
					common.WithDislikes(3),
				),
			},
		},
		{
			name: "default value",
			r:    reaction.Reactions{},
			want: want{
				val: common.NewCounts(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.r.Counts(),
			)
		})
	}
}

func TestReactions_With(t *testing.T) {
	t.Parallel()
	type args struct {
		opts []reaction.Option
	}
	type want struct {
		val reaction.Reactions
	}
	tests := []struct {
		name string
		r    reaction.Reactions
		args args
		want want
	}{
		{
			name: "with likes and dislikes",
			r: reaction.NewReactions(
				reaction.WithCounts(
					common.NewCounts(
						common.WithLikes(24),
						common.WithDislikes(3),
					),
				),
			),
			args: args{
				opts: []reaction.Option{
					reaction.WithCounts(
						common.NewCounts(
							common.WithLikes(32),
							common.WithDislikes(1),
						),
					),
				},
			},
			want: want{
				val: reaction.NewReactions(
					reaction.WithCounts(
						common.NewCounts(
							common.WithLikes(32),
							common.WithDislikes(1),
						),
					),
				),
			},
		},
		{
			name: "without opts",
			r: reaction.NewReactions(
				reaction.WithCounts(
					common.NewCounts(
						common.WithLikes(24),
						common.WithDislikes(3),
					),
				),
			),
			args: args{
				opts: []reaction.Option{},
			},
			want: want{
				val: reaction.NewReactions(
					reaction.WithCounts(
						common.NewCounts(
							common.WithLikes(24),
							common.WithDislikes(3),
						),
					),
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.r.With(
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
