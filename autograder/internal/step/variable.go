package step

import "time"

// Variable holds the variables for the job steps.
type Variable struct {
	// k is the name of the variable, v is the value of the variable.
	k string
	// v is the value of the variable.
	v any
}

// NewVariable creates a new Variable with the given name and value.
func NewVariable(name string, value any) Variable {
	return Variable{
		k: name,
		v: value,
	}
}

// Name returns the name of the variable.
func (v Variable) Name() string {
	return v.k
}

// Value returns the value of the variable.
func (v Variable) Value() any {
	return v.v
}

// Empty returns true if the variable is empty, false otherwise.
func (v Variable) Empty() bool {
	return v.k == "" && v.v == nil
}

// AsString returns the string value of the variable, or an empty string if the variable is not a string.
func (v Variable) AsString() string {
	if str, ok := v.v.(string); ok {
		return str
	}
	return ""
}

// AsInt returns the int value of the variable, or 0 if the variable is not an int.
func (v Variable) AsInt() int {
	if i, ok := v.v.(int); ok {
		return i
	}
	return 0
}

// AsDuration returns the Duration value of the variable,
// or 0 if the variable is not a Duration.
func (v Variable) AsDuration() time.Duration {
	if d, ok := v.v.(time.Duration); ok {
		return d
	}
	return 0
}

// AsTime returns the Time value of the variable,
// or the zero time if the variable is not a Time.
func (v Variable) AsTime() time.Time {
	if t, ok := v.v.(time.Time); ok {
		return t
	}
	return time.Time{}
}
