package category

// Type represents the type of the event.
type Type uint8

const (
	// Other represents an other event.
	Other Type = 0
	// Meetup represents a meetup event.
	Meetup Type = 1
	// Concert represents a concert event.
	Concert Type = 2
	// Exhibition represents an exhibition event.
	Exhibition Type = 3
	// Party represents a party event.
	Party Type = 4
)

// Parse parses the event type by a string.
//
// Returns Other if not parsed or unknown.
func Parse(s string) Type {
	switch s {
	case Meetup.String():
		return Meetup
	case Concert.String():
		return Concert
	case Exhibition.String():
		return Exhibition
	case Party.String():
		return Party
	default:
		return Other
	}
}

// String returns a string representation of the event type.
func (t Type) String() string {
	switch t {
	case Meetup:
		return "meetup"
	case Concert:
		return "concert"
	case Exhibition:
		return "exhibition"
	case Party:
		return "party"
	default:
		return "other"
	}
}
