package session

import (
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/http/response/cookie"
)

// MustParseSession creates a new Session instance from a slice of HTTP cookies.
//
// It panics if the required session cookie does not exist.
func MustParseSession(ckk []*http.Cookie) Session {
	return NewSession(
		cookie.
			NewCookies(ckk).
			MustGet(Name),
	)
}
