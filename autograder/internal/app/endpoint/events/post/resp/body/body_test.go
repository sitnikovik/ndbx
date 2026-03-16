package body_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/post/resp/body"
)

func TestBody_ID(t *testing.T) {
	t.Parallel()
	type want struct {
		val string
	}
	tests := []struct {
		name string
		b    body.Body
		want want
	}{
		{
			name: "ok",
			b:    body.NewBody("123"),
			want: want{
				val: "123",
			},
		},
		{
			name: "empty id",
			b:    body.NewBody(""),
			want: want{
				val: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.b.ID(),
			)
		})
	}
}
