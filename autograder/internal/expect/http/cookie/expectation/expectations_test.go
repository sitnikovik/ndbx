package expectation_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/cookie"
	impl "github.com/sitnikovik/ndbx/autograder/internal/expect/http/cookie/expectation"
)

func TestExpectations_Assert(t *testing.T) {
	t.Parallel()
	type args struct {
		ckk []*http.Cookie
	}
	type want struct {
		err error
	}
	tests := []struct {
		name string
		e    impl.Expectations
		args args
		want want
	}{
		{
			name: "ok",
			e: impl.NewExpectations(
				"foo",
				impl.WithAsserts(
					cookie.AssertExists,
					cookie.AssertExistsMaxAge,
				),
				impl.WithAssertsValueFn(
					func(v string) error {
						if v != "bar" {
							return errs.ErrInvalidValue
						}
						return nil
					},
				),
			),
			args: args{
				ckk: []*http.Cookie{
					{
						Name:   "foo",
						Value:  "bar",
						MaxAge: 900,
					},
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "not satisfies value",
			e: impl.NewExpectations(
				"foo",
				impl.WithAsserts(
					cookie.AssertExists,
					cookie.AssertExistsMaxAge,
				),
				impl.WithAssertsValueFn(
					func(v string) error {
						if v != "foo" {
							return errs.ErrInvalidValue
						}
						return nil
					},
				),
			),
			args: args{
				ckk: []*http.Cookie{
					{
						Name:   "foo",
						Value:  "bar",
						MaxAge: 900,
					},
				},
			},
			want: want{
				err: errs.ErrInvalidValue,
			},
		},
		{
			name: "not satisfies max age",
			e: impl.NewExpectations(
				"foo",
				impl.WithAsserts(
					cookie.AssertExists,
					cookie.AssertExistsMaxAge,
				),
				impl.WithAssertsValueFn(
					func(v string) error {
						if v != "bar" {
							return errs.ErrInvalidValue
						}
						return nil
					},
				),
			),
			args: args{
				ckk: []*http.Cookie{
					{
						Name:   "foo",
						Value:  "bar",
						MaxAge: 0,
					},
				},
			},
			want: want{
				err: errs.ErrInvalidValue,
			},
		},
		{
			name: "not found by name",
			e: impl.NewExpectations(
				"bar",
				impl.WithAsserts(
					cookie.AssertExists,
					cookie.AssertExistsMaxAge,
				),
				impl.WithAssertsValueFn(
					func(v string) error {
						if v != "bar" {
							return errs.ErrInvalidValue
						}
						return nil
					},
				),
			),
			args: args{
				ckk: []*http.Cookie{
					{
						Name:   "foo",
						Value:  "bar",
						MaxAge: 0,
					},
				},
			},
			want: want{
				err: errs.ErrMissedCookie,
			},
		},
		{
			name: "without asserts",
			e: impl.NewExpectations(
				"bar",
				impl.WithAsserts(),
			),
			args: args{
				ckk: []*http.Cookie{
					{
						Name:   "foo",
						Value:  "bar",
						MaxAge: 0,
					},
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "empty name",
			e: impl.NewExpectations(
				"",
				impl.WithAsserts(
					cookie.AssertExists,
				),
			),
			args: args{
				ckk: []*http.Cookie{
					{
						Name:   "foo",
						Value:  "bar",
						MaxAge: 0,
					},
				},
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
				tt.e.Assert(
					tt.args.ckk,
				),
				tt.want.err,
			)
		})
	}
}
