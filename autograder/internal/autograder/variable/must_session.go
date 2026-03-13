package variable

import "github.com/sitnikovik/ndbx/autograder/internal/app/user/session"

// MustSession retrieves the session variable from the original variables
// and returns it as a Session instance.
//
// Panics if the session variable is not set or has an invalid type.
func (v Values) MustSession() session.Session {
	sess, ok := v.orig.MustGet(Session).Value().(session.Session)
	if !ok {
		panic("session variable is not set or has invalid type")
	}
	return sess
}
