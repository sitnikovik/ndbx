package response_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
)

func TestAssertOKStatus(t *testing.T) {
	t.Parallel()
	type args struct {
		rsp *http.Response
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "201",
			args: args{
				rsp: &http.Response{
					StatusCode: http.StatusCreated,
				},
			},
			wantErr: true,
		},
		{
			name: "200",
			args: args{
				rsp: &http.Response{
					StatusCode: http.StatusOK,
				},
			},
			wantErr: false,
		},
		{
			name: "400",
			args: args{
				rsp: &http.Response{
					StatusCode: http.StatusBadRequest,
				},
			},
			wantErr: true,
		},
		{
			name: "500",
			args: args{
				rsp: &http.Response{
					StatusCode: http.StatusInternalServerError,
				},
			},
			wantErr: true,
		},
		{
			name: "100",
			args: args{
				rsp: &http.Response{
					StatusCode: http.StatusContinue,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := response.AssertOKStatus(tt.args.rsp)
			if tt.wantErr {
				assert.ErrorIs(t, got, errs.ErrExpectationFailed)
				assert.ErrorIs(t, got, errs.ErrInvalidHTTPStatus)
			} else {
				assert.NoError(t, got)
			}
		})
	}
}
