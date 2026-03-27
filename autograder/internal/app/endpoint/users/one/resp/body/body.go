package body

import "github.com/sitnikovik/ndbx/autograder/internal/app/user"

// Body represents the HTTP response body of the user endpoint.
type Body struct {
	// user is the user got by the endpoint.
	user user.User
}

// NewBody creates a new Body instance.
func NewBody(usr user.User) Body {
	return Body{
		user: usr,
	}
}

// User returns the User got by the endpoint.
func (b Body) User() user.User {
	return b.user
}
