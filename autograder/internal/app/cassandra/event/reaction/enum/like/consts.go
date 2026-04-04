package like

// Value represents the value
// for 'like' field in the reactions table.
type Value int8

const (
	// Unspecified defines the unspecified value.
	Unspecified Value = 0
	// Like defines the like value.
	Like Value = 1
	// Dislike defines the dislike value.
	Dislike Value = -1
)
