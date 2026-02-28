package body

import "encoding/json"

// MustParseIn parses a JSON body into the provided type and panics if parsing fails.
//
// Parameters:
//   - v: A pointer to the variable where the parsed data will be stored.
func (b Body) MustParseIn(v any) {
	err := b.ParseIn(v)
	if err != nil {
		panic(err)
	}
}

// ParseIn parses a JSON body into the provided type.
//
// Parameters:
//   - v: A pointer to the variable where the parsed data will be stored.
//
// Returns an error if parsing fails.
func (b Body) ParseIn(v any) error {
	return json.NewDecoder(b.reader).Decode(v)
}
