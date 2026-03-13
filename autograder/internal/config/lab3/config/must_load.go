package config

import (
	"github.com/sitnikovik/ndbx/autograder/internal/config/app"
	"github.com/sitnikovik/ndbx/autograder/internal/config/mongo"
	"github.com/sitnikovik/ndbx/autograder/internal/config/redis"
)

// MustLoad loads the job configuration
// and panics if any required configuration is missing or invalid.
//
// It is intended to be used in situations
// where the configuration must be valid for the application to run,
// such as in the main function of the application.
func MustLoad() Config {
	cfg := NewConfig(
		redis.MustLoad(),
		mongo.MustLoad(),
		app.MustLoad(),
	)
	err := cfg.Validate()
	if err != nil {
		panic(err)
	}
	return cfg
}
