package step

import (
	"maps"
	"sync"
)

// Variables represents the variables that can be used across steps.
type Variables interface {
	// Get returns the variable with the given name
	// and a boolean indicating whether the variable was found.
	Get(name string) (Variable, bool)
	// MustGet returns the variable with the given name,
	// or panics if the variable is not found.
	MustGet(name string) Variable
	// Set sets the variable with the given name and value.
	Set(name string, value any)
	// Len returns the number of variables in the struct.
	Len() int
	// Empty returns true if the struct has no variables, false otherwise.
	Empty() bool
	// With adds the given variables to the struct and returns a new one.
	With(vv ...Variable) *Vars
	// Copy creates a copy of the struct and returns it.
	Copy() *Vars
}

// Vars holds the variables for the lab steps.
type Vars struct {
	// mu is a mutex to protect the map of variables.
	mu sync.RWMutex
	// m is a map of variable name to variable value.
	m map[string]Variable
}

// NewVariables creates a new Vars struct with an empty map.
func NewVariables() Variables {
	return &Vars{
		m: make(map[string]Variable),
	}
}

// Len returns the number of variables in the Vars struct.
func (v *Vars) Len() int {
	return len(v.m)
}

// Empty returns true if the Vars struct has no variables, false otherwise.
func (v *Vars) Empty() bool {
	return len(v.m) == 0
}

// Get returns the variable with the given name
// and a boolean indicating whether the variable was found.
func (v *Vars) Get(name string) (Variable, bool) {
	v.mu.RLock()
	variable, ok := v.m[name]
	v.mu.RUnlock()
	return variable, ok
}

// Set sets the variable with the given name and value.
func (v *Vars) Set(name string, value any) {
	v.mu.Lock()
	v.m[name] = NewVariable(name, value)
	v.mu.Unlock()
}

// MustGet returns the variable with the given name,
// or panics if the variable is not found.
func (v *Vars) MustGet(name string) Variable {
	variable, ok := v.Get(name)
	if !ok {
		panic("variable not found: " + name)
	}
	return variable
}

// With adds the given variables to the Vars struct and returns a new Vars struct.
func (v *Vars) With(vv ...Variable) *Vars {
	newVars := v.Copy()
	for _, variable := range vv {
		newVars.m[variable.k] = variable
	}
	return newVars
}

// Copy creates a copy of the Vars struct and returns it.
func (v *Vars) Copy() *Vars {
	newVars, ok := NewVariables().(*Vars)
	if !ok {
		panic("failed to create a new Vars struct")
	}
	maps.Copy(newVars.m, v.m)
	return newVars
}
