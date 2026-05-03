package count_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/rating"
	impl "github.com/sitnikovik/ndbx/autograder/internal/app/review/count"
)

func TestCounts_Empty(t *testing.T) {
	t.Parallel()
	type want struct {
		value bool
	}
	tests := []struct {
		name string
		c    impl.Counts
		want want
	}{
		{
			name: "all set",
			c: impl.NewCounts(
				impl.WithRating(4.3),
				impl.WithCount(3),
			),
			want: want{
				value: false,
			},
		},
		{
			name: "only rating set",
			c: impl.NewCounts(
				impl.WithRating(4.8),
			),
			want: want{
				value: false,
			},
		},
		{
			name: "only count set",
			c: impl.NewCounts(
				impl.WithCount(3),
			),
			want: want{
				value: false,
			},
		},
		{
			name: "zeros",
			c: impl.NewCounts(
				impl.WithRating(0),
				impl.WithCount(0),
			),
			want: want{
				value: true,
			},
		},
		{
			name: "default value",
			c:    impl.Counts{},
			want: want{
				value: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.c.Empty()
			if tt.want.value {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}

func TestCounts_Rating(t *testing.T) {
	t.Parallel()
	type want struct {
		value rating.Rating
	}
	tests := []struct {
		name string
		c    impl.Counts
		want want
	}{
		{
			name: "set",
			c: impl.NewCounts(
				impl.WithRating(4.8),
			),
			want: want{
				value: rating.NewRating(4.8),
			},
		},
		{
			name: "zero",
			c: impl.NewCounts(
				impl.WithRating(0),
			),
			want: want{
				value: rating.NewRating(0),
			},
		},
		{
			name: "default values",
			c:    impl.Counts{},
			want: want{
				value: rating.None,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.c.Rating(),
			)
		})
	}
}

func TestCounts_Count(t *testing.T) {
	t.Parallel()
	type want struct {
		value int
	}
	tests := []struct {
		name string
		c    impl.Counts
		want want
	}{
		{
			name: "set",
			c: impl.NewCounts(
				impl.WithCount(4),
			),
			want: want{
				value: 4,
			},
		},
		{
			name: "zero",
			c: impl.NewCounts(
				impl.WithRating(0),
			),
			want: want{
				value: 0,
			},
		},
		{
			name: "default values",
			c:    impl.Counts{},
			want: want{
				value: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.c.Count(),
			)
		})
	}
}
