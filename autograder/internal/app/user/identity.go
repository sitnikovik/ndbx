package user

// Identity represents the unique identifier of a user in the autograder application.
type Identity struct {
	// id is the unique identifier of the user.
	id ID
}

// NewIdentity creates and returns a new Identity instance.
func NewIdentity(id ID) Identity {
	return Identity{
		id: id,
	}
}

// ID returns the unique identifier of the user.
func (i Identity) ID() ID {
	return i.id
}
