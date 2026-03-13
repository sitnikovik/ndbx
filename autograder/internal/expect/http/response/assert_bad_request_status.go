package response

import (
	"errors"
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
)

// AssertBadRequestStatus checks if the HTTP response has a status code of 400 Bad Request
// and returns an error if it does not.
func AssertBadRequestStatus(rsp *http.Response) error {
	err := numbers.AssertEquals(
		http.StatusBadRequest,
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
