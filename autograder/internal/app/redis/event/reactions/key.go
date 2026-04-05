package reactions

import (
	redisEvent "github.com/sitnikovik/ndbx/autograder/internal/app/redis/event"
)

// Key returns the key for the reactions set in Redis.
func Key(sfx string) string {
	return redisEvent.Key(sfx) + ":reactions"
}
