package user

import "github.com/sitnikovik/ndbx/autograder/internal/app/user"

// NewJohnSmith creates a fixture of the User used in tests.
func NewJohnSmith() user.User {
	return user.NewUser(
		user.NewID("4"),
		"johnsm1th",
		"John Smith",
	)
}
