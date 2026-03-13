package variable

import "github.com/sitnikovik/ndbx/autograder/internal/app/user"

// MustUser retrieves the user variable from the original variables
// and returns it as a User instance.
//
// Panics if the user variable is not set or has an invalid type.
func (v Values) MustUser() user.User {
	usr, ok := v.orig.MustGet(User).Value().(user.User)
	if !ok {
		panic("user variable is not set or has invalid type")
	}
	return usr
}
