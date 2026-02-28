package user_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/config/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/config/app/user/session"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

func TestConfig_Validate(t *testing.T) {
	t.Parallel()
	type fields struct {
		session session.Config
	}
	tests := []struct {
		name            string
		fields          fields
		wantErr         bool
		wantErrContains string
	}{
		{
			name: "ok",
			fields: fields{
				session: session.NewConfig(1 * time.Minute),
			},
			wantErr: false,
		},
		{
			name: "zero ttl for session",
			fields: fields{
				session: session.NewConfig(0),
			},
			wantErr:         true,
			wantErrContains: "session: TTL must be greater than zero",
		},
		{
			name: "lt zero ttl for session",
			fields: fields{
				session: session.NewConfig(-1),
			},
			wantErr:         true,
			wantErrContains: "session: TTL must be greater than zero",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := user.NewConfig(tt.fields.session)
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
