package redis_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/config/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

func TestConfig_Validate(t *testing.T) {
	t.Parallel()
	type fields struct {
		addr     string
		password string
		db       int
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
				addr:     "localhost:6379",
				password: "secret",
				db:       0,
			},
			wantErr: false,
		},
		{
			name: "missing address",
			fields: fields{
				addr:     "",
				password: "secret",
				db:       0,
			},
			wantErr:         true,
			wantErrContains: "redis address is required",
		},
		{
			name: "invalid DB",
			fields: fields{
				addr:     "localhost:6379",
				password: "secret",
				db:       -1,
			},
			wantErr:         true,
			wantErrContains: "invalid redis DB",
		},
		{
			name: "empty password",
			fields: fields{
				addr:     "localhost:6379",
				password: "",
				db:       0,
			},
			wantErr: false,
		},
		{
			name:            "default values",
			fields:          fields{},
			wantErr:         true,
			wantErrContains: "redis address is required",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := redis.NewConfig(
				tt.fields.addr,
				tt.fields.password,
				tt.fields.db,
			)
			err := c.Validate()
			if tt.wantErr {
				assert.ErrorIs(t, err, errs.ErrInvalidConfig)
				assert.ErrorContains(t, err, tt.wantErrContains)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
