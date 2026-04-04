package reaction

import (
	"github.com/sitnikovik/ndbx/autograder/internal/app/cassandra"
	"github.com/sitnikovik/ndbx/autograder/internal/app/cassandra/event/reaction/filter"
)

// Reactions represents event reactions from database.
type Reactions struct {
	// db is database client.
	db cassandra.Selectable
	// ftr is filter for reactions.
	ftr filter.Filter
	// limit is the maximum number of reactions to return.
	limit int
}

// NewReactions creates new Reactions instance.
func NewReactions(
	db cassandra.Selectable,
	opts ...Option,
) *Reactions {
	r := &Reactions{
		db: db,
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}
