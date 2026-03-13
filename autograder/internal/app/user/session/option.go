package session

// Option represents a functional option for configuring a Session.
type Option func(*Session)

// WithUser returns an Option that sets the user of a session.
func WithUser(usr User) Option {
	return func(s *Session) {
		s.user = usr
	}
}
