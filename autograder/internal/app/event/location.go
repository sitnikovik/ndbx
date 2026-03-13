package event

// Location represents the physical or virtual location of an event.
type Location struct {
	// addr is the address of the event location.
	addr string
}

// NewLocation creates a new Location instance.
func NewLocation(addr string) Location {
	return Location{
		addr: addr,
	}
}

// Address returns the address of the event location.
func (l Location) Address() string {
	return l.addr
}
