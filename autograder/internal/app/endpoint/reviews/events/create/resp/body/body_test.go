package body_test

import (
	"testing"

	impl "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/reviews/events/create/resp/body"
	"github.com/stretchr/testify/assert"
)

func TestBody_ID(t *testing.T) {
	t.Parallel()
	type want struct {
		val string
	}
	tests := []struct {
		name string
		b    impl.Body
		want want
	}{
		{
			name: "ok",
			b:    impl.NewBody("123"),
			want: want{
				val: "123",
			},
		},
		{
			name: "empty id",
			b:    impl.NewBody(""),
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
