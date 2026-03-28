package body

import (
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

// Body represents the HTTP response body of the list of users endpoint.
type Body struct {
	// users is the list of the users got from response.
	users []user.User
	// n is total count of the list.
	n int
}

// NewBody creates a new Body instance.
func NewBody(users []user.User, n int) Body {
	return Body{
		users: users,
		n:     n,
	}
}

// Users returns a list of the users got from response.
func (b Body) Users() []user.User {
	return b.users
}

// Count returns total count of the list.
func (b Body) Count() int {
	return b.n
}
