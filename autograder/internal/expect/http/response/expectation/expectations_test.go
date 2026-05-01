package expectation_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/cookie"
	cookiexpct "github.com/sitnikovik/ndbx/autograder/internal/expect/http/cookie/expectation"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
	impl "github.com/sitnikovik/ndbx/autograder/internal/expect/http/response/expectation"
)

func TestExpectations_Assert(t *testing.T) {
	t.Parallel()
	type args struct {
		resp *http.Response
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
				impl.WithAsserts(
					response.AssertNoContentStatus,
					response.AssertEmptyContent,
				),
				impl.WithCookies(
					cookiexpct.NewExpectations(
						"session_id",
						cookiexpct.WithAsserts(
							cookie.AssertExistsMaxAge,
							cookie.AssertExistsHTTPOnly,
						),
					),
				),
			),
			args: args{
				resp: &http.Response{
					StatusCode: http.StatusNoContent,
					Body:       http.NoBody,
					Header: http.Header{
						"Set-Cookie": []string{
							"session_id=abc123; HttpOnly; Secure; SameSite=Strict; Path=/; Max-Age=86400",
						},
					},
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "resp not passed",
			e: impl.NewExpectations(
				impl.WithAsserts(
					response.AssertBadRequestStatus,
				),
				impl.WithCookies(
					cookiexpct.NewExpectations(
						"session_id",
						cookiexpct.WithAsserts(
							cookie.AssertExistsMaxAge,
							cookie.AssertExistsHTTPOnly,
						),
					),
				),
			),
			args: args{
				resp: &http.Response{
					StatusCode: http.StatusNoContent,
					Body:       http.NoBody,
					Header: http.Header{
						"Set-Cookie": []string{
							"session_id=abc123; HttpOnly; Secure; SameSite=Strict; Path=/; Max-Age=86400",
						},
					},
				},
			},
			want: want{
				err: errs.ErrInvalidHTTPStatus,
			},
		},
		{
			name: "cookie value not passed",
			e: impl.NewExpectations(
				impl.WithAsserts(
					response.AssertNoContentStatus,
				),
				impl.WithCookies(
					cookiexpct.NewExpectations(
						"session_id",
						cookiexpct.WithAsserts(
							cookie.AssertExistsMaxAge,
						),
						cookiexpct.WithAssertsValueFn(
							func(v string) error {
								if v != "abc" {
									return errs.ErrInvalidValue
								}
								return nil
							},
						),
					),
				),
			),
			args: args{
				resp: &http.Response{
					StatusCode: http.StatusNoContent,
					Body:       http.NoBody,
					Header: http.Header{
						"Set-Cookie": []string{
							"session_id=abc123; HttpOnly; Secure; SameSite=Strict; Path=/; Max-Age=86400",
						},
					},
				},
			},
			want: want{
				err: errs.ErrInvalidValue,
			},
		},
		{
			name: "without resp asserts but cookies",
			e: impl.NewExpectations(
				impl.WithCookies(
					cookiexpct.NewExpectations(
						"session_id",
						cookiexpct.WithAsserts(
							cookie.AssertExistsMaxAge,
						),
					),
				),
			),
			args: args{
				resp: &http.Response{
					StatusCode: http.StatusNoContent,
					Body:       http.NoBody,
					Header: http.Header{
						"Set-Cookie": []string{
							"session_id=abc123; HttpOnly; Secure; SameSite=Strict; Path=/; Max-Age=0",
						},
					},
				},
			},
			want: want{
				err: errs.ErrInvalidValue,
			},
		},
		{
			name: "without cookies asserts",
			e: impl.NewExpectations(
				impl.WithAsserts(
					response.AssertBadRequestStatus,
				),
			),
			args: args{
				resp: &http.Response{
					StatusCode: http.StatusNoContent,
					Body:       http.NoBody,
					Header: http.Header{
						"Set-Cookie": []string{
							"session_id=abc123; HttpOnly; Secure; SameSite=Strict; Path=/; Max-Age=0",
						},
					},
				},
			},
			want: want{
				err: errs.ErrInvalidHTTPStatus,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.ErrorIs(
				t,
				tt.e.Assert(
					tt.args.resp,
				),
				tt.want.err,
			)
		})
	}
}
