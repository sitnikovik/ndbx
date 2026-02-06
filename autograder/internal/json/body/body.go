package body

import (
	"io"
)

// Body represents a JSON body that can be parsed into a Go struct.
type Body struct {
	// reader is the underlying reader for the JSON body.
	reader io.Reader
}

// NewBody creates a new Body instance.
//
// Parameters:
//   - reader: The io.Reader that provides the JSON data.
func NewBody(reader io.Reader) Body {
	return Body{
		reader: reader,
	}
}
