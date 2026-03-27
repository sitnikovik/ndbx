package main

import (
	"context"
	"os"

	"github.com/sitnikovik/ndbx/autograder/internal/autograder"
	authFailedEndpoint "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/auth/inv-creds/endpoint"
	logoutEndpoint "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/auth/logout/ok/endpoint"
	logoutRedis "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/auth/logout/ok/redis"
	logoutUnauthEndpoint "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/auth/logout/unauth/endpoint"
	authEndpoint "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/auth/ok/endpoint"
	createEventEndpoint "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/events/create/ok/endpoint"
	createNewEventMongo "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/events/create/ok/mongo"
	createNewEventRedis "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/events/create/ok/redis"
	createEventUnauthEndpoint "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/events/create/unauth/endpoint"
	listAllEndpoint "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/events/list/ok/all/endpoint"
	listByTitleEndpoint "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/events/list/ok/by/title/endpoint"
	mongoSetup "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/setup/mongo"
	redisSetup "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/setup/redis"
	signupEmptyPwdEndpoint "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/sign-up/empty-pwd/endpoint"
	signupEndpoint "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/sign-up/ok/endpoint"
	signupMongo "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/sign-up/ok/mongo"
	signupRedis "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/sign-up/ok/redis"
	mongoTeardown "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/teardown/mongo"
	redisTeardown "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/teardown/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
	"github.com/sitnikovik/ndbx/autograder/internal/client/httpx"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo"
	"github.com/sitnikovik/ndbx/autograder/internal/client/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/config/lab3/config"
	"github.com/sitnikovik/ndbx/autograder/internal/console"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	"github.com/sitnikovik/ndbx/autograder/internal/step/wait"
)

// main is the entry point for the Lab 3 autograder.
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
			createEventUnauthEndpoint.NewStep(
				httpcli,
				baseURL,
			),
			logoutUnauthEndpoint.NewStep(
				httpcli,
				baseURL,
			),
			signupEmptyPwdEndpoint.NewStep(
				httpcli,
				baseURL,
			),
			signupEndpoint.NewStep(
				httpcli,
				baseURL,
			),
			signupMongo.NewStep(
				mongocli,
			),
			signupRedis.NewStep(
				rediscli,
			),
			authFailedEndpoint.NewStep(
				httpcli,
				baseURL,
			),
			authEndpoint.NewStep(
				httpcli,
				baseURL,
			),
			wait.NewWaitingStep(
				sessionTTL/2,
			),
			createEventEndpoint.NewStep(
				httpcli,
				baseURL,
			),
			createNewEventMongo.NewStep(
				mongocli,
			),
			createNewEventRedis.NewStep(
				rediscli,
			),
			wait.NewWaitingStep(
				sessionTTL/2,
			),
			logoutEndpoint.NewStep(
				httpcli,
				baseURL,
			),
			logoutRedis.NewStep(
				rediscli,
			),
			listAllEndpoint.NewStep(
				httpcli,
				baseURL,
			),
			listByTitleEndpoint.NewStep(
				httpcli,
				baseURL,
			),
			createEventUnauthEndpoint.NewStep(
				httpcli,
				baseURL,
			),
			listAllEndpoint.NewStep(
				httpcli,
				baseURL,
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
		console.Fatal("Lab 3 autograder failed: %v", err)
	}
}
