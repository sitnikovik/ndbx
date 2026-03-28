package event

// ID represents the unique identifier for an event.
type ID string

// NewID creates a new ID instance from a string.
func NewID(id string) ID {
	return ID(id)
}

// String returns the string representation of the ID.
func (id ID) String() string {
	return string(id)
}

// Empty defines if id is empty.
func (id ID) Empty() bool {
	return id == ""
}
