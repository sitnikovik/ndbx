package endpoint

import "github.com/sitnikovik/ndbx/autograder/internal/app/user"

// UserRecommendations returns the URL for the user's recommendations endpoint
// of the autograder application.
func UserRecommendations(id user.ID) string {
	return "/users/" + id.String() + "/recommendations"
}
