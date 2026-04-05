package main

import (
	"context"
	"os"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event/category"
	"github.com/sitnikovik/ndbx/autograder/internal/app/money"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder"
	logoutEndpoint "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/auth/logout/ok/endpoint"
	authEndpoint "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/auth/ok/endpoint"
	mongoSetup "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/setup/mongo"
	redisSetup "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/setup/redis"
	signupEndpoint "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/sign-up/ok/endpoint"
	mongoTeardown "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/teardown/mongo"
	redisTeardown "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/teardown/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
	"github.com/sitnikovik/ndbx/autograder/internal/client/cassandra"
	"github.com/sitnikovik/ndbx/autograder/internal/client/httpx"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo"
	"github.com/sitnikovik/ndbx/autograder/internal/client/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/config/lab5/config"
	"github.com/sitnikovik/ndbx/autograder/internal/console"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	cassandraSetup "github.com/sitnikovik/ndbx/autograder/internal/step/cassandra/setup"
	cassandraTeardown "github.com/sitnikovik/ndbx/autograder/internal/step/cassandra/teardown"
	createOneEventMongo "github.com/sitnikovik/ndbx/autograder/internal/step/events/create/one/mongo"
	likeOneEventEndpoint "github.com/sitnikovik/ndbx/autograder/internal/step/events/one/like/endpoint"
	getEventLikesRedis "github.com/sitnikovik/ndbx/autograder/internal/step/events/one/like/redis"
	eventLikesCassandra "github.com/sitnikovik/ndbx/autograder/internal/step/reaction/event/like/list/cassandra"
	createOneUserMongo "github.com/sitnikovik/ndbx/autograder/internal/step/user/create/one/mongo"
	userfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

// main is the entry point for the Lab 4 autograder.
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
	cassandracli := cassandra.NewClient(cfg.Cassandra())
	_ = cassandracli
	httpcli := httpx.NewClient(httpx.WithEmptyCookieJar())
	baseURL := cfg.App().Address()
	sessionTTL := cfg.App().User().Session().TTL()
	reactonTTL := cfg.App().Event().Reactions().TTL()
	ctx := context.Background()
	vars := step.NewVariables()
	vars.Set(
		variable.SessionTTL,
		sessionTTL,
	)
	samSepiol := userfx.NewSamSepiol()
	johnDoe := userfx.NewJohnDoe()
	wonderLandEvents := []event.Event{
		event.NewEvent(
			event.NewID("1"),
			event.NewContent(
				"В стране чудес",
				"Бесплатная выставка картин и иллюстраций Екатерины Ващинской",
				event.WithCategory(category.Exhibition),
			),
			event.NewLocation(
				"Москва. Ходынский бульвар 20а",
				event.WithCity("Москва"),
			),
			event.NewCreated(
				timex.MustRFC3339("2026-01-01T11:33:00Z"),
				samSepiol.Idendity(),
			),
			event.NewDates(
				timex.MustRFC3339("2026-03-24T10:00:00Z"),
				timex.MustRFC3339("2026-03-24T12:00:00Z"),
			),
			event.WithCosts(
				event.NewCosts(
					money.NewMoney(0, 00),
				),
			),
		),
		event.NewEvent(
			event.NewID("2"),
			event.NewContent(
				"В стране чудес",
				"Бесплатная выставка картин и иллюстраций Екатерины Ващинской",
				event.WithCategory(category.Exhibition),
			),
			event.NewLocation(
				"Москва. Ходынский бульвар 20а",
			),
			event.NewCreated(
				timex.MustRFC3339("2026-01-01T11:33:00Z"),
				samSepiol.Idendity(),
			),
			event.NewDates(
				timex.MustRFC3339("2026-03-24T12:00:00Z"),
				timex.MustRFC3339("2026-03-24T14:00:00Z"),
			),
			event.WithCosts(
				event.NewCosts(
					money.NewMoney(0, 00),
				),
			),
		),
		event.NewEvent(
			event.NewID("3"),
			event.NewContent(
				"В стране чудес",
				"Бесплатная выставка картин и иллюстраций Екатерины Ващинской",
				event.WithCategory(category.Exhibition),
			),
			event.NewLocation(
				"Москва. Ходынский бульвар 20а",
				event.WithCity("Москва"),
			),
			event.NewCreated(
				timex.MustRFC3339("2026-01-01T11:33:00Z"),
				johnDoe.Idendity(),
			),
			event.NewDates(
				timex.MustRFC3339("2026-03-24T14:00:00Z"),
				timex.MustRFC3339("2026-03-24T16:00:00Z"),
			),
			event.WithCosts(
				event.NewCosts(
					money.NewMoney(0, 00),
				),
			),
		),
	}
	likeEventsEndpoints := make([]step.Runner, 0, 24)
	for range 14 {
		likeEventsEndpoints = append(
			likeEventsEndpoints,
			likeOneEventEndpoint.NewStep(
				httpcli,
				baseURL,
				wonderLandEvents[0],
			),
		)
	}
	for range 8 {
		likeEventsEndpoints = append(
			likeEventsEndpoints,
			likeOneEventEndpoint.NewStep(
				httpcli,
				baseURL,
				wonderLandEvents[1],
			),
		)
	}
	for range 2 {
		likeEventsEndpoints = append(
			likeEventsEndpoints,
			likeOneEventEndpoint.NewStep(
				httpcli,
				baseURL,
				wonderLandEvents[2],
			),
		)
	}
	err := autograder.
		NewAutograder(
			cassandraSetup.NewStep(
				cassandracli,
			),
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
			logoutEndpoint.NewStep(
				httpcli,
				baseURL,
			),
			createOneUserMongo.NewStep(
				mongocli,
				samSepiol,
			),
			createOneUserMongo.NewStep(
				mongocli,
				johnDoe,
			),
			createOneEventMongo.NewStep(
				mongocli,
				wonderLandEvents[0],
			),
			createOneEventMongo.NewStep(
				mongocli,
				wonderLandEvents[1],
			),
			createOneEventMongo.NewStep(
				mongocli,
				wonderLandEvents[2],
			),
			step.NewList(
				likeEventsEndpoints...,
			),
			getEventLikesRedis.NewStep(
				rediscli,
				wonderLandEvents[0],
				24,
				reactonTTL,
			),
			getEventLikesRedis.NewStep(
				rediscli,
				wonderLandEvents[1],
				24,
				reactonTTL,
			),
			getEventLikesRedis.NewStep(
				rediscli,
				wonderLandEvents[2],
				24,
				reactonTTL,
			),
			eventLikesCassandra.NewStep(
				cassandracli,
				wonderLandEvents[0],
				14,
			),
			eventLikesCassandra.NewStep(
				cassandracli,
				wonderLandEvents[1],
				8,
			),
			eventLikesCassandra.NewStep(
				cassandracli,
				wonderLandEvents[2],
				2,
			),
			logoutEndpoint.NewStep(
				httpcli,
				baseURL,
			),
			mongoTeardown.NewStep(
				mongocli,
			),
			redisTeardown.NewStep(
				rediscli,
			),
			cassandraTeardown.NewStep(
				cassandracli,
			),
		).
		Run(ctx, vars)
	if err != nil {
		console.Fatal("Lab 4 autograder failed: %v", err)
	}
}
