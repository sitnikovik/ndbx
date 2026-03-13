package lab3

import "github.com/sitnikovik/ndbx/autograder/internal/app/user"

// NewTestUser creates and returns a new user to be used in the autograder.
func NewTestUser() user.User {
	return user.NewUser(
		user.NewID(""),
		"sams3p1ol",
		"Sam Sepiol",
	)
}
