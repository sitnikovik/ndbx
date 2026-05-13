package config_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/config/app"
	"github.com/sitnikovik/ndbx/autograder/internal/config/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/config/app/user/session"
	impl "github.com/sitnikovik/ndbx/autograder/internal/config/lab7/config"
	"github.com/sitnikovik/ndbx/autograder/internal/config/mongo"
	"github.com/sitnikovik/ndbx/autograder/internal/config/neo4j"
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
				mongo.NewConfig(
					"testdb",
					"testuser",
					"testpass",
					"localhost",
					27017,
				),
				neo4j.NewConfig(
					neo4j.NewConnection("neo4j://localhost:7687"),
					neo4j.NewAuth("user", "password"),
				),
				app.NewConfig(
					user.NewConfig(
						session.NewConfig(1*time.Second),
					),
					"localhost",
					8080,
				),
			),
			wantErr: nil,
		},
		{
			name: "invalid app config",
			c: impl.NewConfig(
				mongo.NewConfig(
					"testdb",
					"testuser",
					"testpass",
					"localhost",
					27017,
				),
				neo4j.NewConfig(
					neo4j.NewConnection("neo4j://localhost:7687"),
					neo4j.NewAuth("user", "password"),
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
			name: "invalid mongo config",
			c: impl.NewConfig(
				mongo.NewConfig(
					"",
					"",
					"",
					"",
					0,
				),
				neo4j.NewConfig(
					neo4j.NewConnection("neo4j://localhost:7687"),
					neo4j.NewAuth("user", "password"),
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
			name: "invalid neo4j config",
			c: impl.NewConfig(
				mongo.NewConfig(
					"testdb",
					"testuser",
					"testpass",
					"localhost",
					27017,
				),
				neo4j.NewConfig(
					neo4j.NewConnection(""),
					neo4j.NewAuth("user", "password"),
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
			wantErrContains: "neo4j",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.c.Validate()
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Contains(t, err.Error(), tt.wantErrContains)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
