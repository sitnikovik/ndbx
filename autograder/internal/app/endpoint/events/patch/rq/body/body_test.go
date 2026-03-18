package body_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/patch/rq/body"
)

func TestBody_MustBytes(t *testing.T) {
	t.Parallel()
	type want struct {
		val   []byte
		panic bool
	}
	tests := []struct {
		name string
		b    body.Body
		want want
	}{
		{
			name: "ok",
			b: body.NewBody(
				body.WithCategory("Music"),
				body.WithCity("New York"),
				body.WithPrice(50),
			),
			want: want{
				val: []byte(`{` +
					`"category":"Music",` +
					`"city":"New York",` +
					`"price":50` +
					`}`,
				),
				panic: false,
			},
		},
		{
			name: "empty fields",
			b:    body.NewBody(),
			want: want{
				val: []byte(`{` +
					`"category":"",` +
					`"city":"",` +
					`"price":0` +
					`}`,
				),
				panic: false,
			},
		},
		{
			name: "price is max uint",
			b: body.NewBody(
				body.WithPrice(math.MaxUint64),
			),
			want: want{
				val: []byte(`{` +
					`"category":"",` +
					`"city":"",` +
					`"price":18446744073709551615` +
					`}`,
				),
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
			} else {
				assert.Equal(
					t,
					tt.want.val,
					tt.b.MustBytes(),
				)
			}
		})
	}
}
