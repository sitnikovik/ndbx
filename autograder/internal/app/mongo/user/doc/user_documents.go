package doc

import "github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"

// UserDocuments represents a list of user documents in MongoDB.
type UserDocuments []UserDocument

// KVs returns the slice of key-value pairs representing the fields of the document.
func (uu UserDocuments) KVs() []doc.KVs {
	if len(uu) == 0 {
		return nil
	}
	res := make([]doc.KVs, len(uu))
	for i, u := range uu {
		res[i] = u.KVs()
	}
	return res
}
