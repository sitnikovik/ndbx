package resp_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/resp"
)

func TestErrorResponse_Error(t *testing.T) {
	t.Parallel()
	type want struct {
		val string
	}
	tests := []struct {
		name string
		e    resp.ErrorResponse
		want want
	}{
		{
			name: "ok",
			e:    resp.NewErrorResponse("already exists"),
			want: want{
				val: "already exists",
			},
		},
		{
			name: "empty message",
			e:    resp.NewErrorResponse(""),
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
				tt.e.Error(),
			)
		})
	}
}
