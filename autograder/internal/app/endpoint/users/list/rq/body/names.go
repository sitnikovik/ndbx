package body

import "net/url"

// Names represents parameters to filter users by them names.
type Names struct {
	// fullName is full name of the user.
	fullName string
}

// URLQuery returns Names representation in url queries.
func (n Names) URLQuery() url.Values {
	q := make(url.Values, 1)
	if n.fullName != "" {
		q.Set("name", n.fullName)
	}
	return q
}
