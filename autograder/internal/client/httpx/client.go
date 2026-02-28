package httpx

import (
	"net/http"
	"net/http/cookiejar"
	"time"
)

const (
	// DefaultTimeout is the default timeout for HTTP requests made by the Client.
	DefaultTimeout = 5 * time.Second
)

// Client is a wrapper around the standard http.Client
// that provides additional configuration options.
type Client struct {
	// cli is the client used to make HTTP requests.
	cli *http.Client
}

// ClientOption is a function type that defines an option for configuring the Client.
type ClientOption func(*Client)

// WithTimeout sets the timeout for the client's HTTP requests.
func WithTimeout(dur time.Duration) ClientOption {
	return func(c *Client) {
		c.cli.Timeout = dur
	}
}

// WithEmptyCookieJar sets an empty cookie jar for the client, which will not store any cookies.
func WithEmptyCookieJar() ClientOption {
	return func(c *Client) {
		jar, _ := cookiejar.New(nil)
		c.cli.Jar = jar
	}
}

// NewClient creates a new Client with the provided options.
func NewClient(opts ...ClientOption) *Client {
	c := &Client{
		cli: &http.Client{
			Timeout: DefaultTimeout,
		},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Timeout returns the timeout duration for the client's HTTP requests.
func (c *Client) Timeout() time.Duration {
	return c.cli.Timeout
}
