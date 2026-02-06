package body

import "encoding/json"

// MustParseIn parses a JSON body into the provided type and panics if parsing fails.
//
// Parameters:
//   - v: A pointer to the variable where the parsed data will be stored.
func (b Body) MustParseIn(v any) {
	err := json.NewDecoder(b.reader).Decode(v)
	if err != nil {
		panic(err)
	}
}
