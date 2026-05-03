package body_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/reviews/events/create/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/rating"
)

func TestBody_MustBytes(t *testing.T) {
	t.Parallel()
	type want struct {
		value []byte
		panic bool
	}
	tests := []struct {
		name string
		b    impl.Body
		want want
	}{
		{
			name: "full",
			b: impl.NewBody(
				impl.WithComment("test"),
				impl.WithRating(rating.Five),
			),
			want: want{
				value: []byte(
					`{` +
						`"comment":"test",` +
						`"rating":5` +
						`}`,
				),
				panic: false,
			},
		},
		{
			name: "full but empty rating",
			b: impl.NewBody(
				impl.WithComment("test"),
				impl.WithRating(rating.None),
			),
			want: want{
				value: []byte(
					`{` +
						`"comment":"test"` +
						`}`,
				),
				panic: false,
			},
		},
		{
			name: "only comment",
			b: impl.NewBody(
				impl.WithComment("test"),
			),
			want: want{
				value: []byte(
					`{` +
						`"comment":"test"` +
						`}`,
				),
				panic: false,
			},
		},
		{
			name: "only rating",
			b: impl.NewBody(
				impl.WithRating(rating.Four),
			),
			want: want{
				value: []byte(
					`{` +
						`"rating":4` +
						`}`,
				),
				panic: false,
			},
		},
		{
			name: "custom float rating",
			b: impl.NewBody(
				impl.WithRating(3.54123),
			),
			want: want{
				value: []byte(
					`{` +
						`"rating":4` +
						`}`,
				),
				panic: false,
			},
		},
		{
			name: "negative rating",
			b: impl.NewBody(
				impl.WithRating(-1),
			),
			want: want{
				value: nil,
				panic: true,
			},
		},
		{
			name: "empty comment and without rating",
			b: impl.NewBody(
				impl.WithComment(""),
			),
			want: want{
				value: []byte(`{}`),
				panic: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = tt.b.MustBytes()
				})
				return
			}
			assert.Equal(
				t,
				tt.want.value,
				tt.b.MustBytes(),
			)
		})
	}
}
