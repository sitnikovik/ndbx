package response_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
)

func TestAssertAll(t *testing.T) {
	t.Parallel()
	type args struct {
		rsp *http.Response
		ff  []response.AssertFunc
	}
	type want struct {
		errs  []error
		panic bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "all assertions pass",
			args: args{
				rsp: &http.Response{
					StatusCode:    http.StatusCreated,
					ContentLength: 0,
				},
				ff: []response.AssertFunc{
					response.AssertCreatedStatus,
					response.AssertEmptyContent,
				},
			},
			want: want{errs: nil},
		},
		{
			name: "one assertion fails",
			args: args{
				rsp: &http.Response{
					StatusCode:    http.StatusOK,
					ContentLength: 0,
				},
				ff: []response.AssertFunc{
					response.AssertCreatedStatus,
					response.AssertEmptyContent,
				},
			},
			want: want{
				errs: []error{
					errs.ErrInvalidHTTPStatus,
					errs.ErrExpectationFailed,
				},
			},
		},
		{
			name: "the second assertion fails",
			args: args{
				rsp: &http.Response{
					StatusCode:    http.StatusCreated,
					ContentLength: 10,
				},
				ff: []response.AssertFunc{
					response.AssertCreatedStatus,
					response.AssertEmptyContent,
				},
			},
			want: want{
				errs: []error{
					errs.ErrExpectationFailed,
				},
			},
		},
		{
			name: "empty assertion list",
			args: args{
				rsp: &http.Response{
					StatusCode:    http.StatusOK,
					ContentLength: 10,
				},
				ff: []response.AssertFunc{},
			},
			want: want{
				errs:  nil,
				panic: true,
			},
		},
		{
			name: "nil response",
			args: args{
				rsp: nil,
				ff: []response.AssertFunc{
					response.AssertCreatedStatus,
					response.AssertEmptyContent,
				},
			},
			want: want{
				errs:  nil,
				panic: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = response.AssertAll(
						tt.args.rsp,
						tt.args.ff...,
					)
				})
				return
			}
			err := response.AssertAll(
				tt.args.rsp,
				tt.args.ff...,
			)
			if len(tt.want.errs) > 0 {
				for _, want := range tt.want.errs {
					assert.ErrorIs(t, err, want)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
