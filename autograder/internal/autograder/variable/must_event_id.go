package variable

// MustEventID retrieves the EventID variable from the original variables and returns it as a string.
func (v Values) MustEventID() string {
	return v.orig.MustGet(EventID).AsString()
}
