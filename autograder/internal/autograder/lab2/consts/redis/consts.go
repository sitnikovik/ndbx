package redis

const (
	// SessionKeyPrefix is the prefix for Redis keys that store session information.
	SessionKeyPrefix = "sid:"
	// SessionCreatedAtField is the field name for the creation timestamp in Redis.
	SessionCreatedAtField = "created_at"
	// SessionTTL is the field name for the time-to-live of a session in Redis.
	SessionTTL = "ttl"
)
