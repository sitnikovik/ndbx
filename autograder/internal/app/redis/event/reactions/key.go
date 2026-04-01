package reactions

import (
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	redisEvent "github.com/sitnikovik/ndbx/autograder/internal/app/redis/event"
)

// Key returns the key for the reactions set in Redis.
func Key(id event.ID) string {
	return redisEvent.Key(id) + ":reactions"
}
