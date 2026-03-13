package session

const (
	// SessionKeyPrefix is the prefix for Redis keys that store session information.
	SessionKeyPrefix = "sid:"
	// CreatedAtField is the field name for the creation timestamp in Redis.
	CreatedAtField = "created_at"
	// UpdatedAtField is the field name for the update timestamp in Redis.
	UpdatedAtField = "updated_at"
	// UserIDField is the field name for the user ID associated with the session in Redis.
	UserIDField = "user_id"
)
