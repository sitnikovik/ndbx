package httpx

import (
	"errors"
	"io"
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/http/content"
)

// PostJSON sends a POST request with a JSON body
// to the specified URL and returns the response.
func (c *Client) PostJSON(
	url string,
	body io.Reader,
) (*http.Response, error) {
	resp, err := c.cli.Post(
		url,
		content.ApplicationJSON.String(),
		body,
	)
	if err != nil {
		return nil, errors.Join(err, errs.ErrHTTPFailed)
	}
	return resp, nil
}
