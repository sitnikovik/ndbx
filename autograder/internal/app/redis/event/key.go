package event

// Key returns the key for the event set in Redis.
func Key(sfx string) string {
	return "event:" + sfx
}
