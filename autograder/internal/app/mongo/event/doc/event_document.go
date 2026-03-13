package doc

import "github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"

// EventDocument represents an event document in MongoDB.
type EventDocument struct {
	// orig is the original MongoDB document that contains the event data.
	orig doc.Document
}

// NewEventDocument creates a new EventDocument struct from a MongoDB document.
func NewEventDocument(orig doc.Document) EventDocument {
	return EventDocument{
		orig: orig,
	}
}
