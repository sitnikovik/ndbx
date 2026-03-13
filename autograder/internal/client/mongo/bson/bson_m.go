package bson

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

// M is a type definition for bson.M, which is a map used in MongoDB operations.
type M bson.M

// NewBsonM creates a new bson M from a bson.M.
func NewBsonM(b bson.M) M {
	return M(b)
}


