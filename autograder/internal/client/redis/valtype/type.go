package valtype

import "slices"

// Type defines the Type type,
// which represents the type of a value stored in Redis.
type Type string

const (
	// String represents a Redis string value.
	String Type = "string"
	// Hash represents a Redis hash value.
	Hash Type = "hash"
	// None represents the absence of a value or an unknown type.
	None Type = "none"
)

// ParseType converts a string representation of a Redis value type
// into a Type value.
//
// It returns the corresponding Type for known Redis types,
// or None for unknown types.
func ParseType(t string) Type {
	switch t {
	case "string":
		return String
	case "hash":
		return Hash
	default:
		return None
	}
}

// IsNone checks if the Type is None.
func (t Type) IsNone() bool {
	return t == None
}

// IsString checks if the Type is String.
func (t Type) IsString() bool {
	return t == String
}

// IsHash checks if the Type is Hash.
func (t Type) IsHash() bool {
	return t == Hash
}

// In checks if the Type is in the provided list of Types.
func (t Type) In(types ...Type) bool {
	return slices.Contains(types, t)
}

// String returns the string representation of the Type.
func (t Type) String() string {
	return string(t)
}
