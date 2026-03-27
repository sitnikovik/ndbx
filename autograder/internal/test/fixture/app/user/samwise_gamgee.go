package user

import (
	app "github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

// NewSamwiseGamgee creates a fixture of the User used in tests.
func NewSamwiseGamgee() app.User {
	return app.NewUser(
		app.NewID("5"),
		"samw1seGamgee",
		"Samwise Gamgee",
	)
}
