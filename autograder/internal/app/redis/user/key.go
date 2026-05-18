package user

// Key returns the key for the user set in Redis.
func Key(sfx string) string {
	return "user:" + sfx
}
