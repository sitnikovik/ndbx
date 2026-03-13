package variable

// MustUserPassword retrieves the user password variable from the original variables
// and returns it as a string.
//
// Panics if the user password variable is not set or has an invalid type.
func (v Values) MustUserPassword() string {
	pass, ok := v.orig.MustGet(UserPassword).Value().(string)
	if !ok {
		panic("user password variable is not set or has invalid type")
	}
	return pass
}
