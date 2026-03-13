package response

import (
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

// AssertNotEmptyContent checks if the HTTP response has a non-empty content
// and returns an error if it does not.
func AssertNotEmptyContent(rsp *http.Response) error {
	if rsp.ContentLength == 0 {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"expected non-empty content, but got the empty one",
		)
	}
	return nil
}
