package doc

import "github.com/sitnikovik/ndbx/autograder/internal/app/user"

// FromUsers converts the list of User to the list of UserDocument related to MongoDB.
func FromUsers(uu []user.User) UserDocuments {
	if len(uu) == 0 {
		return nil
	}
	res := make([]UserDocument, len(uu))
	for i, u := range uu {
		res[i] = FromUser(u)
	}
	return res
}
