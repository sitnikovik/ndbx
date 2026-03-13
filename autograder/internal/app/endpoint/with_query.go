package endpoint

import "net/url"

// WithQuery returns the URL for the given endpoint with the provided URL query parameters.
func WithQuery(endpoint string, q url.Values) string {
	if len(q) == 0 {
		return endpoint
	}
	u, err := url.Parse(endpoint)
	if err != nil {
		panic(err)
	}
	u.RawQuery = q.Encode()
	return u.String()
}
