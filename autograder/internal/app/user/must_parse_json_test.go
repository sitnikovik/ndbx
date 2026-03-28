package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

func TestMustParseJSON(t *testing.T) {
	t.Parallel()
	type args struct {
		bb []byte
	}
	type want struct {
		val   user.User
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
				bb: []byte(
					`{` +
						`"id": "1",` +
						`"username": "sams3p1ol",` +
						`"full_name": "Sam Sepiol"` +
						`}`,
				),
			},
			want: want{
				val: user.NewUser(
					user.NewID("1"),
					"sams3p1ol",
					"Sam Sepiol",
				),
				panic: false,
			},
		},
		{
			name: "empty",
			args: args{
				bb: []byte(`{}`),
			},
			want: want{
				val:   user.User{},
				panic: false,
			},
		},
		{
			name: "not a json",
			args: args{
				bb: []byte(`foo`),
			},
			want: want{
				val:   user.User{},
				panic: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = user.MustParseJSON(tt.args.bb)
				})
				return
			}
			assert.Equal(
				t,
				tt.want.val,
				user.MustParseJSON(tt.args.bb),
			)
		})
	}
}
