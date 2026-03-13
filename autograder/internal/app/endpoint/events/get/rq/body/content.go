package body

import "net/url"

// Content represents the content of the request body for the list events endpoint.
type Content struct {
	// Title is the title of the event to filter by.
	title string
}

// URLQuery converts the Content into url.Values.
func (c Content) URLQuery() url.Values {
	q := make(url.Values, 1)
	if v := c.title; v != "" {
		q.Set("title", c.title)
	}
	return q
}
