package tag

// Tag represents an event tag displays the area covered by the event
// or interests that related with the event.
type Tag string

// NewTag creates a new instance of Tag.
func NewTag(s string) Tag {
	return Tag(s)
}

// String returns string representation of the tag.
func (t Tag) String() string {
	return string(t)
}
