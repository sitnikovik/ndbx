package response_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
)

func TestAssertNotEmptyContent(t *testing.T) {
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
			name: "empty content",
			args: args{
				rsp: &http.Response{
					ContentLength: 0,
				},
			},
			wantErr: true,
		},
		{
			name: "non-empty content",
			args: args{
				rsp: &http.Response{
					ContentLength: 10,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := response.AssertNotEmptyContent(tt.args.rsp)
			if tt.wantErr {
				assert.ErrorIs(t, got, errs.ErrExpectationFailed)
			} else {
				assert.NoError(t, got)
			}
		})
	}
}
