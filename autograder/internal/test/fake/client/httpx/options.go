package httpx

import (
	"io"
	"net/http"
)

// Option are used to configure the behavior of the FakeClient.
type Option func(*FakeClient)

// WithGet sets the function that will be executed
// when the FakeClient's Get method is called.
func WithGet(
	fn func(
		url string,
	) (*http.Response, error),
) Option {
	return func(c *FakeClient) {
		c.funcs.Get = fn
	}
}

// WithPostJSON sets the function that will be executed
// when the FakeClient's PostJSON method is called.
func WithPostJSON(
	fn func(
		url string,
		body io.Reader,
	) (*http.Response, error),
) Option {
	return func(c *FakeClient) {
		c.funcs.PostJSON = fn
	}
}
