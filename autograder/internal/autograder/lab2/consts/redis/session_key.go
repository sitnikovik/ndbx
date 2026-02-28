package redis

// SessionKey constructs the Redis key for a given session ID.
func SessionKey(sessionID string) string {
	return SessionKeyPrefix + sessionID
}
