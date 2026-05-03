package cookie

import (
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

// AssertValueFn asserts that cookie value satisfies given function.
func AssertValueFn(
	ckk []*http.Cookie,
	name string,
	fn func(v string) error,
) error {
	for _, ck := range ckk {
		if ck.Name == name {
			return errs.Wrap(fn(ck.Value), name)
		}
	}
	return errs.Wrap(errs.ErrMissedCookie, name)
}
