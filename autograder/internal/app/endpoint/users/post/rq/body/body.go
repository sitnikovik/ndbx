package body

import (
	"encoding/json"
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

// Body represents the request body for the sign-up endpoint.
type Body struct {
	// usr holds all user's data.
	usr user.User
	// pwd is the password of the user.
	pwd string
}

// NewBody creates a new Body instance with the provided user data.
func NewBody(usr user.User, pwd string) Body {
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
		"full_name": b.usr.FullName(),
		"username":  b.usr.Username(),
		"password":  b.pwd,
	}
	bb, err := json.Marshal(m)
	if err != nil {
		panic(errors.Join(errs.ErrMarshallFailed, err))
	}
	return bb
}
