package response

import (
	"errors"
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
)

// AssertOKStatus checks if the HTTP response has a status code of 200 OK
// and returns an error if it does not.
func AssertOKStatus(rsp *http.Response) error {
	err := numbers.AssertEquals(
		http.StatusOK,
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
