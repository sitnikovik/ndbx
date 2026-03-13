package response

import (
	"errors"
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
)

// AssertUnauthorizedStatus checks if the HTTP response has a status code of 401 Unauthorized
// and returns an error if it does not.
func AssertUnauthorizedStatus(rsp *http.Response) error {
	err := numbers.AssertEquals(
		http.StatusUnauthorized,
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
