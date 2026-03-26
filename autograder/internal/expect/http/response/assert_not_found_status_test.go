package response_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
)

func TestAssertNotFoundStatus(t *testing.T) {
	t.Parallel()
	type args struct {
		rsp *http.Response
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
			name: "404",
			args: args{
				rsp: &http.Response{
					StatusCode: http.StatusNotFound,
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "400",
			args: args{
				rsp: &http.Response{
					StatusCode: http.StatusBadRequest,
				},
			},
			want: want{
				err: errs.ErrInvalidHTTPStatus,
			},
		},
		{
			name: "401",
			args: args{
				rsp: &http.Response{
					StatusCode: http.StatusUnauthorized,
				},
			},
			want: want{
				err: errs.ErrInvalidHTTPStatus,
			},
		},
		{
			name: "409",
			args: args{
				rsp: &http.Response{
					StatusCode: http.StatusConflict,
				},
			},
			want: want{
				err: errs.ErrInvalidHTTPStatus,
			},
		},
		{
			name: "201",
			args: args{
				rsp: &http.Response{
					StatusCode: http.StatusCreated,
				},
			},
			want: want{
				err: errs.ErrInvalidHTTPStatus,
			},
		},
		{
			name: "200",
			args: args{
				rsp: &http.Response{
					StatusCode: http.StatusOK,
				},
			},
			want: want{
				err: errs.ErrInvalidHTTPStatus,
			},
		},
		{
			name: "500",
			args: args{
				rsp: &http.Response{
					StatusCode: http.StatusInternalServerError,
				},
			},
			want: want{
				err: errs.ErrInvalidHTTPStatus,
			},
		},
		{
			name: "100",
			args: args{
				rsp: &http.Response{
					StatusCode: http.StatusContinue,
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
				response.AssertNotFoundStatus(
					tt.args.rsp,
				),
				tt.want.err,
			)
		})
	}
}
