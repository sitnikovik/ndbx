package session_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/config/app/user/session"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

func TestConfig_Validate(t *testing.T) {
	t.Parallel()
	type fields struct {
		ttl time.Duration
	}
	tests := []struct {
		name            string
		fields          fields
		wantErr         bool
		wantErrContains string
	}{
		{
			name: "gt zero",
			fields: fields{
				ttl: time.Minute,
			},
			wantErr: false,
		},
		{
			name: "zero",
			fields: fields{
				ttl: 0,
			},
			wantErr:         true,
			wantErrContains: "TTL must be greater than zero",
		},
		{
			name: "lt zero",
			fields: fields{
				ttl: -1,
			},
			wantErr:         true,
			wantErrContains: "TTL must be greater than zero",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := session.NewConfig(tt.fields.ttl)
			err := c.Validate()
			if tt.wantErr {
				assert.ErrorIs(t, err, errs.ErrInvalidConfig)
				assert.Contains(t, err.Error(), tt.wantErrContains)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
