package httpx

import (
	"io"
	"net/http"
)

// FakeClient is a mock implementation of an HTTP client
// that allows you to define custom behavior for its methods.
type FakeClient struct {
	// funcs holds the functions that define the behavior of the FakeClient's methods.
	funcs funcs
}

// NewFakeClient creates a new instance of FakeClient
// with the provided options to configure its behavior.
func NewFakeClient(opts ...Option) *FakeClient {
	c := &FakeClient{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Get simulates an HTTP GET request to the specified URL,
// returning the response body as a string or an error if the operation fails.
//
// Panics if the behavior for the Get method is not specified.
func (c *FakeClient) Get(url string) (*http.Response, error) {
	if c.funcs.Get == nil {
		panic("not specified behavior for Get method")
	}
	return c.funcs.Get(url)
}

// PostJSON simulates an HTTP POST request with a JSON body to the specified URL,
// returning the response body as a string or an error if the operation fails.
//
// Panics if the behavior for the PostJSON method is not specified.
func (c *FakeClient) PostJSON(
	url string,
	body io.Reader,
) (*http.Response, error) {
	if c.funcs.PostJSON == nil {
		panic("not specified behavior for PostJSON method")
	}
	return c.funcs.PostJSON(url, body)
}

// Patch simulates an HTTP PATCH request with a JSON body to the specified URL,
// returning the response body as a string or an error if the operation fails.
//
// Panics if the behavior for the Patch method is not specified.
func (c *FakeClient) Patch(
	url string,
	body io.Reader,
) (*http.Response, error) {
	if c.funcs.Patch == nil {
		panic("not specified behavior for Patch method")
	}
	return c.funcs.Patch(url, body)
}
