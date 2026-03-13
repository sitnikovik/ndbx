package session

// ID represents a session ID, which is a unique identifier for a user session.
type ID string

// NewID creates a new ID instance from the given string.
func NewID(id string) ID {
	return ID(id)
}

// String returns the string representation of the session ID.
func (i ID) String() string {
	return string(i)
}
