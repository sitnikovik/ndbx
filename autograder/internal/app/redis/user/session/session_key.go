package session

// Key constructs the Redis key for a given session ID.
func Key(sessionID string) string {
	return SessionKeyPrefix + sessionID
}
