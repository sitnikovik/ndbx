package event

// Equals checks if two events are equal by comparing their fields.
func (e Event) Equals(other Event) bool {
	return e == other
}
