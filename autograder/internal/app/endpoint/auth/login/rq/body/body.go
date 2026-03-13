package body

import (
	"encoding/json"
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

// Body represents the request body for the authentication endpoint.
type Body struct {
	// usr is the username of the user.
	usr string
	// pwd is the password of the user.
	pwd string
}

// NewBody creates a new Body instance with the provided user data.
func NewBody(usr, pwd string) Body {
	return Body{
		usr: usr,
		pwd: pwd,
	}
}

// MustBytes returns the JSON representation of the Body as a byte slice.
//
// Panics if marshalling the data fails.
func (b Body) MustBytes() []byte {
	m := map[string]string{
		"username": b.usr,
		"password": b.pwd,
	}
	bb, err := json.Marshal(m)
	if err != nil {
		panic(errors.Join(errs.ErrMarshallFailed, err))
	}
	return bb
}
