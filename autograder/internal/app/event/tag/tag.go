package tag

// Tag represents an event tag displays the area covered by the event
// or interests that related with the event.
type Tag string

// String returns string representation of the tag.
func (t Tag) String() string {
	return string(t)
}
