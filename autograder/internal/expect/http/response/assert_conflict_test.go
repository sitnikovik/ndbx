package response_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
)

func TestAssertConflictStatus(t *testing.T) {
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
			name: "409",
			args: args{
				rsp: &http.Response{
					StatusCode: http.StatusConflict,
				},
			},
			want: want{
				err: nil,
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
			got := response.AssertConflictStatus(tt.args.rsp)
			if tt.want.err != nil {
				assert.ErrorIs(t, got, tt.want.err)
				assert.ErrorIs(t, got, errs.ErrInvalidHTTPStatus)
			} else {
				assert.NoError(t, got)
			}
		})
	}
}
