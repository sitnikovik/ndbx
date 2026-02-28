package httpx

import (
	"io"
	"net/http"
)

// funcs holds the functions that will be executed
// when the corresponding methods of the FakeClient are called.
type funcs struct {
	// Get is the function that will be executed
	// when the FakeClient's Get method is called.
	Get func(url string) (*http.Response, error)
	// PostJSON is the function that will be executed
	// when the FakeClient's PostJSON method is called.
	PostJSON func(
		url string,
		body io.Reader,
	) (*http.Response, error)
}
