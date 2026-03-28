package response

import (
	"errors"
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
)

// AssertNotFoundStatus checks if the HTTP response
// has a status code of 404 Not Found and returns an error if it does not.
func AssertNotFoundStatus(rsp *http.Response) error {
	err := numbers.AssertEquals(
		http.StatusNotFound,
		rsp.StatusCode,
	)
	if err != nil {
		return errors.Join(
			errs.ErrInvalidHTTPStatus,
			err,
		)
	}
	return nil
}
