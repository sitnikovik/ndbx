package include

import (
	"net/url"
	"strings"
)

// Include represents the include query parameter.
type Include struct {
	// values is the list of values to include.
	values []string
}

// NewInclude creates a new Include instance.
func NewInclude(values ...string) Include {
	return Include{
		values: values,
	}
}

// URLQuery converts the Include into url.Values.
func (i Include) URLQuery() url.Values {
	q := make(url.Values, 1)
	if len(i.values) > 0 {
		q.Set(
			"include",
			strings.Join(i.values, ","),
		)
	}
	return q
}
