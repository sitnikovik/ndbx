package user

import (
	"encoding/json"
)

// MustParseJSON parses the JSON content of the User.
//
// Panics if parsing the JSON fails or if the required fields are missing.
func MustParseJSON(bb []byte) User {
	var v struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		FullName string `json:"full_name"`
	}
	err := json.Unmarshal(bb, &v)
	if err != nil {
		panic(err)
	}
	return NewUser(
		NewID(v.ID),
		v.Username,
		v.FullName,
	)
}
