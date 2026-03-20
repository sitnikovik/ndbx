package httpx

import (
	"io"
	"net/http"
	"net/url"

	"github.com/sitnikovik/ndbx/autograder/internal/http/content"
)

// Patch sends a PATCH request with a JSON body
// to the specified URL and returns the response.
func (c *Client) Patch(rawURL string, body io.Reader) (*http.Response, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	return c.cli.Do(&http.Request{
		Method: http.MethodPatch,
		URL:    u,
		Header: http.Header{
			"Content-Type": []string{
				content.ApplicationJSON.String(),
			},
		},
		Body: body.(io.ReadCloser),
	})
}
