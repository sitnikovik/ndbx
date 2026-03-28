package event

import "github.com/sitnikovik/ndbx/autograder/internal/app/money"

// Costs represents the cost information related to the event.
type Costs struct {
	// entry is the entry price the attendee have to pay.
	entry money.Money
}

// NewCosts creates a new Costs instance.
func NewCosts(entry money.Money) Costs {
	return Costs{
		entry: entry,
	}
}

// Entry returns the entry price the attendee have to pay.
func (c Costs) Entry() money.Money {
	return c.entry
}
