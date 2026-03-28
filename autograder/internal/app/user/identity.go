package user

import (
	"crypto/md5"
	"encoding/hex"
)

// Identity represents the unique identifier of a user in the autograder application.
type Identity struct {
	// username is the username of the user.
	username string
	// id is the unique identifier of the user.
	id ID
}

// IdentityOption represents a functional option for an Idendity.
type IdentityOption func(id *Identity)

// WithUsername sets the provided username to the Identity.
func WithUsername(usr string) IdentityOption {
	return func(id *Identity) {
		id.username = usr
	}
}

// NewIdentity creates and returns a new Identity instance.
func NewIdentity(id ID, opts ...IdentityOption) Identity {
	idnt := Identity{
		id: id,
	}
	for _, opt := range opts {
		opt(&idnt)
	}
	return idnt
}

// ID returns the unique identifier of the user.
func (i Identity) ID() ID {
	return i.id
}

// Username returns the username of the user.
func (i Identity) Username() string {
	return i.username
}

// Hash represents the user's identity as a hash.
func (i Identity) Hash() string {
	usr := i.Username()
	if usr == "" {
		return ""
	}
	hash := md5.Sum([]byte(usr))
	return hex.EncodeToString(hash[:])
}
