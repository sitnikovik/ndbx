package user

import (
	app "github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

// NewSamSepiol creates a fixture of the User used in tests.
func NewSamSepiol() app.User {
	return app.NewUser(
		app.NewID(
			"1",
		),
		"samsep1ol",
		"Sam Sepiol",
	)
}
