package httpx

import (
	"errors"
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

// Get sends a GET request to the specified URL and returns the response.
func (c *Client) Get(url string) (*http.Response, error) {
	resp, err := c.cli.Get(url)
	if err != nil {
		return nil, errors.Join(err, errs.ErrHTTPFailed)
	}
	return resp, nil
}
