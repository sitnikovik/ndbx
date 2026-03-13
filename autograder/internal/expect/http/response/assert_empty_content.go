package response

import (
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

// AssertEmptyContent checks if the HTTP response has an empty content (Content-Length of 0)
// and returns an error if it does not.
func AssertEmptyContent(rsp *http.Response) error {
	if rsp.ContentLength != 0 {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"expected empty content, got content length %d",
			rsp.ContentLength,
		)
	}
	return nil
}
