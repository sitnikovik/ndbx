package resp_test

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/resp"
)

func TestMustParseError(t *testing.T) {
	t.Parallel()
	type args struct {
		body io.ReadCloser
	}
	type want struct {
		val   resp.ErrorResponse
		panic bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "ok",
			args: args{
				body: io.NopCloser(
					strings.NewReader(`{"message":"already exists"}`),
				),
			},
			want: want{
				val:   resp.NewErrorResponse("already exists"),
				panic: false,
			},
		},
		{
			name: "invalid JSON",
			args: args{
				body: io.NopCloser(
					strings.NewReader(`{"message":"already exists`),
				),
			},
			want: want{
				panic: true,
			},
		},
		{
			name: "missing message field",
			args: args{
				body: io.NopCloser(
					strings.NewReader(`{"name":"value"}`),
				),
			},
			want: want{
				val:   resp.NewErrorResponse(""),
				panic: false,
			},
		},
		{
			name: "empty body",
			args: args{
				body: io.NopCloser(
					strings.NewReader(``),
				),
			},
			want: want{
				panic: true,
			},
		},
		{
			name: "message field is empty",
			args: args{
				body: io.NopCloser(
					strings.NewReader(`{"message":""}`),
				),
			},
			want: want{
				val:   resp.NewErrorResponse(""),
				panic: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = resp.MustParseError(tt.args.body)
				})
				return
			}
			assert.Equal(
				t,
				tt.want.val,
				resp.MustParseError(tt.args.body),
			)
		})
	}
}
