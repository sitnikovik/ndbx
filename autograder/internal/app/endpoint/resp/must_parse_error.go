package resp

import (
	"io"

	jsonbody "github.com/sitnikovik/ndbx/autograder/internal/json/body"
)

// MustParseError reads the HTTP response body and extracts the error message.
//
// Panics if parsing the body fails or if the required fields are missing.
func MustParseError(body io.ReadCloser) ErrorResponse {
	var v struct {
		Message string `json:"message"`
	}
	jsonbody.NewBody(body).MustParseIn(&v)
	return ErrorResponse{
		msg: v.Message,
	}
}
