package doc

import "github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"

// EventDocuments represents a list of event documents in MongoDB.
type EventDocuments []EventDocument

// KVs returns the slice of key-value pairs representing the fields of the document.
func (ee EventDocuments) KVs() []doc.KVs {
	if len(ee) == 0 {
		return nil
	}
	res := make([]doc.KVs, len(ee))
	for i, e := range ee {
		res[i] = e.KVs()
	}
	return res
}
