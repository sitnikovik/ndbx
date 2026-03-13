package mongo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/config/mongo"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

func TestConfig_Validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		c               mongo.Config
		wantErr         error
		wantErrContains string
	}{
		{
			name: "ok",
			c: mongo.NewConfig(
				"testdb",
				"testuser",
				"testpass",
				"localhost",
				27017,
			),
			wantErr:         nil,
			wantErrContains: "",
		},
		{
			name: "missing database",
			c: mongo.NewConfig(
				"",
				"testuser",
				"testpass",
				"localhost",
				27017,
			),
			wantErr:         errs.ErrInvalidConfig,
			wantErrContains: "database name is required",
		},
		{
			name: "missing username",
			c: mongo.NewConfig(
				"testdb",
				"",
				"testpass",
				"localhost",
				27017,
			),
			wantErr:         nil,
			wantErrContains: "",
		},
		{
			name: "missing password",
			c: mongo.NewConfig(
				"testdb",
				"testuser",
				"",
				"localhost",
				27017,
			),
			wantErr:         nil,
			wantErrContains: "",
		},
		{
			name: "missing host",
			c: mongo.NewConfig(
				"testdb",
				"testuser",
				"testpass",
				"",
				27017,
			),
			wantErr:         errs.ErrInvalidConfig,
			wantErrContains: "host is required",
		},
		{
			name: "invalid port",
			c: mongo.NewConfig(
				"testdb",
				"testuser",
				"testpass",
				"localhost",
				-1,
			),
			wantErr:         errs.ErrInvalidConfig,
			wantErrContains: "invalid port",
		},
		{
			name: "port too large",
			c: mongo.NewConfig(
				"testdb",
				"testuser",
				"testpass",
				"localhost",
				100000,
			),
			wantErr:         errs.ErrInvalidConfig,
			wantErrContains: "invalid port",
		},
		{
			name:            "empty config",
			c:               mongo.NewConfig("", "", "", "", 0),
			wantErr:         errs.ErrInvalidConfig,
			wantErrContains: "database name is required",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.c.Validate()
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				assert.ErrorContains(t, err, tt.wantErrContains)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
