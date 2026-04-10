package config_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/client/cassandra/consistency"
	"github.com/sitnikovik/ndbx/autograder/internal/config/app"
	"github.com/sitnikovik/ndbx/autograder/internal/config/app/event"
	reaction "github.com/sitnikovik/ndbx/autograder/internal/config/app/reaction/event"
	"github.com/sitnikovik/ndbx/autograder/internal/config/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/config/app/user/session"
	"github.com/sitnikovik/ndbx/autograder/internal/config/cassandra"
	impl "github.com/sitnikovik/ndbx/autograder/internal/config/lab5/config"
	"github.com/sitnikovik/ndbx/autograder/internal/config/mongo"
	"github.com/sitnikovik/ndbx/autograder/internal/config/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/console"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

func TestConfig_Validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		c               impl.Config
		wantErr         error
		wantErrContains string
	}{
		{
			name: "ok",
			c: impl.NewConfig(
				redis.NewConfig(
					"localhost:6379",
					"",
					0,
				),
				mongo.NewConfig(
					"testdb",
					"testuser",
					"testpass",
					"localhost",
					27017,
				),
				cassandra.NewConfig(
					cassandra.NewConnection(
						[]string{"localhost"},
						9042,
					),
					cassandra.NewAuth("", ""),
					cassandra.NewDatabase(
						"testkeyspace",
						consistency.Quorum,
					),
				),
				app.NewConfig(
					user.NewConfig(
						session.NewConfig(1*time.Second),
					),
					"localhost",
					8080,
					app.WithEvent(
						event.NewConfig(
							reaction.NewConfig(1*time.Minute),
						),
					),
				),
			),
			wantErr: nil,
		},
		{
			name: "invalid redis config",
			c: impl.NewConfig(
				redis.NewConfig(
					"",
					"",
					0,
				),
				mongo.NewConfig(
					"testdb",
					"testuser",
					"testpass",
					"localhost",
					27017,
				),
				cassandra.NewConfig(
					cassandra.NewConnection(
						[]string{"localhost"},
						9042,
					),
					cassandra.NewAuth("", ""),
					cassandra.NewDatabase(
						"testkeyspace",
						consistency.Quorum,
					),
				),
				app.NewConfig(
					user.NewConfig(
						session.NewConfig(0),
					),
					"",
					0,
				),
			),
			wantErr:         errs.ErrInvalidConfig,
			wantErrContains: "redis",
		},
		{
			name: "invalid app config",
			c: impl.NewConfig(
				redis.NewConfig(
					"localhost:6379",
					"",
					0,
				),
				mongo.NewConfig(
					"testdb",
					"testuser",
					"testpass",
					"localhost",
					27017,
				),
				cassandra.NewConfig(
					cassandra.NewConnection(
						[]string{"localhost"},
						9042,
					),
					cassandra.NewAuth("", ""),
					cassandra.NewDatabase(
						"testkeyspace",
						consistency.Quorum,
					),
				),
				app.NewConfig(
					user.NewConfig(
						session.NewConfig(0),
					),
					"",
					0,
				),
			),
			wantErr:         errs.ErrInvalidConfig,
			wantErrContains: "app",
		},
		{
			name: "invalid app event reations ttl",
			c: impl.NewConfig(
				redis.NewConfig(
					"localhost:6379",
					"",
					0,
				),
				mongo.NewConfig(
					"testdb",
					"testuser",
					"testpass",
					"localhost",
					27017,
				),
				cassandra.NewConfig(
					cassandra.NewConnection(
						[]string{"localhost"},
						9042,
					),
					cassandra.NewAuth("", ""),
					cassandra.NewDatabase(
						"testkeyspace",
						consistency.Quorum,
					),
				),
				app.NewConfig(
					user.NewConfig(
						session.NewConfig(0),
					),
					"",
					0,
					app.WithEvent(
						event.NewConfig(
							reaction.NewConfig(0),
						),
					),
				),
			),
			wantErr:         errs.ErrInvalidConfig,
			wantErrContains: "app",
		},
		{
			name: "invalid mongo config",
			c: impl.NewConfig(
				redis.NewConfig(
					"localhost:6379",
					"",
					0,
				),
				mongo.NewConfig(
					"",
					"",
					"",
					"",
					0,
				),
				cassandra.NewConfig(
					cassandra.NewConnection(
						[]string{"localhost"},
						9042,
					),
					cassandra.NewAuth("", ""),
					cassandra.NewDatabase(
						"testkeyspace",
						consistency.Quorum,
					),
				),
				app.NewConfig(
					user.NewConfig(
						session.NewConfig(0),
					),
					"",
					0,
				),
			),
			wantErr:         errs.ErrInvalidConfig,
			wantErrContains: "mongo",
		},
		{
			name: "invalid cassandra config",
			c: impl.NewConfig(
				redis.NewConfig(
					"localhost:6379",
					"",
					0,
				),
				mongo.NewConfig(
					"testdb",
					"testuser",
					"testpass",
					"localhost",
					27017,
				),
				cassandra.NewConfig(
					cassandra.NewConnection(
						[]string{""},
						0,
					),
					cassandra.NewAuth("", ""),
					cassandra.NewDatabase(
						"",
						consistency.Quorum,
					),
				),
				app.NewConfig(
					user.NewConfig(
						session.NewConfig(1*time.Second),
					),
					"localhost",
					8080,
				),
			),
			wantErr:         errs.ErrInvalidConfig,
			wantErrContains: "cassandra",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.c.Validate()
			console.Log("got %v", err)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Contains(t, err.Error(), tt.wantErrContains)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
