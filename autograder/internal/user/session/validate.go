package session

import (
	"errors"
	"strings"
)

// Validate checks if the provided session ID is valid and returns an error if it is not.
//
// A valid session ID is expected to be a hexadecimal string of at least 32 characters in length.
func Validate(sid string) error {
	if len(sid) < 32 {
		return errors.New("must be at least 32 characters long")
	}
	for _, chr := range sid {
		if !strings.ContainsRune("0123456789abcdefABCDEF", chr) {
			return errors.New("must be a hexadecimal string")
		}
	}
	return nil
}
