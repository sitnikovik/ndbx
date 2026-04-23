package rating_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/sitnikovik/ndbx/autograder/internal/app/rating"
)

func TestAverage(t *testing.T) {
	t.Parallel()
	type args struct {
		rr []impl.Rating
	}
	type want struct {
		value impl.Rating
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "only consts",
			args: args{
				rr: []impl.Rating{
					impl.Five,
					impl.One,
					impl.Five,
					impl.Two,
				},
			},
			want: want{
				value: impl.NewRating(3.3),
			},
		},
		{
			name: "customs",
			args: args{
				rr: []impl.Rating{
					impl.NewRating(5),
					impl.NewRating(4.8),
					impl.NewRating(2.5),
					impl.NewRating(1.8),
					impl.NewRating(1.22131),
					impl.NewRating(3.5984),
				},
			},
			want: want{
				value: impl.NewRating(3.2),
			},
		},
		{
			name: "empty list",
			args: args{
				rr: []impl.Rating{},
			},
			want: want{
				value: impl.NewRating(0),
			},
		},
		{
			name: "one value",
			args: args{
				rr: []impl.Rating{
					impl.NewRating(3.5984),
				},
			},
			want: want{
				value: impl.NewRating(3.6),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				impl.Average(tt.args.rr...),
			)
		})
	}
}
