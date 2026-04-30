package cookie

import (
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

// AssertExistsHTTPOnly asserts that cookie with given name exists
// in the list of cookies and has HttpOnly flag set and returns an error otherwise.
func AssertExistsHTTPOnly(ckk []*http.Cookie, name string) error {
	for _, ck := range ckk {
		if ck.Name == name {
			if ck.HttpOnly {
				return nil
			}
			return errs.Wrap(
				errs.ErrInvalidValue,
				"%s: HttpOnly",
				name,
			)
		}
	}
	return errs.Wrap(errs.ErrMissedCookie, name)
}
