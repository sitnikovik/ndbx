package session

import (
	"net/http"
)

// Name is the name of the session cookie.
const Name = "X-Session-Id"

// Session represents a user session, encapsulating the session cookie.
type Session struct {
	// ck is the HTTP cookie that contains the session information.
	ck *http.Cookie
}

// NewSession creates a new Session instance from the provided HTTP cookie.
func NewSession(ck *http.Cookie) Session {
	return Session{ck: ck}
}

// String returns a string representation of the session.
func (s Session) String() string {
	return s.ck.Value
}

// Expired checks if the session cookie has expired based on its MaxAge field.
func (s Session) Expired() bool {
	return s.ck.MaxAge <= 0
}
