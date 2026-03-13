package bson

import "go.mongodb.org/mongo-driver/v2/bson"

// MList is a type definition for a slice of bson.M, which is a map used in MongoDB operations.
type MList []M

// NewBsonMList creates a new slice of bson M from a variadic list of bson.M.
func NewBsonMList(bs ...bson.M) MList {
	if len(bs) == 0 {
		return nil
	}
	ms := make(MList, len(bs))
	for i, b := range bs {
		ms[i] = NewBsonM(b)
	}
	return ms
}
