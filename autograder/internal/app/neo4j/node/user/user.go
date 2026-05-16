package user

import "github.com/sitnikovik/ndbx/autograder/internal/app/user"

// User represents a user in the Neo4j database
// of the target application.
type User struct {
	// id is the unique identifier of the user.
	id user.ID
}

// NewUser creates a new User
// stored in Neo4j with the given user ID.
func NewUser(id user.ID) User {
	return User{
		id: id,
	}
}

// ID returns the user's identifier.
func (u User) ID() user.ID {
	return u.id
}

// Equals checks if the given user is equal to the current one.
func (u User) Equals(other User) bool {
	return u.id == other.id
}
