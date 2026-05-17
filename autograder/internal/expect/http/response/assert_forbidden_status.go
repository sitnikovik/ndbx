package response

import (
	"errors"
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
)

// AssertForbiddenStatus checks if the HTTP response has a status code of 403 Forbidden
// and returns an error if it does not.
func AssertForbiddenStatus(rsp *http.Response) error {
	err := numbers.AssertEquals(
		http.StatusForbidden,
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
