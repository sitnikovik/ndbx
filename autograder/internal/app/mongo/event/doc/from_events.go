package doc

import "github.com/sitnikovik/ndbx/autograder/internal/app/event"

// FromEvents converts the list of Event to the list of EventDocument related to MongoDB.
func FromEvents(ee []event.Event) EventDocuments {
	if len(ee) == 0 {
		return nil
	}
	res := make([]EventDocument, len(ee))
	for i, e := range ee {
		res[i] = FromEvent(e)
	}
	return res
}
