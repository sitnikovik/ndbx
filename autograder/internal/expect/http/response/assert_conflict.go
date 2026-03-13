package response

import (
	"errors"
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
)

// AssertConflictStatus checks if the HTTP response has a status code of 409 Conflict
// and returns an error if it does not.
func AssertConflictStatus(rsp *http.Response) error {
	err := numbers.AssertEquals(
		http.StatusConflict,
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
