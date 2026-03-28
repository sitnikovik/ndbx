package body

import (
	"net/url"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event/category"
)

// Content represents the content of the request body for the list events endpoint.
type Content struct {
	// Title is the title of the event to filter by.
	title string
	// cat is the category of the event to filter by.
	cat category.Type
}

// URLQuery converts the Content into url.Values.
func (c Content) URLQuery() url.Values {
	q := make(url.Values, 2)
	if v := c.title; v != "" {
		q.Set("title", c.title)
	}
	if !c.cat.Unspecified() {
		q.Set("category", c.cat.String())
	}
	return q
}
