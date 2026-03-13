package config_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/config"
	"github.com/sitnikovik/ndbx/autograder/internal/config/app"
	"github.com/sitnikovik/ndbx/autograder/internal/config/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/config/app/user/session"
	"github.com/sitnikovik/ndbx/autograder/internal/config/mongo"
	"github.com/sitnikovik/ndbx/autograder/internal/config/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

func TestConfig_Validate(t *testing.T) {
	t.Parallel()
	type fields struct {
		redis redis.Config
		mongo mongo.Config
		app   app.Config
	}
	tests := []struct {
		name            string
		fields          fields
		wantErr         error
		wantErrContains string
	}{
		{
			name: "ok",
			fields: fields{
				redis: redis.NewConfig(
					"localhost:6379",
					"",
					0,
				),
				app: app.NewConfig(
					user.NewConfig(
						session.NewConfig(1*time.Second),
					),
					"localhost",
					8080,
				),
				mongo: mongo.NewConfig(
					"testdb",
					"testuser",
					"testpass",
					"localhost",
					27017,
				),
			},
			wantErr: nil,
		},
		{
			name: "invalid redis config",
			fields: fields{
				redis: redis.NewConfig(
					"",
					"",
					0,
				),
				app: app.NewConfig(
					user.NewConfig(
						session.NewConfig(0),
					),
					"",
					0,
				),
				mongo: mongo.NewConfig(
					"testdb",
					"testuser",
					"testpass",
					"localhost",
					27017,
				),
			},
			wantErr:         errs.ErrInvalidConfig,
			wantErrContains: "redis",
		},
		{
			name: "invalid app config",
			fields: fields{
				redis: redis.NewConfig(
					"localhost:6379",
					"",
					0,
				),
				mongo: mongo.NewConfig(
					"testdb",
					"testuser",
					"testpass",
					"localhost",
					27017,
				),
				app: app.NewConfig(
					user.NewConfig(
						session.NewConfig(0),
					),
					"",
					0,
				),
			},
			wantErr:         errs.ErrInvalidConfig,
			wantErrContains: "app",
		},
		{
			name: "invalid mongo config",
			fields: fields{
				redis: redis.NewConfig(
					"localhost:6379",
					"",
					0,
				),
				app: app.NewConfig(
					user.NewConfig(
						session.NewConfig(1*time.Second),
					),
					"localhost",
					8080,
				),
				mongo: mongo.NewConfig(
					"",
					"",
					"",
					"",
					0,
				),
			},
			wantErr:         errs.ErrInvalidConfig,
			wantErrContains: "mongo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := config.NewConfig(
				tt.fields.redis,
				tt.fields.mongo,
				tt.fields.app,
			)
			err := c.Validate()
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Contains(t, err.Error(), tt.wantErrContains)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
