package bson

import (
	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
)

// ToIndexes converts a list of bson documents to a list of indexes.
func (mm MList) ToIndexes() doc.Indexes {
	if len(mm) == 0 {
		return nil
	}
	indexes := make(doc.Indexes, 0, len(mm))
	for _, v := range mm {
		indexes = append(indexes, v.ToIndex()...)
	}
	return indexes
}

// ToIndex converts a bson document to a list of indexes.
func (m M) ToIndex() doc.Indexes {
	if len(m) == 0 {
		return nil
	}
	keyField, ok := m["key"]
	if !ok {
		return nil
	}
	var keys []string
	if bsonD, ok := keyField.(bson.D); ok {
		keys = make([]string, 0, len(bsonD))
		for _, elem := range bsonD {
			keys = append(keys, elem.Key)
		}
	} else if keyMap, ok := keyField.(map[string]any); ok {
		keys = make([]string, 0, len(keyMap))
		for k := range keyMap {
			keys = append(keys, k)
		}
	} else if bsonM, ok := keyField.(M); ok {
		keys = make([]string, 0, len(bsonM))
		for k := range bsonM {
			keys = append(keys, k)
		}
	} else {
		return nil
	}
	if len(keys) == 0 {
		return nil
	}
	unique := false
	if uniqueField, ok := m["unique"]; ok {
		if uniqueBool, ok := uniqueField.(bool); ok {
			unique = uniqueBool
		}
	}
	var idx doc.Index
	if unique {
		idx = doc.NewUniqueIndex(keys...)
	} else {
		idx = doc.NewIndex(keys...)
	}
	return doc.Indexes{idx}
}
