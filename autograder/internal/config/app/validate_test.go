package app_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/config/app"
	"github.com/sitnikovik/ndbx/autograder/internal/config/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/config/app/user/session"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

func TestConfig_Validate(t *testing.T) {
	t.Parallel()
	type fields struct {
		user user.Config
		host string
		port int
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
				user: user.NewConfig(
					session.NewConfig(1 * time.Minute),
				),
				host: "localhost",
				port: 8080,
			},
		},
		{
			name: "invalid user config",
			fields: fields{
				user: user.NewConfig(
					session.NewConfig(0),
				),
				host: "localhost",
				port: 8080,
			},
			wantErr:         true,
			wantErrContains: "user: session: TTL must be greater than zero",
		},
		{
			name: "empty host",
			fields: fields{
				user: user.NewConfig(
					session.NewConfig(1 * time.Minute),
				),
				host: "",
				port: 8080,
			},
			wantErr:         true,
			wantErrContains: "host is required",
		},
		{
			name: "zero port",
			fields: fields{
				user: user.NewConfig(
					session.NewConfig(1 * time.Minute),
				),
				host: "localhost",
				port: 0,
			},
			wantErr:         true,
			wantErrContains: "invalid port number",
		},
		{
			name: "too large port",
			fields: fields{
				user: user.NewConfig(
					session.NewConfig(1 * time.Minute),
				),
				host: "localhost",
				port: 100000,
			},
			wantErr:         true,
			wantErrContains: "invalid port number",
		},
		{
			name:            "default value",
			fields:          fields{},
			wantErr:         true,
			wantErrContains: "user: session: TTL must be greater than zero",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := app.NewConfig(
				tt.fields.user,
				tt.fields.host,
				tt.fields.port,
			)
			err := c.Validate()
			if tt.wantErr {
				assert.ErrorIs(t, err, errs.ErrInvalidConfig)
				assert.Containsf(t, err.Error(), tt.wantErrContains, "got error: %v", err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
