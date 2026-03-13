package doc

import (
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
)

// UserDocument represents a user document in MongoDB.
type UserDocument struct {
	// orig is the original MongoDB document that contains the user data.
	orig doc.Document
}

// NewUserDocument creates a new UserDocument struct from a MongoDB document.
func NewUserDocument(orig doc.Document) UserDocument {
	return UserDocument{
		orig: orig,
	}
}
