package main

import (
	"context"
	"os"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event/category"
	"github.com/sitnikovik/ndbx/autograder/internal/app/money"
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
	"github.com/sitnikovik/ndbx/autograder/internal/config/lab7/config"
	"github.com/sitnikovik/ndbx/autograder/internal/console"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
	respXpct "github.com/sitnikovik/ndbx/autograder/internal/expect/http/response/expectation"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	cassandraSetup "github.com/sitnikovik/ndbx/autograder/internal/step/cassandra/setup"
	cassandraTeardown "github.com/sitnikovik/ndbx/autograder/internal/step/cassandra/teardown"
	createOneEventMongo "github.com/sitnikovik/ndbx/autograder/internal/step/events/create/one/mongo"
	dislikeOneEventEndpoint "github.com/sitnikovik/ndbx/autograder/internal/step/events/one/dislike/endpoint"
	likeOneEventEndpoint "github.com/sitnikovik/ndbx/autograder/internal/step/events/one/like/endpoint"
	login "github.com/sitnikovik/ndbx/autograder/internal/step/user/auth/login"
	logout "github.com/sitnikovik/ndbx/autograder/internal/step/user/auth/logout"
	signup "github.com/sitnikovik/ndbx/autograder/internal/step/user/create/sign-up"
	recommsEndpoint "github.com/sitnikovik/ndbx/autograder/internal/step/user/one/recommendations/endpoint"
	recommsEndpointXpct "github.com/sitnikovik/ndbx/autograder/internal/step/user/one/recommendations/endpoint/expect"
	recommsRedis "github.com/sitnikovik/ndbx/autograder/internal/step/user/one/recommendations/redis"
	recommsRedisXpct "github.com/sitnikovik/ndbx/autograder/internal/step/user/one/recommendations/redis/expect"
	userfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

// main is the entry point for the Lab 7 autograder.
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
	recommsTTL := cfg.App().User().Recommendations().Events().TTL()
	ctx := context.Background()
	vars := step.NewVariables()
	vars.Set(variable.SessionTTL, sessionTTL)
	// Пользователи
	samSepiol := userfx.NewSamSepiol()
	johnDoe := userfx.NewJohnDoe()
	alexSmith := userfx.NewAlexSmith()
	johnSmith := userfx.NewJohnSmith()
	samwiseGamgee := userfx.NewSamwiseGamgee()
	pwd := "supa_dxpa_pwd"
	// Мепроприятия
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
			event.WithCosts(event.NewCosts(money.NewMoney(0, 00))),
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
				event.WithCity("Москва"),
			),
			event.NewCreated(
				timex.MustRFC3339("2026-01-01T11:34:00Z"),
				samSepiol.Idendity(),
			),
			event.NewDates(
				timex.MustRFC3339("2026-03-24T12:00:00Z"),
				timex.MustRFC3339("2026-03-24T14:00:00Z"),
			),
			event.WithCosts(event.NewCosts(money.NewMoney(0, 00))),
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
				timex.MustRFC3339("2026-01-01T11:35:00Z"),
				johnDoe.Idendity(),
			),
			event.NewDates(
				timex.MustRFC3339("2026-03-24T14:00:00Z"),
				timex.MustRFC3339("2026-03-24T16:00:00Z"),
			),
			event.WithCosts(event.NewCosts(money.NewMoney(0, 00))),
		),
	}
	concertEvent := event.NewEvent(
		event.NewID("4"),
		event.NewContent(
			"Ночь в опере",
			"Гала-концерт лучших солистов Большого театра",
			event.WithCategory(category.Concert),
		),
		event.NewLocation(
			"Москва, ул. Большая Дмитровка, д. 6",
			event.WithCity("Москва"),
		),
		event.NewCreated(
			timex.MustRFC3339("2026-01-02T10:00:00Z"),
			alexSmith.Idendity(),
		),
		event.NewDates(
			timex.MustRFC3339("2026-04-05T19:00:00Z"),
			timex.MustRFC3339("2026-04-05T21:00:00Z"),
		),
		event.WithCosts(event.NewCosts(money.NewMoney(2500, 00))),
	)
	shakespeareMeetup := event.NewEvent(
		event.NewID("5"),
		event.NewContent(
			"Обсуждение Гамлета",
			"Встреча поклонников Шекспира: анализ пьесы, дискуссия, читка сцены",
			event.WithCategory(category.Meetup),
		),
		event.NewLocation(
			"Москва, Новый Арбат, д. 22, библиотека им. Ленина",
			event.WithCity("Москва"),
		),
		event.NewCreated(
			timex.MustRFC3339("2026-01-03T15:00:00Z"),
			johnDoe.Idendity(),
		),
		event.NewDates(
			timex.MustRFC3339("2026-04-10T18:30:00Z"),
			timex.MustRFC3339("2026-04-10T20:30:00Z"),
		),
		event.WithCosts(event.NewCosts(money.NewMoney(0, 00))),
	)
	err := autograder.NewAutograder(
		// Настройка инфраструктуры
		cassandraSetup.NewStep(cassandracli),
		mongoSetup.NewStep(mongocli),
		redisSetup.NewStep(rediscli),

		// Регистрация пользователей
		signup.NewStep(httpcli, mongocli, baseURL, samSepiol, pwd),
		signup.NewStep(httpcli, mongocli, baseURL, johnDoe, pwd),
		signup.NewStep(httpcli, mongocli, baseURL, alexSmith, pwd),
		signup.NewStep(httpcli, mongocli, baseURL, johnSmith, pwd),
		signup.NewStep(httpcli, mongocli, baseURL, samwiseGamgee, pwd),

		// Create events
		createOneEventMongo.NewStep(mongocli, wonderLandEvents[0]),
		createOneEventMongo.NewStep(mongocli, wonderLandEvents[1]),
		createOneEventMongo.NewStep(mongocli, wonderLandEvents[2]),
		createOneEventMongo.NewStep(mongocli, concertEvent),
		createOneEventMongo.NewStep(mongocli, shakespeareMeetup),

		// Sam Sepiol likes exhibition
		login.NewStep(httpcli, baseURL, samSepiol, pwd),
		likeOneEventEndpoint.NewStep(httpcli, baseURL, wonderLandEvents[0]),
		likeOneEventEndpoint.NewStep(httpcli, baseURL, wonderLandEvents[1]),
		likeOneEventEndpoint.NewStep(httpcli, baseURL, wonderLandEvents[2]),
		logout.NewStep(httpcli, baseURL),

		// John Doe likes exhibition and concert
		login.NewStep(httpcli, baseURL, johnDoe, pwd),
		likeOneEventEndpoint.NewStep(httpcli, baseURL, wonderLandEvents[0]),
		likeOneEventEndpoint.NewStep(httpcli, baseURL, concertEvent),
		logout.NewStep(httpcli, baseURL),

		// Alex Smith likes exhibition and concert
		login.NewStep(httpcli, baseURL, alexSmith, pwd),
		likeOneEventEndpoint.NewStep(httpcli, baseURL, wonderLandEvents[0]),
		likeOneEventEndpoint.NewStep(httpcli, baseURL, wonderLandEvents[2]),
		likeOneEventEndpoint.NewStep(httpcli, baseURL, concertEvent),
		logout.NewStep(httpcli, baseURL),

		// Samwise Gamgee likes the exhibition (then dislikes), and joins the Shakespeare meetup
		login.NewStep(httpcli, baseURL, samwiseGamgee, pwd),
		likeOneEventEndpoint.NewStep(httpcli, baseURL, wonderLandEvents[0]),
		dislikeOneEventEndpoint.NewStep(httpcli, baseURL, wonderLandEvents[0]),
		likeOneEventEndpoint.NewStep(httpcli, baseURL, shakespeareMeetup),
		logout.NewStep(httpcli, baseURL),

		// Check recommendations
		login.NewStep(httpcli, baseURL, samwiseGamgee, pwd),
		recommsEndpoint.NewStep(
			step.NewDesc(
				"User's recommendations endpoint",
				"Checking recommendations for Samwise Gamgee. "+
					"Likes build the recommendation graph, dislikes are ignored in this lab. "+
					"Samwise liked the exhibition before disliking it, so the graph still connects him to other users who liked it. "+
					"Those users also liked other exhibition events and the concert, so after deduplication by title and excluding already liked events "+
					"the recommendations are wonderLandEvents[1] and concertEvent",
			),
			httpcli,
			baseURL,
			recommsEndpointXpct.NewExpectations(
				recommsEndpointXpct.WithEvents(
					wonderLandEvents[1],
					concertEvent,
				),
				recommsEndpointXpct.WithResponse(
					respXpct.NewExpectations(
						respXpct.WithAsserts(
							response.AssertOKStatus,
							response.AssertNotEmptyContent,
						),
					),
				),
			),
		),
		recommsRedis.NewStep(
			step.NewDesc(
				"User's recommendations Redis",
				"Recommendations for Samwise Gamgee must not be in Redis cache before first request",
			),
			rediscli,
			samwiseGamgee,
			recommsRedisXpct.NewExpectations(
				recommsRedisXpct.WithNoEvents(),
			),
		),
		recommsEndpoint.NewStep(
			step.NewDesc(
				"User's recommendations endpoint",
				"Checking recommendations for Samwise Gamgee. "+
					"Likes build the recommendation graph, dislikes are ignored in this lab. "+
					"Samwise liked the exhibition before disliking it, so the graph still connects him to other users who liked it. "+
					"Those users also liked other exhibition events and the concert, so after deduplication by title and excluding already liked events "+
					"the recommendations are wonderLandEvents[1] and concertEvent",
			),
			httpcli,
			baseURL,
			recommsEndpointXpct.NewExpectations(
				recommsEndpointXpct.WithEvents(
					wonderLandEvents[1],
					concertEvent,
				),
				recommsEndpointXpct.WithResponse(
					respXpct.NewExpectations(
						respXpct.WithAsserts(
							response.AssertOKStatus,
							response.AssertNotEmptyContent,
						),
					),
				),
			),
		),
		recommsRedis.NewStep(
			step.NewDesc(
				"User's recommendations Redis",
				"Recommendations for Samwise Gamgee must not be in Redis cache before first request",
			),
			rediscli,
			samwiseGamgee,
			recommsRedisXpct.NewExpectations(
				recommsRedisXpct.WithEvents(
					wonderLandEvents[1],
					concertEvent,
				),
				recommsRedisXpct.WithTTL(
					recommsTTL,
				),
			),
		),
		logout.NewStep(httpcli, baseURL),
		login.NewStep(httpcli, baseURL, johnSmith, pwd),
		recommsEndpoint.NewStep(
			step.NewDesc(
				"User's recommendations endpoint",
				"Empty list is expected for John Doe",
			),
			httpcli,
			baseURL,
			recommsEndpointXpct.NewExpectations(
				recommsEndpointXpct.WithNoEvents(),
			),
		),
		logout.NewStep(httpcli, baseURL),
		recommsEndpoint.NewStep(
			step.NewDesc(
				"User's recommendations endpoint",
				"Try to get recommendations on user is authenticated",
			),
			httpcli,
			baseURL,
			recommsEndpointXpct.NewExpectations(
				recommsEndpointXpct.WithResponse(
					respXpct.NewExpectations(
						respXpct.WithAsserts(
							response.AssertUnauthorizedStatus,
							response.AssertEmptyContent,
						),
					),
				),
			),
		),
		mongoTeardown.NewStep(mongocli),
		redisTeardown.NewStep(rediscli),
		cassandraTeardown.NewStep(cassandracli),
	).Run(ctx, vars)
	if err != nil {
		console.Fatal("Lab 7 autograder failed: %v", err)
	}
}
