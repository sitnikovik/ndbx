package body_test

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/event/post/resp/body"
)

func TestMustParseBody(t *testing.T) {
	t.Parallel()
	type args struct {
		body io.ReadCloser
	}
	type want struct {
		val   body.Body
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
					strings.NewReader(`{"id":"123"}`),
				),
			},
			want: want{
				val:   body.NewBody("123"),
				panic: false,
			},
		},
		{
			name: "invalid JSON",
			args: args{
				body: io.NopCloser(
					strings.NewReader(`{"id":123}`),
				),
			},
			want: want{
				panic: true,
			},
		},
		{
			name: "missing id field",
			args: args{
				body: io.NopCloser(
					strings.NewReader(`{"name":"Event Name"}`),
				),
			},
			want: want{
				val:   body.NewBody(""),
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
			name: "id field is empty",
			args: args{
				body: io.NopCloser(
					strings.NewReader(`{"id":""}`),
				),
			},
			want: want{
				val:   body.NewBody(""),
				panic: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = body.MustParseBody(tt.args.body)
				})
				return
			}
			assert.Equal(
				t,
				tt.want.val,
				body.MustParseBody(tt.args.body),
			)
		})
	}
}
