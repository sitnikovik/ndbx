package session_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/user/session"
)

func TestValidate(t *testing.T) {
	t.Parallel()
	type args struct {
		sid string
	}
	tests := []struct {
		name            string
		args            args
		wantErrContains string
	}{
		{
			name: "valid session ID",
			args: args{
				sid: "0123456789abcdef0123456789abcdef",
			},
		},
		{
			name: "too short session ID",
			args: args{
				sid: "0123456789abcdef0123456789abcde",
			},
			wantErrContains: "must be at least 32 characters long",
		},
		{
			name: "non-hexadecimal characters in session ID",
			args: args{
				sid: "0123456789abcdef0123456789abcdeg",
			},
			wantErrContains: "must be a hexadecimal string",
		},
		{
			name: "empty session ID",
			args: args{
				sid: "",
			},
			wantErrContains: "must be at least 32 characters long",
		},
		{
			name: "with uppercase hexadecimal characters",
			args: args{
				sid: "0123456789ABCDEF0123456789ABCDEF",
			},
			wantErrContains: "",
		},
		{
			name: "with mixed case hexadecimal characters",
			args: args{
				sid: "0123456789AbCdEf0123456789aBcDeF",
			},
			wantErrContains: "",
		},
		{
			name: "with special characters",
			args: args{
				sid: "0123456,789abcdef.0123456789abcde!",
			},
			wantErrContains: "must be a hexadecimal string",
		},
		{
			name: "uuid",
			args: args{
				sid: "3fa56e2b-8b7d-485b-aa79-c5e90c959cb7",
			},
			wantErrContains: "must be a hexadecimal string",
		},
		{
			name: "only digits",
			args: args{
				sid: "01234567890123456789012345678901",
			},
			wantErrContains: "",
		},
		{
			name: "only letters",
			args: args{
				sid: "abcdefabcdefabcdefabcdefabcdefab",
			},
			wantErrContains: "",
		},
		{
			name: "with whitespace",
			args: args{
				sid: "0123456789abcdef0123456789abcde ",
			},
			wantErrContains: "must be a hexadecimal string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := session.Validate(tt.args.sid)
			if tt.wantErrContains != "" {
				assert.ErrorContains(t, err, tt.wantErrContains)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
