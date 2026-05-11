package body_test

import (
	"math"
	"net/url"
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
				body.WithTags("culture", "exhibition"),
			),
			want: want{
				val: []byte(`{` +
					`"category":"Music",` +
					`"city":"New York",` +
					`"price":50,` +
					`"tags":["culture","exhibition"]` +
					`}`,
				),
				panic: false,
			},
		},
		{
			name: "empty fields",
			b:    body.NewBody(),
			want: want{
				val:   []byte(`{}`),
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

func TestBody_URLQuery(t *testing.T) {
	t.Parallel()
	type want struct {
		value url.Values
	}
	tests := []struct {
		name string
		b    body.Body
		want want
	}{
		{
			name: "only with cascade",
			b: body.NewBody(
				body.WithCascade(),
			),
			want: want{
				value: func() url.Values {
					q := make(url.Values, 1)
					q.Set("cascade", "true")
					return q
				}(),
			},
		},
		{
			name: "with cascade and event fields",
			b: body.NewBody(
				body.WithCategory("Music"),
				body.WithCity("New York"),
				body.WithPrice(50),
				body.WithTags("culture", "exhibition"),
				body.WithCascade(),
			),
			want: want{
				value: func() url.Values {
					q := make(url.Values, 1)
					q.Set("cascade", "true")
					return q
				}(),
			},
		},
		{
			name: "empty",
			b:    body.NewBody(),
			want: want{
				value: make(url.Values),
			},
		},
		{
			name: "default value",
			b:    body.Body{},
			want: want{
				value: make(url.Values),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.b.URLQuery(),
			)
		})
	}
}
