package session

import "github.com/sitnikovik/ndbx/autograder/internal/app/user"

// User represents a user base information stored in the session.
type User struct {
	// ID is the unique identifier of the user.
	id user.ID
}

// NewUser creates and returns a new User with the given ID.
func NewUser(id user.ID) User {
	return User{
		id: id,
	}
}

// ID returns the unique identifier of the user.
func (u User) ID() user.ID {
	return u.id
}
