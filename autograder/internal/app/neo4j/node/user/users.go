package user

import "github.com/sitnikovik/ndbx/autograder/internal/app/user"

// Users represents a collection of users
// stored in Neo4j database of the target application.
type Users []User

// NewUsers creates a new collection of users.
func NewUsers(uu ...User) Users {
	return uu
}

// OneWithID returns the user with the given identifier.
//
// If there is no such user, it returns an empty user.
func (uu Users) OneWithID(id user.ID) User {
	for _, u := range uu {
		if u.ID() == id {
			return u
		}
	}
	return User{}
}
