package user

import (
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	app "github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

// NewAlexSmith creates a fixture of the User used in tests.
func NewAlexSmith() user.User {
	return app.NewUser(
		app.NewID(
			"3",
		),
		"al3xsm1th256",
		"Alex Smith Jr.",
	)
}
