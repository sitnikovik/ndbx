package event

import (
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event/category"
	"github.com/sitnikovik/ndbx/autograder/internal/app/money"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

// NewBirthdayParty returns new fixtured event to use in tests.
func NewBirthdayParty(
	when event.Dates,
	createdAt time.Time,
	createdBy user.User,
) event.Event {
	return event.NewEvent(
		event.NewID("1"),
		event.NewContent(
			"Мой день рождения",
			"Приглашаю вас отпраздновать мое 30-с-чем-то-летие",
			event.WithCategory(category.Party),
		),
		event.NewLocation(
			"У меня дома",
		),
		event.NewCreated(
			createdAt,
			createdBy.Idendity(),
		),
		when,
		event.WithCosts(
			event.NewCosts(
				money.NewMoney(0, 00),
			),
		),
	)
}
