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

// KVs returns the slice of key-value pairs representing the fields of the document.
func (e EventDocument) KVs() doc.KVs {
	return e.orig.KVs()
}
