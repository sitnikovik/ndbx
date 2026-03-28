package user

import app "github.com/sitnikovik/ndbx/autograder/internal/app/user"

// NewJohnDoe creates a fixture of the User used in tests.
func NewJohnDoe() app.User {
	return app.NewUser(
		app.NewID(
			"2",
		),
		"j0hnd0e42",
		"John Doe",
	)
}
