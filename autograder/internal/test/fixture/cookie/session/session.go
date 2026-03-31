package session

import (
	"fmt"
	"time"

	cookie "github.com/sitnikovik/ndbx/autograder/internal/app/cookie/session"
)

// NewSession returns a new fixture of cookie session.
func NewSession(
	sid string,
	ttl time.Duration,
) string {
	return fmt.Sprintf(
		"%s=%s; HttpOnly; Max-Age=%d; Secure=true",
		cookie.Name,
		sid,
		int(ttl.Seconds()),
	)
}
