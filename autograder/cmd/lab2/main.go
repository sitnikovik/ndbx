package main

import (
	"context"
	"os"
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/autograder"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/expire"
	expireRedisStep "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/expire/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/health"
	healthEndpointStep "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/health/endpoint"
	healthRedisStep "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/health/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/preserve"
	preserveEndpointStep "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/preserve/endpoint"
	preserveRedisStep "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/preserve/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/refresh"
	refreshEndpointStep "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/refresh/endpoint"
	refreshRedisStep "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/refresh/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/session"
	sessionEndpointStep "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/session/endpoint"
	sessionRedisStep "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/session/redis"
	setup "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/setup/redis"
	teardown "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/teardown/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/client/httpx"
	"github.com/sitnikovik/ndbx/autograder/internal/client/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/config"
	"github.com/sitnikovik/ndbx/autograder/internal/console"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	"github.com/sitnikovik/ndbx/autograder/internal/step/wait"
)

// main is the entry point for the Lab 2 autograder.
func main() {
	defer func() {
		if r := recover(); r != nil {
			console.Panic(r)
			os.Exit(1)
		}
	}()
	cfg := config.MustLoad()
	rediscli := redis.NewClient(cfg.Redis())
	httpcli := httpx.NewClient(httpx.WithEmptyCookieJar())
	baseURL := cfg.App().Address()
	sessionTTL := cfg.App().User().Session().TTL()
	ctx := context.Background()
	vars := step.NewVariables()
	err := autograder.
		NewAutograder(
			setup.NewJob(
				rediscli,
			),
			health.NewJob(
				healthEndpointStep.NewStep(httpcli, baseURL),
				healthRedisStep.NewStep(rediscli),
			),
			session.NewJob(
				sessionEndpointStep.NewStep(httpcli, baseURL),
				sessionRedisStep.NewStep(rediscli),
			),
			preserve.NewJob(
				preserveEndpointStep.NewStep(httpcli, baseURL),
				preserveRedisStep.NewStep(rediscli),
			),
			refresh.NewJob(
				refreshEndpointStep.NewStep(httpcli, baseURL),
				refreshRedisStep.NewStep(rediscli),
			),
			expire.NewJob(
				wait.NewWaitingStep(sessionTTL+5*time.Second),
				expireRedisStep.NewStep(rediscli),
			),
			session.NewJob(
				sessionEndpointStep.NewStep(httpcli, baseURL),
				sessionRedisStep.NewStep(rediscli),
			),
			teardown.NewJob(
				rediscli,
			),
		).
		Run(ctx, vars)
	if err != nil {
		console.Fatal("Lab 2 autograder failed: %v", err)
	}
}
