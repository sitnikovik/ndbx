package response

import (
	"errors"
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
)

// AssertNoContentStatus asserts that the HTTP response has a 204 No Content status code
// and returns an error if it does not.
func AssertNoContentStatus(rsp *http.Response) error {
	err := numbers.AssertEquals(
		http.StatusNoContent,
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
