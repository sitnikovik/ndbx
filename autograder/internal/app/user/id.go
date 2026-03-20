package user

// ID represents a user ID, which is a unique identifier for a user.
type ID string

// NewID creates a new ID instance from the given string.
func NewID(id string) ID {
	return ID(id)
}

// String returns the string representation of the user ID.
func (i ID) String() string {
	return string(i)
}

// Empty defines if id is empty.
func (i ID) Empty() bool {
	return i == ""
}
