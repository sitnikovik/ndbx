package variable

import "github.com/sitnikovik/ndbx/autograder/internal/app/event"

// MustEvenD retrieves the Event variable from the original variables
// and panics if it is not an instance of Event.
func (v Values) MustEvent() event.Event {
	res, ok := (v.orig.MustGet(Event).Value()).(event.Event)
	if !ok {
		panic("expecting the Event type in variables but it's not")
	}
	return res
}
