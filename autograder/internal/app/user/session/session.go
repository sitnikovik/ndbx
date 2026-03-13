package session

// Session represents a user session.
type Session struct {
	// dates holds dates related to the session, such as creation and last update times.
	dates Dates
	// user is the user associated with the session.
	user User
	// id is the unique identifier for the session.
	id ID
}

// NewSession creates a new Session instance.
func NewSession(
	id ID,
	dates Dates,
	opts ...Option,
) Session {
	s := Session{
		id:    id,
		dates: dates,
	}
	for _, opt := range opts {
		opt(&s)
	}
	return s
}

// ID returns the unique identifier for the session.
func (s Session) ID() ID {
	return s.id
}

// Dates returns the dates related to the session, such as creation and last update times.
func (s Session) Dates() Dates {
	return s.dates
}

// User returns the user associated with the session.
func (s Session) User() User {
	return s.user
}

func (s Session) String() string {
	return s.ID().String()
}
