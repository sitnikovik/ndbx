package doc

import (
	"github.com/sitnikovik/ndbx/autograder/internal/app/mongo/user/doc/key"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
)

// FromUser converts the User to a UserDocument related to MongoDB.
func FromUser(u user.User) UserDocument {
	kvs := make(doc.KVs, 0, 2)
	if v := u.FullName(); v != "" {
		kvs = append(kvs, doc.NewKV(
			key.FullName,
			v,
		))
	}
	if v := u.Username(); v != "" {
		kvs = append(kvs, doc.NewKV(
			key.Username,
			v,
		))
	}
	return NewUserDocument(
		doc.NewDocument(
			u.ID().String(),
			kvs...,
		),
	)
}
