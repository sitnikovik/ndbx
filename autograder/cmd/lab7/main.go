package main

import (
	"context"
	"os"

	cookieSession "github.com/sitnikovik/ndbx/autograder/internal/app/cookie/session"
	eventsrq "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/get/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/reviews/events/create/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/rq/include"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event/category"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event/review"
	"github.com/sitnikovik/ndbx/autograder/internal/app/money"
	"github.com/sitnikovik/ndbx/autograder/internal/app/rating"
	"github.com/sitnikovik/ndbx/autograder/internal/app/review/count"
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
	cookieassert "github.com/sitnikovik/ndbx/autograder/internal/expect/http/cookie"
	cookiexpct "github.com/sitnikovik/ndbx/autograder/internal/expect/http/cookie/expectation"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response/expectation"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	cassandraSetup "github.com/sitnikovik/ndbx/autograder/internal/step/cassandra/setup"
	cassandraTeardown "github.com/sitnikovik/ndbx/autograder/internal/step/cassandra/teardown"
	createOneEventMongo "github.com/sitnikovik/ndbx/autograder/internal/step/events/create/one/mongo"
	listEventsEndpoint "github.com/sitnikovik/ndbx/autograder/internal/step/events/list/endpoint"
	listEventsExpectations "github.com/sitnikovik/ndbx/autograder/internal/step/events/list/endpoint/expect"
	reviewRedis "github.com/sitnikovik/ndbx/autograder/internal/step/events/one/review/redis"
	reviewRedisExpects "github.com/sitnikovik/ndbx/autograder/internal/step/events/one/review/redis/expect"
	reviewEventEndpoint "github.com/sitnikovik/ndbx/autograder/internal/step/review/event/create/endpoint"
	eventReviewsCassandra "github.com/sitnikovik/ndbx/autograder/internal/step/review/event/list/cassandra"
	eventReviewsCassandraXpct "github.com/sitnikovik/ndbx/autograder/internal/step/review/event/list/cassandra/expectation"
	reviewListEndpoint "github.com/sitnikovik/ndbx/autograder/internal/step/review/event/list/endpoint"
	reviewListEndpointXpct "github.com/sitnikovik/ndbx/autograder/internal/step/review/event/list/endpoint/expect"
	updateReviewEndpoint "github.com/sitnikovik/ndbx/autograder/internal/step/review/event/update/endpoint"
	login "github.com/sitnikovik/ndbx/autograder/internal/step/user/auth/login"
	logout "github.com/sitnikovik/ndbx/autograder/internal/step/user/auth/logout"
	signup "github.com/sitnikovik/ndbx/autograder/internal/step/user/create/sign-up"
	listUserEventsEndpoint "github.com/sitnikovik/ndbx/autograder/internal/step/user/one/events/endpoint"
	listUserEventsExpectations "github.com/sitnikovik/ndbx/autograder/internal/step/user/one/events/expect"
	userfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
	"github.com/sitnikovik/ndbx/autograder/internal/user/session"
)

