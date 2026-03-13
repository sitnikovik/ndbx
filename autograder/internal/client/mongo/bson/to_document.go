package bson

import (
	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
)

// ToDocument converts a bson document to a Document.
//
// It expects the bson document to contain a valid "_id" field of type ObjectID,
// and will panic if the "_id" field is missing or invalid.
func (m M) ToDocument() doc.Document {
	b := bson.M(m)
	id, ok := b["_id"].(bson.ObjectID)
	if !ok {
		panic("document '_id' must be a type of ObjectID")
	}
	return doc.NewDocument(
		id.Hex(),
		m.ToKV()...,
	)
}

// ToDocuments converts a slice of bson documents to a list of Documents.
//
// It expects each bson document to contain a valid "_id" field of type ObjectID,
// and will panic if any document is missing the "_id" field or contains an invalid "_id".
func (mm MList) ToDocuments() doc.Documents {
	if len(mm) == 0 {
		return nil
	}
	docs := make(doc.Documents, len(mm))
	for i, b := range mm {
		docs[i] = b.ToDocument()
	}
	return docs
}
