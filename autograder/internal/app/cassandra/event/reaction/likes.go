package reaction

import (
	"context"
	"fmt"
	"strconv"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cassandra"
	eventReaction "github.com/sitnikovik/ndbx/autograder/internal/app/reaction/event"
)

// Likes represents event likes from database.
type Likes struct {
	// reactions defines common reactions options.
	reactions *Reactions
}

// NewLikes returns new Likes from database.
func NewLikes(
	db cassandra.Selectable,
	opts ...Option,
) *Likes {
	return &Likes{
		reactions: NewReactions(db, opts...),
	}
}

// Select returns all likes for the event.
func (r *Likes) Select(ctx context.Context) ([]eventReaction.Like, error) {
	q := fmt.Sprintf(
		`SELECT %s, %s, %s, %s FROM %s %s`,
		"event_id",
		"like_value",
		"created_at",
		"created_by",
		Table,
		r.reactions.ftr.Where(),
	)
	if n := r.reactions.limit; n > 0 {
		q += ` LIMIT ` + strconv.Itoa(n)
	}
	itr, err := r.reactions.db.Select(
		ctx,
		q,
		r.reactions.ftr.Args()...,
	)
	if err != nil {
		return nil, err
	}
	return ToLikes(itr), nil
}
