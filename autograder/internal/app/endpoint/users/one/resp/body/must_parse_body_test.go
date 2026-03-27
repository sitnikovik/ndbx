package body_test

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/users/one/resp/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
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
					strings.NewReader(
						`{` +
							`"id": "1",` +
							`"username": "sams3p1ol",` +
							`"full_name": "Sam Sepiol"` +
							`}`,
					),
				),
			},
			want: want{
				val: body.NewBody(
					user.NewUser(
						user.NewID("1"),
						"sams3p1ol",
						"Sam Sepiol",
					),
				),
				panic: false,
			},
		},
		{
			name: "invalid json",
			args: args{
				body: io.NopCloser(
					strings.NewReader(`not json`),
				),
			},
			want: want{
				val:   body.Body{},
				panic: true,
			},
		},
		{
			name: "default args",
			args: args{
				body: nil,
			},
			want: want{
				val:   body.Body{},
				panic: true,
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
