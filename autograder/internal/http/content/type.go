package content

// Type represents the type of content being sent in an HTTP request.
type Type string

const (
	// ApplicationJSON is the content type for JSON data.
	ApplicationJSON Type = "application/json"
)

// String returns the string representation of the content type.
func (t Type) String() string {
	return string(t)
}
