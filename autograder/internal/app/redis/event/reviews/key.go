package reviews

import (
	redisEvent "github.com/sitnikovik/ndbx/autograder/internal/app/redis/event"
)

// Key returns the key for the reviews set in Redis.
func Key(sfx string) string {
	return redisEvent.Key(sfx) + ":reviews"
}
