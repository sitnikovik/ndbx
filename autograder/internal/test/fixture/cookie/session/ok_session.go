package session

import (
	"time"

	sidfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/cookie/session/id"
)

// NewOKSession returns a new valid fixture of cookie session.
func NewOKSession() string {
	return NewSession(
		sidfx.OK,
		3600*time.Second,
	)
}
