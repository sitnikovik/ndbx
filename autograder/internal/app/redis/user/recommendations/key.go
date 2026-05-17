package recommendations

import (
	"github.com/sitnikovik/ndbx/autograder/internal/app/redis/user"
	"github.com/sitnikovik/ndbx/autograder/internal/app/redis/user/recommendations/field"
)

// Key returns the key for the reactions set in Redis.
func Key(sfx string) string {
	return user.Key(sfx) + ":" + field.Recommendations
}
