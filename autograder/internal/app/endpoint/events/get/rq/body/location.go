package body

import "net/url"

// Location represents the filter criteria
// for the events endpoint related to location.
type Location struct {
	// address is the address of the event.
	address string
	// city is the city name.
	city string
}

// URLQuery converts the Location into url.Values.
func (c Location) URLQuery() url.Values {
	q := make(url.Values, 2)
	if v := c.address; v != "" {
		q.Set("address", v)
	}
	if v := c.city; v != "" {
		q.Set("city", v)
	}
	return q
}
