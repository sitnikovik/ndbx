package cookie

import "net/http"

// Cookies is a wrapper around a slice of HTTP cookies.
type Cookies struct {
	// all is the slice of HTTP cookies that the Cookies struct manages.
	all []*http.Cookie
}

// NewCookies creates a new Cookies instance from the provided HTTP cookies.
func NewCookies(cookies []*http.Cookie) *Cookies {
	return &Cookies{
		all: cookies,
	}
}

// Has checks if a cookie with the specified name exists in the collection of cookies.
func (c *Cookies) Has(name string) bool {
	for _, ck := range c.all {
		if ck.Name == name {
			return true
		}
	}
	return false
}

// MustGet retrieves a cookie by its name from the collection of cookies.
// It panics if the cookie is not found.
func (c *Cookies) MustGet(name string) *http.Cookie {
	for _, ck := range c.all {
		if ck.Name == name {
			return ck
		}
	}
	panic("cookie not found: " + name)
}
