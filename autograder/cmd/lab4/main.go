package main

import (
	"context"
	"os"
	"time"

	eventsrq "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/get/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/rq/pagination"
	usersrq "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/users/list/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event/category"
	"github.com/sitnikovik/ndbx/autograder/internal/app/money"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder"
	logoutEndpoint "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/auth/logout/ok/endpoint"
	authEndpoint "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/auth/ok/endpoint"
	createEventEndpoint "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/events/create/ok/endpoint"
	mongoSetup "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/setup/mongo"
	redisSetup "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/setup/redis"
	signupEndpoint "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/sign-up/ok/endpoint"
	mongoTeardown "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/teardown/mongo"
	redisTeardown "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/teardown/redis"
	bulkEventCreation "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab4/job/events/create/ok/mongo"
	listEventsByEndpoint "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab4/job/events/list/by/all/endpoint"
	getNXEventEndpoint "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab4/job/events/one/not-found/endpoint"
	getEventEndpoint "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab4/job/events/one/ok/endpoint"
	updateEventEndpoint "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab4/job/events/update/ok/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
	"github.com/sitnikovik/ndbx/autograder/internal/client/httpx"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo"
	"github.com/sitnikovik/ndbx/autograder/internal/client/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/config/lab3/config"
	"github.com/sitnikovik/ndbx/autograder/internal/console"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	createOneEventMongo "github.com/sitnikovik/ndbx/autograder/internal/step/events/create/one/mongo"
	createOneUserMongo "github.com/sitnikovik/ndbx/autograder/internal/step/user/create/one/mongo"
	listUsersEndpoint "github.com/sitnikovik/ndbx/autograder/internal/step/user/list/by/endpoint/ok"
	getNXUserEndpoint "github.com/sitnikovik/ndbx/autograder/internal/step/user/one/endpoint/not-found"
	getUserEndpoint "github.com/sitnikovik/ndbx/autograder/internal/step/user/one/endpoint/ok"
	listUserEventsEndpoint "github.com/sitnikovik/ndbx/autograder/internal/step/user/one/events/endpoint"
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
	httpcli := httpx.NewClient(httpx.WithEmptyCookieJar())
	baseURL := cfg.App().Address()
	sessionTTL := cfg.App().User().Session().TTL()
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
				user.NewIdentity(samSepiol.ID()),
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
				user.NewIdentity(samSepiol.ID()),
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
				user.NewIdentity(johnDoe.ID()),
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
		event.NewEvent(
			event.NewID("4"),
			event.NewContent(
				"В стране чудес",
				"Платная выставка картин и иллюстраций Екатерины Ващинской",
				event.WithCategory(category.Exhibition),
			),
			event.NewLocation(
				"Москва. Ходынский бульвар 20а",
				event.WithCity("Москва"),
			),
			event.NewCreated(
				timex.MustRFC3339("2026-01-01T11:33:00Z"),
				user.NewIdentity(johnDoe.ID()),
			),
			event.NewDates(
				timex.MustRFC3339("2026-03-24T19:00:00Z"),
				timex.MustRFC3339("2026-03-24T23:59:59Z"),
			),
			event.WithCosts(
				event.NewCosts(
					money.NewMoney(3000, 00),
				),
			),
		),
		event.NewEvent(
			event.NewID("5"),
			event.NewContent(
				"В стране чудес",
				"Выставка картин и иллюстраций Екатерины Ващинской",
				event.WithCategory(category.Exhibition),
			),
			event.NewLocation(
				"Москва. Ходынский бульвар 20а",
				event.WithCity("Москва"),
			),
			event.NewCreated(
				timex.MustRFC3339("2026-01-01T11:33:00Z"),
				user.NewIdentity(samSepiol.ID()),
			),
			event.NewDates(
				timex.MustRFC3339("2026-03-25T10:00:00Z"),
				timex.MustRFC3339("2026-03-25T18:00:00Z"),
			),
			event.WithCosts(
				event.NewCosts(
					money.NewMoney(0, 00),
				),
			),
		),
	}
	veniceClassicsEvents := []event.Event{
		event.NewEvent(
			event.NewID("6"),
			event.NewContent(
				"Концерт «Вечер венской классики»",
				"Приглашаем в библиотеку Культурного центра ЗИЛ на концерт «Вечер венской классики».",
				event.WithCategory(category.Concert),
			),
			event.NewLocation(
				"Культурный центр ЗИЛ",
				event.WithCity("Москва"),
			),
			event.NewCreated(
				timex.MustRFC3339("2025-10-10T18:25:00Z"),
				user.NewIdentity(alexSmith.ID()),
			),
			event.NewDates(
				timex.MustRFC3339("2026-03-25T18:00:00Z"),
				timex.MustRFC3339("2026-03-25T21:00:00Z"),
			),
			event.WithCosts(
				event.NewCosts(
					money.NewMoney(2000, 00),
				),
			),
		),
	}
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
			logoutEndpoint.NewStep(
				httpcli,
				baseURL,
			),
			bulkEventCreation.NewStep(
				mongocli,
			),
			createOneUserMongo.NewStep(
				mongocli,
				samSepiol,
			),
			createOneUserMongo.NewStep(
				mongocli,
				johnDoe,
			),
			createOneUserMongo.NewStep(
				mongocli,
				alexSmith,
			),
			createOneUserMongo.NewStep(
				mongocli,
				johnSmith,
			),
			createOneUserMongo.NewStep(
				mongocli,
				samwiseGamgee,
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
			createOneEventMongo.NewStep(
				mongocli,
				wonderLandEvents[3],
			),
			createOneEventMongo.NewStep(
				mongocli,
				wonderLandEvents[4],
			),
			createOneEventMongo.NewStep(
				mongocli,
				veniceClassicsEvents[0],
			),
			listEventsByEndpoint.NewStep(
				httpcli,
				baseURL,
			),
			getEventEndpoint.NewStep(
				httpcli,
				baseURL,
			),
			getNXEventEndpoint.NewStep(
				httpcli,
				baseURL,
			),
			listUsersEndpoint.NewStep(
				httpcli,
				baseURL,
				usersrq.NewBody(
					usersrq.WithFullName("sam"),
				),
				[]user.User{
					samSepiol,
					samwiseGamgee,
				},
			),
			listUsersEndpoint.NewStep(
				httpcli,
				baseURL,
				usersrq.NewBody(
					usersrq.WithFullName("sam"),
					usersrq.WithPagination(
						pagination.NewPagination(1, 1),
					),
				),
				[]user.User{
					samwiseGamgee,
				},
			),
			listUsersEndpoint.NewStep(
				httpcli,
				baseURL,
				usersrq.NewBody(
					usersrq.WithFullName("Smith"),
					usersrq.WithPagination(
						pagination.NewPagination(1, 0),
					),
				),
				[]user.User{
					alexSmith,
				},
			),
			listUsersEndpoint.NewStep(
				httpcli,
				baseURL,
				usersrq.NewBody(
					usersrq.WithFullName("John"),
				),
				[]user.User{
					johnDoe,
					johnSmith,
				},
			),
			listUsersEndpoint.NewStep(
				httpcli,
				baseURL,
				usersrq.NewBody(
					usersrq.WithIdentity(
						user.NewIdentity(
							user.NewID("2"),
						),
					),
				),
				[]user.User{
					johnDoe,
				},
			),
			getUserEndpoint.NewStep(
				httpcli,
				baseURL,
				samSepiol.ID(),
				samSepiol,
			),
			getNXUserEndpoint.NewStep(
				httpcli,
				baseURL,
				user.NewID("123iuj2ekwo"),
			),
			listUserEventsEndpoint.NewStep(
				httpcli,
				baseURL,
				samSepiol.ID(),
				eventsrq.NewBody(
					eventsrq.WithDates(
						timex.MustRFC3339("2026-03-25T00:00:00Z"),
						time.Time{},
					),
				),
				[]event.Event{
					wonderLandEvents[4],
				},
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