// main is the entry point for the Lab 6 autograder.
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
	reviewsTTL := cfg.App().Event().Reviews().TTL()
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
			event.WithReviews(
				review.NewReviews(
					review.WithCounts(
						count.NewCounts(
							count.WithRating(4),
							count.WithCount(2),
						),
					),
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
				timex.MustRFC3339("2026-01-01T11:34:00Z"),
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
			event.WithReviews(
				review.NewReviews(
					review.WithCounts(
						count.NewCounts(
							count.WithRating(4),
							count.WithCount(1),
						),
					),
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
				timex.MustRFC3339("2026-01-01T11:35:00Z"),
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
			event.WithReviews(
				review.NewReviews(
					review.WithCounts(
						count.NewCounts(
							count.WithRating(4),
							count.WithCount(1),
						),
					),
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
			reviewEventEndpoint.NewStep(
				step.NewDesc(
					"Review event bad request",
					"Makes wrong request for event review without comment by the endpoint",
				),
				httpcli,
				baseURL,
				wonderLandEvents[0],
				body.NewBody(
					body.WithRating(rating.Five),
				),
				expectation.NewExpectations(
					expectation.WithAsserts(
						response.AssertBadRequestStatus,
						response.AssertNotEmptyContent,
					),
					expectation.WithCookies(
						cookiexpct.NewExpectations(
							cookieSession.Name,
							cookiexpct.WithAsserts(
								cookieassert.AssertExistsMaxAge,
								cookieassert.AssertExistsHTTPOnly,
							),
							cookiexpct.WithAssertsValueFn(
								session.Validate,
							),
						),
					),
				),
			),
			reviewEventEndpoint.NewStep(
				step.NewDesc(
					"Review event bad request",
					"Makes wrong request for event review without rating by the endpoint",
				),
				httpcli,
				baseURL,
				wonderLandEvents[0],
				body.NewBody(
					body.WithComment("Отлично!"),
				),
				expectation.NewExpectations(
					expectation.WithAsserts(
						response.AssertBadRequestStatus,
						response.AssertNotEmptyContent,
					),
					expectation.WithCookies(
						cookiexpct.NewExpectations(
							cookieSession.Name,
							cookiexpct.WithAsserts(
								cookieassert.AssertExistsMaxAge,
								cookieassert.AssertExistsHTTPOnly,
							),
							cookiexpct.WithAssertsValueFn(
								session.Validate,
							),
						),
					),
				),
			),
			reviewEventEndpoint.NewStep(
				step.NewDesc(
					"Review event",
					"Makes a review for an event created earlier by the endpoint and validates the response",
				),
				httpcli,
				baseURL,
				wonderLandEvents[0],
				body.NewBody(
					body.WithRating(rating.Five),
					body.WithComment("Рекомендую"),
				),
				expectation.NewExpectations(
					expectation.WithAsserts(
						response.AssertCreatedStatus,
						response.AssertNotEmptyContent,
					),
					expectation.WithCookies(
						cookiexpct.NewExpectations(
							cookieSession.Name,
							cookiexpct.WithAsserts(
								cookieassert.AssertExistsMaxAge,
								cookieassert.AssertExistsHTTPOnly,
							),
							cookiexpct.WithAssertsValueFn(
								session.Validate,
							),
						),
					),
				),
			),
			reviewEventEndpoint.NewStep(
				step.NewDesc(
					"Review event conflict",
					"Repeats the request with other comment and rating for event review that has been already left",
				),
				httpcli,
				baseURL,
				wonderLandEvents[0],
				body.NewBody(
					body.WithRating(rating.Four),
					body.WithComment("Я передумал"),
				),
				expectation.NewExpectations(
					expectation.WithAsserts(
						response.AssertConflictStatus,
						response.AssertNotEmptyContent,
					),
					expectation.WithCookies(
						cookiexpct.NewExpectations(
							cookieSession.Name,
							cookiexpct.WithAsserts(
								cookieassert.AssertExistsMaxAge,
								cookieassert.AssertExistsHTTPOnly,
							),
							cookiexpct.WithAssertsValueFn(
								session.Validate,
							),
						),
					),
				),
			),
			reviewEventEndpoint.NewStep(
				step.NewDesc(
					"Review event",
					"Makes a review for an event created earlier by the endpoint and validates the response",
				),
				httpcli,
				baseURL,
				wonderLandEvents[1],
				body.NewBody(
					body.WithRating(rating.Two),
					body.WithComment("Не рекомендую. Очень плохой спектакль."),
				),
				expectation.NewExpectations(
					expectation.WithAsserts(
						response.AssertCreatedStatus,
						response.AssertNotEmptyContent,
					),
					expectation.WithCookies(
						cookiexpct.NewExpectations(
							cookieSession.Name,
							cookiexpct.WithAsserts(
								cookieassert.AssertExistsMaxAge,
								cookieassert.AssertExistsHTTPOnly,
							),
							cookiexpct.WithAssertsValueFn(
								session.Validate,
							),
						),
					),
				),
			),
			reviewEventEndpoint.NewStep(
				step.NewDesc(
					"Review event",
					"Makes a review for an event created earlier by the endpoint and validates the response",
				),
				httpcli,
				baseURL,
				wonderLandEvents[2],
				body.NewBody(
					body.WithRating(rating.Four),
					body.WithComment("Потрясающие актеры! Но минус балл за плохую организацию"),
				),
				expectation.NewExpectations(
					expectation.WithAsserts(
						response.AssertCreatedStatus,
						response.AssertNotEmptyContent,
					),
					expectation.WithCookies(
						cookiexpct.NewExpectations(
							cookieSession.Name,
							cookiexpct.WithAsserts(
								cookieassert.AssertExistsMaxAge,
								cookieassert.AssertExistsHTTPOnly,
							),
							cookiexpct.WithAssertsValueFn(
								session.Validate,
							),
						),
					),
				),
			),
			updateReviewEndpoint.NewStep(
				step.NewDesc(
					"Update review event",
					"Updates the event review by the endpoin",
				),
				httpcli,
				baseURL,
				wonderLandEvents[0],
				body.NewBody(
					body.WithComment("Рекомендую! Идите скорее пока билеты не раскупили"),
				),
				expectation.NewExpectations(
					expectation.WithAsserts(
						response.AssertNoContentStatus,
						response.AssertEmptyContent,
					),
					expectation.WithCookies(
						cookiexpct.NewExpectations(
							cookieSession.Name,
							cookiexpct.WithAsserts(
								cookieassert.AssertExistsMaxAge,
								cookieassert.AssertExistsHTTPOnly,
							),
							cookiexpct.WithAssertsValueFn(
								session.Validate,
							),
						),
					),
				),
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
			reviewEventEndpoint.NewStep(
				step.NewDesc(
					"Review event",
					"Makes a review for an event created earlier by the endpoint and validates the response",
				),
				httpcli,
				baseURL,
				wonderLandEvents[1],
				body.NewBody(
					body.WithRating(rating.One),
					body.WithComment("Отвратительный спектакль!"),
				),
				expectation.NewExpectations(
					expectation.WithAsserts(
						response.AssertCreatedStatus,
						response.AssertNotEmptyContent,
					),
					expectation.WithCookies(
						cookiexpct.NewExpectations(
							cookieSession.Name,
							cookiexpct.WithAsserts(
								cookieassert.AssertExistsMaxAge,
								cookieassert.AssertExistsHTTPOnly,
							),
							cookiexpct.WithAssertsValueFn(
								session.Validate,
							),
						),
					),
				),
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
			reviewEventEndpoint.NewStep(
				step.NewDesc(
					"Review event",
					"Makes a review for an event created earlier by the endpoint and validates the response",
				),
				httpcli,
				baseURL,
				wonderLandEvents[0],
				body.NewBody(
					body.WithRating(rating.Three),
					body.WithComment("На один раз и то вод вопросом"),
				),
				expectation.NewExpectations(
					expectation.WithAsserts(
						response.AssertCreatedStatus,
						response.AssertNotEmptyContent,
					),
					expectation.WithCookies(
						cookiexpct.NewExpectations(
							cookieSession.Name,
							cookiexpct.WithAsserts(
								cookieassert.AssertExistsMaxAge,
								cookieassert.AssertExistsHTTPOnly,
							),
							cookiexpct.WithAssertsValueFn(
								session.Validate,
							),
						),
					),
				),
			),
			reviewEventEndpoint.NewStep(
				step.NewDesc(
					"Review event",
					"Makes a review for an event created earlier by the endpoint and validates the response",
				),
				httpcli,
				baseURL,
				wonderLandEvents[2],
				body.NewBody(
					body.WithRating(rating.Five),
					body.WithComment("Рекомендую. Идите и не думайте"),
				),
				expectation.NewExpectations(
					expectation.WithAsserts(
						response.AssertCreatedStatus,
						response.AssertNotEmptyContent,
					),
					expectation.WithCookies(
						cookiexpct.NewExpectations(
							cookieSession.Name,
							cookiexpct.WithAsserts(
								cookieassert.AssertExistsMaxAge,
								cookieassert.AssertExistsHTTPOnly,
							),
							cookiexpct.WithAssertsValueFn(
								session.Validate,
							),
						),
					),
				),
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
			reviewEventEndpoint.NewStep(
				step.NewDesc(
					"Review event",
					"Makes a review for an event created earlier by the endpoint and validates the response",
				),
				httpcli,
				baseURL,
				wonderLandEvents[0],
				body.NewBody(
					body.WithRating(rating.Five),
					body.WithComment("Отлично!"),
				),
				expectation.NewExpectations(
					expectation.WithAsserts(
						response.AssertCreatedStatus,
						response.AssertNotEmptyContent,
					),
					expectation.WithCookies(
						cookiexpct.NewExpectations(
							cookieSession.Name,
							cookiexpct.WithAsserts(
								cookieassert.AssertExistsMaxAge,
								cookieassert.AssertExistsHTTPOnly,
							),
							cookiexpct.WithAssertsValueFn(
								session.Validate,
							),
						),
					),
				),
			),
			reviewRedis.NewStep(
				step.NewDesc(
					"Event reviews in Redis",
					"Checks how the event reviews stored in Redis",
				),
				rediscli,
				wonderLandEvents[0],
				reviewRedis.WithExpectations(
					reviewRedisExpects.NewExpectations(
						reviewRedisExpects.WithCounts(
							count.NewCounts(
								count.WithRating(3.6),
								count.WithCount(7),
							),
						),
						reviewRedisExpects.WithTTL(
							reviewsTTL,
						),
					),
				),
			),
			listUserEventsEndpoint.NewStep(
				httpcli,
				baseURL,
				samSepiol,
				eventsrq.NewBody(
					eventsrq.WithInclude(
						include.NewInclude("reviews"),
					),
				),
				[]event.Event{
					wonderLandEvents[0],
					wonderLandEvents[1],
				},
				listUserEventsEndpoint.WithExpectations(
					listUserEventsExpectations.NewExpectations(
						listUserEventsExpectations.WithReviews(
							review.NewReviews(
								review.WithCounts(
									count.NewCounts(
										count.WithRating(3.6),
										count.WithCount(7),
									),
								),
							),
							review.NewReviews(
								review.WithCounts(
									count.NewCounts(
										count.WithRating(3.6),
										count.WithCount(7),
									),
								),
							),
						),
					),
				),
			),
			listEventsEndpoint.NewStep(
				step.NewDesc(
					"List events including reviews",
					"Gets lists of all events with reviews",
				),
				httpcli,
				baseURL,
				eventsrq.NewBody(
					eventsrq.WithInclude(
						include.NewInclude("reviews"),
					),
				),
				listEventsExpectations.NewExpectations(
					listEventsExpectations.WithEvents(
						wonderLandEvents[0],
						wonderLandEvents[1],
						wonderLandEvents[2],
					),
					listEventsExpectations.WithReviews(
						review.NewReviews(
							review.WithCounts(
								count.NewCounts(
									count.WithRating(3.6),
									count.WithCount(7),
								),
							),
						),
						review.NewReviews(
							review.WithCounts(
								count.NewCounts(
									count.WithRating(3.6),
									count.WithCount(7),
								),
							),
						),
						review.NewReviews(
							review.WithCounts(
								count.NewCounts(
									count.WithRating(3.6),
									count.WithCount(7),
								),
							),
						),
					),
				),
			),
			logout.NewStep(
				httpcli,
				baseURL,
			),
			reviewEventEndpoint.NewStep(
				step.NewDesc(
					"Event reviews in Redis",
					"Checks how the event reviews stored in Redis",
				),
				httpcli,
				baseURL,
				wonderLandEvents[0],
				body.NewBody(
					body.WithRating(rating.Five),
					body.WithComment("Отлично!"),
				),
				expectation.NewExpectations(
					expectation.WithAsserts(
						response.AssertUnauthorizedStatus,
						response.AssertEmptyContent,
					),
				),
			),
			reviewRedis.NewStep(
				step.NewDesc(
					"Event reviews in Redis",
					"Checks how the event reviews stored in Redis",
				),
				rediscli,
				wonderLandEvents[1],
				reviewRedis.WithExpectations(
					reviewRedisExpects.NewExpectations(
						reviewRedisExpects.WithCounts(
							count.NewCounts(
								count.WithRating(3.6),
								count.WithCount(7),
							),
						),
						reviewRedisExpects.WithTTL(
							reviewsTTL,
						),
					),
				),
			),
			reviewRedis.NewStep(
				step.NewDesc(
					"Event reviews in Redis",
					"Checks how the event reviews stored in Redis",
				),
				rediscli,
				wonderLandEvents[2],
				reviewRedis.WithExpectations(
					reviewRedisExpects.NewExpectations(
						reviewRedisExpects.WithCounts(
							count.NewCounts(
								count.WithRating(3.6),
								count.WithCount(7),
							),
						),
						reviewRedisExpects.WithTTL(
							reviewsTTL,
						),
					),
				),
			),
			eventReviewsCassandra.NewStep(
				step.NewDesc(
					"Get event reviews from Cassandra",
					"Selects event reviews from Cassandra databases and validates the rows",
				),
				cassandracli,
				wonderLandEvents[0],
				eventReviewsCassandraXpct.NewExpectations(
					eventReviewsCassandraXpct.WithCount(3),
				),
			),
			eventReviewsCassandra.NewStep(
				step.NewDesc(
					"Get event reviews from Cassandra",
					"Selects event reviews from Cassandra databases and validates the rows",
				),
				cassandracli,
				wonderLandEvents[1],
				eventReviewsCassandraXpct.NewExpectations(
					eventReviewsCassandraXpct.WithCount(2),
				),
			),
			eventReviewsCassandra.NewStep(
				step.NewDesc(
					"Get event reviews from Cassandra",
					"Selects event reviews from Cassandra databases and validates the rows",
				),
				cassandracli,
				wonderLandEvents[2],
				eventReviewsCassandraXpct.NewExpectations(
					eventReviewsCassandraXpct.WithCount(2),
				),
			),
			reviewListEndpoint.NewStep(
				step.NewDesc(
					"Get event reviews by endpoint",
					"Gets reviews for the event by the endpoint and validates the response",
				),
				httpcli,
				wonderLandEvents[0],
				baseURL,
				reviewListEndpointXpct.NewExpectations(
					reviewListEndpointXpct.WithCount(3),
				),
			),
			reviewListEndpoint.NewStep(
				step.NewDesc(
					"Get event reviews by endpoint",
					"Gets reviews for the event by the endpoint and validates the response",
				),
				httpcli,
				wonderLandEvents[1],
				baseURL,
				reviewListEndpointXpct.NewExpectations(
					reviewListEndpointXpct.WithCount(2),
				),
			),
			reviewListEndpoint.NewStep(
				step.NewDesc(
					"Get event reviews by endpoint",
					"Gets reviews for the event by the endpoint and validates the response",
				),
				httpcli,
				wonderLandEvents[2],
				baseURL,
				reviewListEndpointXpct.NewExpectations(
					reviewListEndpointXpct.WithCount(2),
				),
			),
			listUserEventsEndpoint.NewStep(
				httpcli,
				baseURL,
				samSepiol,
				eventsrq.NewBody(
					eventsrq.WithInclude(
						include.NewInclude("reviews"),
					),
				),
				[]event.Event{
					wonderLandEvents[0],
					wonderLandEvents[1],
				},
				listUserEventsEndpoint.WithExpectations(
					listUserEventsExpectations.NewExpectations(
						listUserEventsExpectations.WithReviews(
							review.NewReviews(
								review.WithCounts(
									count.NewCounts(
										count.WithRating(3.6),
										count.WithCount(7),
									),
								),
							),
							review.NewReviews(
								review.WithCounts(
									count.NewCounts(
										count.WithRating(3.6),
										count.WithCount(7),
									),
								),
							),
						),
					),
				),
			),
			listEventsEndpoint.NewStep(
				step.NewDesc(
					"List events after reviewed",
					"Gets lists of all events with reviews after someone reviewed",
				),
				httpcli,
				baseURL,
				eventsrq.NewBody(
					eventsrq.WithInclude(
						include.NewInclude("reviews"),
					),
				),
				listEventsExpectations.NewExpectations(
					listEventsExpectations.WithEvents(
						wonderLandEvents[0],
						wonderLandEvents[1],
						wonderLandEvents[2],
					),
					listEventsExpectations.WithReviews(
						review.NewReviews(
							review.WithCounts(
								count.NewCounts(
									count.WithRating(3.6),
									count.WithCount(7),
								),
							),
						),
						review.NewReviews(
							review.WithCounts(
								count.NewCounts(
									count.WithRating(3.6),
									count.WithCount(7),
								),
							),
						),
						review.NewReviews(
							review.WithCounts(
								count.NewCounts(
									count.WithRating(3.6),
									count.WithCount(7),
								),
							),
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
		console.Fatal("Lab 6 autograder failed: %v", err)
	}
}
