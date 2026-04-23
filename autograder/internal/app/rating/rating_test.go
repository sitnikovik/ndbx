package rating_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/sitnikovik/ndbx/autograder/internal/app/rating"
)

func TestRating_String(t *testing.T) {
	t.Parallel()
	type want struct {
		value string
		panic bool
	}
	tests := []struct {
		name string
		r    impl.Rating
		want want
	}{
		{
			name: "one",
			r:    impl.One,
			want: want{
				value: "1.0",
				panic: false,
			},
		},
		{
			name: "two",
			r:    impl.Two,
			want: want{
				value: "2.0",
				panic: false,
			},
		},
		{
			name: "three",
			r:    impl.Three,
			want: want{
				value: "3.0",
				panic: false,
			},
		},
		{
			name: "four",
			r:    impl.Four,
			want: want{
				value: "4.0",
				panic: false,
			},
		},
		{
			name: "five",
			r:    impl.Five,
			want: want{
				value: "5.0",
				panic: false,
			},
		},
		{
			name: "float ceil",
			r:    impl.NewRating(3.987656789),
			want: want{
				value: "4.0",
				panic: false,
			},
		},
		{
			name: "float floor",
			r:    impl.NewRating(3.937656789),
			want: want{
				value: "3.9",
				panic: false,
			},
		},
		{
			name: "float more than five",
			r:    impl.NewRating(5.00001),
			want: want{
				value: "",
				panic: true,
			},
		},
		{
			name: "float lower than one",
			r:    impl.NewRating(0.9999),
			want: want{
				value: "",
				panic: true,
			},
		},
		{
			name: "zero",
			r:    impl.NewRating(0),
			want: want{
				value: "0.0",
				panic: false,
			},
		},
		{
			name: "six",
			r:    impl.NewRating(6),
			want: want{
				value: "",
				panic: true,
			},
		},
		{
			name: "negative",
			r:    impl.NewRating(-1),
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
					_ = tt.r.String()
				})
				return
			}
			assert.Equal(
				t,
				tt.want.value,
				tt.r.String(),
			)
		})
	}
}

func TestRating_Round(t *testing.T) {
	t.Parallel()
	type want struct {
		value int
		panic bool
	}
	tests := []struct {
		name string
		r    impl.Rating
		want want
	}{
		{
			name: "one",
			r:    impl.One,
			want: want{
				value: 1,
				panic: false,
			},
		},
		{
			name: "two",
			r:    impl.Two,
			want: want{
				value: 2,
				panic: false,
			},
		},
		{
			name: "three",
			r:    impl.Three,
			want: want{
				value: 3,
				panic: false,
			},
		},
		{
			name: "four",
			r:    impl.Four,
			want: want{
				value: 4,
				panic: false,
			},
		},
		{
			name: "five",
			r:    impl.Five,
			want: want{
				value: 5,
				panic: false,
			},
		},
		{
			name: "float ceil",
			r:    impl.NewRating(3.987656789),
			want: want{
				value: 4,
				panic: false,
			},
		},
		{
			name: "float floor",
			r:    impl.NewRating(3.937656789),
			want: want{
				value: 4,
				panic: false,
			},
		},
		{
			name: "float more than five",
			r:    impl.NewRating(5.00001),
			want: want{
				value: 0,
				panic: true,
			},
		},
		{
			name: "float lower than one",
			r:    impl.NewRating(0.9999),
			want: want{
				value: 0,
				panic: true,
			},
		},
		{
			name: "zero",
			r:    impl.NewRating(0),
			want: want{
				value: 0,
				panic: false,
			},
		},
		{
			name: "six",
			r:    impl.NewRating(6),
			want: want{
				value: 0,
				panic: true,
			},
		},
		{
			name: "negative",
			r:    impl.NewRating(-1),
			want: want{
				value: 0,
				panic: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = tt.r.Round()
				})
				return
			}
			assert.Equal(
				t,
				tt.want.value,
				tt.r.Round(),
			)
		})
	}
}

func TestRating_Int(t *testing.T) {
	t.Parallel()
	type want struct {
		value int
		panic bool
	}
	tests := []struct {
		name string
		r    impl.Rating
		want want
	}{
		{
			name: "one",
			r:    impl.One,
			want: want{
				value: 1,
				panic: false,
			},
		},
		{
			name: "two",
			r:    impl.Two,
			want: want{
				value: 2,
				panic: false,
			},
		},
		{
			name: "three",
			r:    impl.Three,
			want: want{
				value: 3,
				panic: false,
			},
		},
		{
			name: "four",
			r:    impl.Four,
			want: want{
				value: 4,
				panic: false,
			},
		},
		{
			name: "five",
			r:    impl.Five,
			want: want{
				value: 5,
				panic: false,
			},
		},
		{
			name: "float ceil",
			r:    impl.NewRating(3.987656789),
			want: want{
				value: 4,
				panic: false,
			},
		},
		{
			name: "float floor",
			r:    impl.NewRating(3.937656789),
			want: want{
				value: 4,
				panic: false,
			},
		},
		{
			name: "float more than five",
			r:    impl.NewRating(5.00001),
			want: want{
				value: 0,
				panic: true,
			},
		},
		{
			name: "float lower than one",
			r:    impl.NewRating(0.9999),
			want: want{
				value: 0,
				panic: true,
			},
		},
		{
			name: "zero",
			r:    impl.NewRating(0),
			want: want{
				value: 0,
				panic: false,
			},
		},
		{
			name: "six",
			r:    impl.NewRating(6),
			want: want{
				value: 0,
				panic: true,
			},
		},
		{
			name: "negative",
			r:    impl.NewRating(-1),
			want: want{
				value: 0,
				panic: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = tt.r.Int()
				})
				return
			}
			assert.Equal(
				t,
				tt.want.value,
				tt.r.Int(),
			)
		})
	}
}
