package main

import (
	"context"
	"os"

	"github.com/sitnikovik/ndbx/autograder/internal/autograder"
	authEndpoint "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/auth/ok/endpoint"
	createEventEndpoint "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/events/create/ok/endpoint"
	mongoSetup "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/setup/mongo"
	redisSetup "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/setup/redis"
	signupEndpoint "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/sign-up/ok/endpoint"
	mongoTeardown "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/teardown/mongo"
	redisTeardown "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/teardown/redis"
	bulkEventCreation "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab4/job/events/create/ok/mongo"
	updateEventEndpoint "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab4/job/events/update/ok/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
	"github.com/sitnikovik/ndbx/autograder/internal/client/httpx"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo"
	"github.com/sitnikovik/ndbx/autograder/internal/client/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/config/lab3/config"
	"github.com/sitnikovik/ndbx/autograder/internal/console"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// main is the entry point for the Lab 4ы autograder.
func main() {
	defer func() {
		if r := recover(); r != nil {
			console.Panic(r)
			os.Exit(1)
		}
	}()
	cfg := config.MustLoad()
	mongocli := mongo.NewClient(cfg.Mongo())
	rediscli := redis.NewClient(cfg.Redis())
	httpcli := httpx.NewClient(httpx.WithEmptyCookieJar())
	baseURL := cfg.App().Address()
	sessionTTL := cfg.App().User().Session().TTL()
	ctx := context.Background()
	vars := step.NewVariables()
	vars.Set(
		variable.SessionTTL,
		sessionTTL,
	)
	err := autograder.
		NewAutograder(
			mongoSetup.NewStep(
				mongocli,
			),
			redisSetup.NewStep(
				rediscli,
			),
			signupEndpoint.NewStep(
				httpcli,
				baseURL,
			),
			authEndpoint.NewStep(
				httpcli,
				baseURL,
			),
			createEventEndpoint.NewStep(
				httpcli,
				baseURL,
			),
			updateEventEndpoint.NewStep(
				httpcli,
				baseURL,
			),
			bulkEventCreation.NewStep(
				mongocli,
			),
			mongoTeardown.NewStep(
				mongocli,
			),
			redisTeardown.NewStep(
				rediscli,
			),
		).
		Run(ctx, vars)
	if err != nil {
		console.Fatal("Lab 4 autograder failed: %v", err)
	}
}
