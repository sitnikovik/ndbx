package bson

import (
	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
)

// ToKV converts a bson, which is a map used in MongoDB operations, to a list of key-value pairs.
//
// Useful for converting query results from MongoDB into a more convenient format.
//
// Note that the "_id" field is excluded from the resulting list of key-value pairs,
// as it is typically handled separately.
func (m M) ToKV() doc.KVs {
	if len(m) == 0 {
		return nil
	}
	kvs := make(doc.KVs, 0, len(m))
	for k, v := range m {
		if k == "_id" {
			continue
		}
		normv := v
		if bsonD, ok := v.(bson.D); ok {
			if len(bsonD) == 1 && bsonD[0].Key == "_id" {
				if oid, ok := bsonD[0].Value.(bson.ObjectID); ok {
					normv = oid.Hex()
				} else {
					normv = bsonD[0].Value
				}
			}
		}
		kvs = append(kvs, doc.NewKV(k, normv))
	}
	return kvs
}
