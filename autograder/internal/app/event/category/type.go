package category

// Type represents the type of the event.
type Type uint8

const (
	// Unspecified represents an unspecified event type.
	Unspecified Type = 0
	// Other represents an other event.
	Other Type = 1
	// Meetup represents a meetup event.
	Meetup Type = 2
	// Concert represents a concert event.
	Concert Type = 3
	// Exhibition represents an exhibition event.
	Exhibition Type = 4
	// Party represents a party event.
	Party Type = 5
)

// Unspecified defines if the event type is unspecified.
func (t Type) Unspecified() bool {
	return t == Unspecified
}

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
	case Other.String():
		return Other
	default:
		return Unspecified
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
	case Other:
		return "other"
	default:
		return ""
	}
}
