package creation

// Stamp represents a creation stamp of some object in the app or database.
type Stamp struct {
	// created holds creation details.
	created Created
}

// NewStamp create a new Stamp instance
// represents a creation stamp of some object in the app or database.
func NewStamp(created Created) Stamp {
	return Stamp{
		created: created,
	}
}

// Created returns the creation details of the object.
func (s Stamp) Created() Created {
	return s.created
}
