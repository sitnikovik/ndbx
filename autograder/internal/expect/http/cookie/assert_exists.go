package cookie

import (
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

// AssertExists asserts that cookie with given name exists in the list of cookies.
func AssertExists(ckk []*http.Cookie, name string) error {
	for _, ck := range ckk {
		if ck.Name == name {
			return nil
		}
	}
	return errs.Wrap(errs.ErrMissedCookie, name)
}
