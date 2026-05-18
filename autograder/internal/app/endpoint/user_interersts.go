package endpoint

import "github.com/sitnikovik/ndbx/autograder/internal/app/user"

// UserInterests returns the URL for the user's interests endpoint
// of the autograder application.
func UserInterests(id user.ID) string {
	return "/users/" + id.String() + "/interests"
}
