package main

import (
	"context"
	"os"

	eventsrq "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/get/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/rq/include"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event/category"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event/reaction"
	"github.com/sitnikovik/ndbx/autograder/internal/app/money"
	"github.com/sitnikovik/ndbx/autograder/internal/app/reaction/count"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder"
	mongoSetup "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/setup/mongo"
	redisSetup "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/setup/redis"
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
	login "github.com/sitnikovik/ndbx/autograder/internal/step/user/auth/login"
	logout "github.com/sitnikovik/ndbx/autograder/internal/step/user/auth/logout"
	signup "github.com/sitnikovik/ndbx/autograder/internal/step/user/create/sign-up"
	listUserEventsEndpoint "github.com/sitnikovik/ndbx/autograder/internal/step/user/one/events/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/step/user/one/events/expect"
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
	alexSmith := userfx.NewAlexSmith()
	johnSmith := userfx.NewJohnSmith()
	samwiseGamgee := userfx.NewSamwiseGamgee()
	pwd := "supa_dxpa_pwd"
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
			signup.NewStep(
				httpcli,
				mongocli,
				baseURL,
				samSepiol,
				pwd,
			),
			signup.NewStep(
				httpcli,
				mongocli,
				baseURL,
				johnDoe,
				pwd,
			),
			signup.NewStep(
				httpcli,
				mongocli,
				baseURL,
				alexSmith,
				pwd,
			),
			signup.NewStep(
				httpcli,
				mongocli,
				baseURL,
				johnSmith,
				pwd,
			),
			signup.NewStep(
				httpcli,
				mongocli,
				baseURL,
				samwiseGamgee,
				pwd,
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
			login.NewStep(
				httpcli,
				baseURL,
				samSepiol,
				pwd,
			),
			likeOneEventEndpoint.NewStep(
				httpcli,
				baseURL,
				wonderLandEvents[0],
			),
			likeOneEventEndpoint.NewStep(
				httpcli,
				baseURL,
				wonderLandEvents[1],
			),
			likeOneEventEndpoint.NewStep(
				httpcli,
				baseURL,
				wonderLandEvents[2],
			),
			logout.NewStep(
				httpcli,
				baseURL,
			),
			login.NewStep(
				httpcli,
				baseURL,
				johnDoe,
				pwd,
			),
			likeOneEventEndpoint.NewStep(
				httpcli,
				baseURL,
				wonderLandEvents[0],
			),
			logout.NewStep(
				httpcli,
				baseURL,
			),
			login.NewStep(
				httpcli,
				baseURL,
				alexSmith,
				pwd,
			),
			likeOneEventEndpoint.NewStep(
				httpcli,
				baseURL,
				wonderLandEvents[0],
			),
			likeOneEventEndpoint.NewStep(
				httpcli,
				baseURL,
				wonderLandEvents[2],
			),
			logout.NewStep(
				httpcli,
				baseURL,
			),
			login.NewStep(
				httpcli,
				baseURL,
				samwiseGamgee,
				pwd,
			),
			likeOneEventEndpoint.NewStep(
				httpcli,
				baseURL,
				wonderLandEvents[0],
			),
			logout.NewStep(
				httpcli,
				baseURL,
			),
			getEventLikesRedis.NewStep(
				rediscli,
				wonderLandEvents[0],
				4,
				reactonTTL,
			),
			getEventLikesRedis.NewStep(
				rediscli,
				wonderLandEvents[1],
				4,
				reactonTTL,
			),
			getEventLikesRedis.NewStep(
				rediscli,
				wonderLandEvents[2],
				4,
				reactonTTL,
			),
			eventLikesCassandra.NewStep(
				cassandracli,
				wonderLandEvents[0],
				2,
			),
			eventLikesCassandra.NewStep(
				cassandracli,
				wonderLandEvents[1],
				1,
			),
			eventLikesCassandra.NewStep(
				cassandracli,
				wonderLandEvents[2],
				1,
			),
			listUserEventsEndpoint.NewStep(
				httpcli,
				baseURL,
				samSepiol,
				eventsrq.NewBody(
					eventsrq.WithInclude(
						include.NewInclude("reactions"),
					),
				),
				[]event.Event{
					wonderLandEvents[0],
					wonderLandEvents[1],
				},
				listUserEventsEndpoint.WithExpectations(
					expect.NewExpectations(
						expect.WithReactions(
							[]reaction.Reactions{
								reaction.NewReactions(
									reaction.WithCounts(
										count.NewCounts(
											count.WithLikes(4),
											count.WithDislikes(0),
										),
									),
								),
								reaction.NewReactions(
									reaction.WithCounts(
										count.NewCounts(
											count.WithLikes(4),
											count.WithDislikes(0),
										),
									),
								),
							},
						),
					),
				),
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
		console.Fatal("Lab 5 autograder failed: %v", err)
	}
}
