package cassandra

import "context"

// Scanner implements the scanner
// that can be used to scan the rows.
type Scanner interface {
	// Scan scans the next like and writes
	// the values to the provided variables.
	Scan(...any) bool
	// Close closes the scanner.
	Close() error
}

// Selectable implements the entities
// stored in Apache Cassandra to select.
type Selectable interface {
	// Select runs SELECT query
	// and returns iterator to scan rows.
	Select(
		ctx context.Context,
		query string,
		args ...any,
	) (Scanner, error)
}
