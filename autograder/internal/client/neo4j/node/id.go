package node

// ID represents the unique identifier of a node (element id).
type ID string

// NewID creates a new ID from the given string.
func NewID(id string) ID {
	return ID(id)
}

// String returns the string representation of the identifier.
func (id ID) String() string {
	return string(id)
}

// Equals defines if the identified is equal to another one.
func (id ID) Equals(other ID) bool {
	return id == other
}

// Empty defines if the identified is empty.
func (id ID) Empty() bool {
	return id == ""
}
