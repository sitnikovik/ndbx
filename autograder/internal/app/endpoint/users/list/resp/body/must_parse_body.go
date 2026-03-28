package body

import (
	"encoding/json"
	"io"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	jsonbody "github.com/sitnikovik/ndbx/autograder/internal/json/body"
)

// MustParseBody reads the HTTP response body represents the list of events.
//
// Panics if parsing the body fails or if the required fields are missing.
func MustParseBody(body io.ReadCloser) Body {
	var v struct {
		Users []json.RawMessage `json:"users"`
		Count int               `json:"count"`
	}
	jsonbody.NewBody(body).MustParseIn(&v)
	users := make([]user.User, 0, len(v.Users))
	for _, jsn := range v.Users {
		users = append(users, user.MustParseJSON(jsn))
	}
	return NewBody(users, v.Count)
}
