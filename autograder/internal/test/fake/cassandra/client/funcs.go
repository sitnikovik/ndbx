package client

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cassandra"
)

// funcs holds function fields
// for mocking the behavior of the Client methods.
type funcs struct {
	// selectfn is a function that will be used
	// to mock the behavior of the Select method.
	selectfn func(
		_ context.Context,
		query string,
		args ...any,
	) (cassandra.Scanner, error)
}
