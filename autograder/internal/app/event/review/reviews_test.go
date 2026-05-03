package review_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/sitnikovik/ndbx/autograder/internal/app/event/review"
	common "github.com/sitnikovik/ndbx/autograder/internal/app/review/count"
)

func TestReviews_Counts(t *testing.T) {
	t.Parallel()

	type want struct {
		val common.Counts
	}
	tests := []struct {
		name string
		r    impl.Reviews
		want want
	}{
		{
			name: "with rating and counts",
			r: impl.NewReviews(
				impl.WithCounts(
					common.NewCounts(
						common.WithRating(4.8),
						common.WithCount(3),
					),
				),
			),
			want: want{
				val: common.NewCounts(
					common.WithRating(4.8),
					common.WithCount(3),
				),
			},
		},
		{
			name: "with rating only",
			r: impl.NewReviews(
				impl.WithCounts(
					common.NewCounts(
						common.WithRating(4.8),
					),
				),
			),
			want: want{
				val: common.NewCounts(
					common.WithRating(4.8),
				),
			},
		},
		{
			name: "with count only",
			r: impl.NewReviews(
				impl.WithCounts(
					common.NewCounts(
						common.WithCount(3),
					),
				),
			),
			want: want{
				val: common.NewCounts(
					common.WithCount(3),
				),
			},
		},
		{
			name: "default value",
			r:    impl.Reviews{},
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

func TestReviews_With(t *testing.T) {
	t.Parallel()
	type args struct {
		opts []impl.Option
	}
	type want struct {
		val impl.Reviews
	}
	tests := []struct {
		name string
		r    impl.Reviews
		args args
		want want
	}{
		{
			name: "with rating and count",
			r: impl.NewReviews(
				impl.WithCounts(
					common.NewCounts(
						common.WithRating(4.8),
						common.WithCount(3),
					),
				),
			),
			args: args{
				opts: []impl.Option{
					impl.WithCounts(
						common.NewCounts(
							common.WithRating(3.2),
							common.WithCount(1),
						),
					),
				},
			},
			want: want{
				val: impl.NewReviews(
					impl.WithCounts(
						common.NewCounts(
							common.WithRating(3.2),
							common.WithCount(1),
						),
					),
				),
			},
		},
		{
			name: "without opts",
			r: impl.NewReviews(
				impl.WithCounts(
					common.NewCounts(
						common.WithRating(4.8),
						common.WithCount(3),
					),
				),
			),
			args: args{
				opts: []impl.Option{},
			},
			want: want{
				val: impl.NewReviews(
					impl.WithCounts(
						common.NewCounts(
							common.WithRating(4.8),
							common.WithCount(3),
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
