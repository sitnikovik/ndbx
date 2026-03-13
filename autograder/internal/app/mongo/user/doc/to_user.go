package doc

import (
	"github.com/sitnikovik/ndbx/autograder/internal/app/mongo/user/doc/key"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

// ToUser converts the UserDocument struct to a user.User struct.
func (u UserDocument) ToUser() user.User {
	var (
		name     string
		username string
	)
	for _, kv := range u.orig.KVs() {
		switch kv.Key() {
		case key.FullName:
			name = kv.Value().(string)
		case key.Username:
			username = kv.Value().(string)
		}
	}
	return user.NewUser(
		user.NewID(u.orig.ID()),
		username,
		name,
	)
}
