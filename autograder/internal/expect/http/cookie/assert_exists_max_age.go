package cookie

import (
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

// AssertExistsMaxAge asserts that cookie with given name exists
// in the list of cookies and has MaxAge > 0.
func AssertExistsMaxAge(ckk []*http.Cookie, name string) error {
	for _, ck := range ckk {
		if ck.Name == name {
			if ck.MaxAge > 0 {
				return nil
			}
			return errs.Wrap(
				errs.ErrInvalidValue,
				"%s: MaxAge",
				name,
			)
		}
	}
	return errs.Wrap(errs.ErrMissedCookie, name)
}
