package main

import (
	"context"
	"os"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event/category"
	"github.com/sitnikovik/ndbx/autograder/internal/app/money"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
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
	createOneEventEndpoint "github.com/sitnikovik/ndbx/autograder/internal/step/events/create/one/endpoint/ok"
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
			createOneEventEndpoint.NewStep(
				httpcli,
				baseURL,
				event.NewEvent(
					event.NewID(""),
					event.NewContent(
						"В стране чудес",
						"Выставка картин и иллюстраций Екатерины Ващинской.\n\n"+
							"10 марта в галерее YOLKA art открылась выставка Екатерины Ващинской «В стране Чудес».\n\n"+
							"Художница широко известна, как иллюстратор, однако ее живопись выходит за пределы книжной графики."+
							"В отличие от большинства иллюстраторов она работает маслом"+
							"и строит изображение через сложную фактуру,"+
							" что придает работам ощущение материальности и одновременно сказочной хрупкости."+
							"\n\n"+
							"Выставка создана по мотивам иллюстраций Екатерины Ващинской"+
							" к новому изданию «Алисы в стране Чудес»,"+
							" вышедшему в 2025 году. "+
							"Здесь привычные персонажи оказываются внутри мира,"+
							"где пространство напоминает театральную сцену,"+
							" шахматную доску или фантастический сад. "+
							"Художница соединяет элементы декоративной живописи,"+
							" книжной иллюстрации и фантастического реализма."+
							"\n\n"+
							"Созданный Ващинской мир сохраняет дух книги Кэрролла и "+
							"одновременно продолжает длинную художественную традицию интерпретаций Страны Чудес."+
							" Здесь Алиса вновь оказывается проводником в пространство, "+
							"где границы между реальностью и фантазией становятся условными.",
						event.WithCategory(category.Exhibition),
					),
					event.NewLocation(
						"Москва. Ходынский бульвар 20а",
						event.WithCity("Москва"),
					),
					event.NewCreated(
						timex.MustRFC3339("2026-01-01T11:33:00Z"),
						user.NewIdentity(user.ID("123")),
					),
					event.NewDates(
						timex.MustRFC3339("2026-03-24T10:00:00Z"),
						timex.MustRFC3339("2025-01-24T18:00:00Z"),
					),
					event.WithCosts(
						event.NewCosts(
							money.NewMoney(0, 00),
						),
					),
				),
			),
			createOneEventEndpoint.NewStep(
				httpcli,
				baseURL,
				event.NewEvent(
					event.NewID(""),
					event.NewContent(
						"В стране чудес",
						"Выставка картин и иллюстраций Екатерины Ващинской.\n\n"+
							"10 марта в галерее YOLKA art открылась выставка Екатерины Ващинской «В стране Чудес».\n\n"+
							"Художница широко известна, как иллюстратор, однако ее живопись выходит за пределы книжной графики."+
							"В отличие от большинства иллюстраторов она работает маслом"+
							"и строит изображение через сложную фактуру,"+
							" что придает работам ощущение материальности и одновременно сказочной хрупкости."+
							"\n\n"+
							"Выставка создана по мотивам иллюстраций Екатерины Ващинской"+
							" к новому изданию «Алисы в стране Чудес»,"+
							" вышедшему в 2025 году. "+
							"Здесь привычные персонажи оказываются внутри мира,"+
							"где пространство напоминает театральную сцену,"+
							" шахматную доску или фантастический сад. "+
							"Художница соединяет элементы декоративной живописи,"+
							" книжной иллюстрации и фантастического реализма."+
							"\n\n"+
							"Созданный Ващинской мир сохраняет дух книги Кэрролла и "+
							"одновременно продолжает длинную художественную традицию интерпретаций Страны Чудес."+
							" Здесь Алиса вновь оказывается проводником в пространство, "+
							"где границы между реальностью и фантазией становятся условными.",
						event.WithCategory(category.Exhibition),
					),
					event.NewLocation(
						"Москва. Ходынский бульвар 20а",
						event.WithCity("Москва"),
					),
					event.NewCreated(
						timex.MustRFC3339("2026-01-01T11:33:00Z"),
						user.NewIdentity(user.ID("123")),
					),
					event.NewDates(
						timex.MustRFC3339("2026-03-25T10:00:00Z"),
						timex.MustRFC3339("2025-01-25T18:00:00Z"),
					),
					event.WithCosts(
						event.NewCosts(
							money.NewMoney(0, 00),
						),
					),
				),
			),
			createOneEventEndpoint.NewStep(
				httpcli,
				baseURL,
				event.NewEvent(
					event.NewID(""),
					event.NewContent(
						"Концерт «Вечер венской классики»",
						"26 марта в 18:00 приглашаем в библиотеку Культурного центра ЗИЛ на концерт «Вечер венской классики»."+
							"Исполнители:\n"+
							"Алексей Неберикутин(скрипка)\n"+
							"Елена Фришман (скрипка)\n"+
							"Александра Тельманова(альт)\n"+
							"Валерий Мастеров(виолончель)"+
							"\n\n"+
							"В программе:\n"+
							"Йозеф Гайдн\n"+
							"Струнный квартет ми мажор\n"+
							"op.17 №1, Hob III:25"+
							"\n\n"+
							"Вольфганг Амадей Моцарт\n"+
							"Струнный квартет №19 до мажор\n"+
							"«Диссонанс»\n"+
							"KV 465",
						event.WithCategory(category.Concert),
					),
					event.NewLocation(
						"Культурный центр ЗИЛ",
						event.WithCity("Москва"),
					),
					event.NewCreated(
						timex.MustRFC3339("2025-10-10T18:25:00Z"),
						user.NewIdentity(user.ID("123")),
					),
					event.NewDates(
						timex.MustRFC3339("2026-03-25T18:00:00Z"),
						timex.MustRFC3339("2025-01-25T21:00:00Z"),
					),
					event.WithCosts(
						event.NewCosts(
							money.NewMoney(2000, 00),
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
		).
		Run(ctx, vars)
	if err != nil {
		console.Fatal("Lab 4 autograder failed: %v", err)
	}
}
