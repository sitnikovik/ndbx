package event

// Location represents the physical or virtual location of an event.
type Location struct {
	// addr is the address of the event location.
	addr string
	// city is the city where the event is happening.
	city string
}

// LocationOption represents a functional option for configuring a Location of the event.
type LocationOption func(l *Location)

// WithCity sets the provided city to the Location.
func WithCity(city string) LocationOption {
	return func(l *Location) {
		l.city = city
	}
}

// NewLocation creates a new Location instance.
func NewLocation(addr string, opts ...LocationOption) Location {
	l := Location{
		addr: addr,
	}
	for _, opt := range opts {
		opt(&l)
	}
	return l
}

// Address returns the address of the event location.
func (l Location) Address() string {
	return l.addr
}

// City return the city where the event is happening.
func (l Location) City() string {
	return l.city
}

// Empty defines if the location is empty.
func (l Location) Empty() bool {
	return l.Address() == "" && l.City() == ""
}
