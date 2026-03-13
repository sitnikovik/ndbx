package response

import (
	"errors"
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
)

// AssertCreatedStatus checks if the HTTP response has a status code of 201 Created
// and returns an error if it does not.
func AssertCreatedStatus(rsp *http.Response) error {
	err := numbers.AssertEquals(
		http.StatusCreated,
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
