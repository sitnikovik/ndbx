package cookie_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/cookie"
)

func TestAssertValuef(t *testing.T) {
	t.Parallel()
	type args struct {
		ckk  []*http.Cookie
		name string
		f    func(v string) error
	}
	type want struct {
		err   error
		panic bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "exists but not passed",
			args: args{
				ckk: []*http.Cookie{
					{
						Name:  "test",
						Value: "",
					},
				},
				name: "test",
				f: func(v string) error {
					if v == "" {
						return errs.ErrInvalidValue
					}
					return nil
				},
			},
			want: want{
				err:   errs.ErrInvalidValue,
				panic: false,
			},
		},
		{
			name: "not exists",
			args: args{
				ckk: []*http.Cookie{
					{
						Name:  "test1",
						Value: "",
					},
				},
				name: "test",
				f: func(v string) error {
					if v == "" {
						return errs.ErrInvalidValue
					}
					return nil
				},
			},
			want: want{
				err:   errs.ErrMissedCookie,
				panic: false,
			},
		},
		{
			name: "empty list",
			args: args{
				ckk:  []*http.Cookie{},
				name: "test",
				f: func(v string) error {
					if v == "" {
						return errs.ErrInvalidValue
					}
					return nil
				},
			},
			want: want{
				err:   errs.ErrMissedCookie,
				panic: false,
			},
		},
		{
			name: "empty name",
			args: args{
				ckk: []*http.Cookie{
					{
						Name:  "test",
						Value: "",
					},
				},
				name: "",
				f: func(v string) error {
					if v == "" {
						return errs.ErrInvalidValue
					}
					return nil
				},
			},
			want: want{
				err:   errs.ErrMissedCookie,
				panic: false,
			},
		},
		{
			name: "nil f",
			args: args{
				ckk: []*http.Cookie{
					{
						Name:  "test",
						Value: "",
					},
				},
				name: "test",
				f:    nil,
			},
			want: want{
				err:   nil,
				panic: true,
			},
		},
		{
			name: "default args",
			args: args{
				ckk:  nil,
				name: "",
				f:    nil,
			},
			want: want{
				err:   errs.ErrMissedCookie,
				panic: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = cookie.AssertValueFn(
						tt.args.ckk,
						tt.args.name,
						tt.args.f,
					)
				})
				return
			}
			assert.ErrorIs(
				t,
				cookie.AssertValueFn(
					tt.args.ckk,
					tt.args.name,
					tt.args.f,
				),
				tt.want.err,
			)
		})
	}
}
