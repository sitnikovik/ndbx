package endpoint_test

import (
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	eventfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/event"
	userfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

var (
	// descFixture is a step description fixture used in the tests cases.
	descFixture = step.NewDesc(
		"Test step",
		"Test description",
	)
	// eventFixture is an event fixture used in the tests cases.
	eventFixture = eventfx.NewBirthdayParty(
		event.NewDates(
			timex.MustRFC3339("2026-03-31T15:00:00Z"),
			timex.MustRFC3339("2026-03-31T23:00:00Z"),
		),
		timex.MustRFC3339("2026-03-14T12:31:00Z"),
		userfx.NewSamwiseGamgee(),
	)
	// varsFixture is the step variables fixture used in the tests cases.
	varsFixture = func() step.Variables {
		vv := step.NewVariables()
		vv.Set(
			eventFixture.Hash(),
			"13298",
		)
		return vv
	}()
)
