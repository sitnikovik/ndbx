package cookie_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/cookie"
)

func TestAssertExistsHTTPOnly(t *testing.T) {
	t.Parallel()
	type args struct {
		ckk  []*http.Cookie
		name string
	}
	type want struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "ok",
			args: args{
				ckk: []*http.Cookie{
					{
						Name:     "test",
						Value:    "",
						HttpOnly: true,
					},
				},
				name: "test",
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "exists but http only is false",
			args: args{
				ckk: []*http.Cookie{
					{
						Name:     "test",
						Value:    "",
						HttpOnly: false,
					},
				},
				name: "test",
			},
			want: want{
				err: errs.ErrInvalidValue,
			},
		},
		{
			name: "not exists",
			args: args{
				ckk: []*http.Cookie{
					{
						Name:     "test1",
						Value:    "",
						HttpOnly: true,
					},
				},
				name: "test",
			},
			want: want{
				err: errs.ErrMissedCookie,
			},
		},
		{
			name: "empty list",
			args: args{
				ckk:  []*http.Cookie{},
				name: "test",
			},
			want: want{
				err: errs.ErrMissedCookie,
			},
		},
		{
			name: "empty name",
			args: args{
				ckk: []*http.Cookie{
					{
						Name:     "test",
						Value:    "",
						HttpOnly: true,
					},
				},
				name: "",
			},
			want: want{
				err: errs.ErrMissedCookie,
			},
		},
		{
			name: "default args",
			args: args{
				ckk:  nil,
				name: "",
			},
			want: want{
				err: errs.ErrMissedCookie,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.ErrorIs(
				t,
				cookie.AssertExistsHTTPOnly(
					tt.args.ckk,
					tt.args.name,
				),
				tt.want.err,
			)
		})
	}
}
