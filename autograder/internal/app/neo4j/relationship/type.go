package relationship

// Type represents the type of a relationship.
type Type string

const (
	// Liked represents relationship type for liking an event by a user.
	Liked Type = "LIKED"
)

// NewType creates a new Type by the given string.
func NewType(s string) Type {
	return Type(s)
}

// String returns the string representation of the Type.
func (t Type) String() string {
	return string(t)
}
