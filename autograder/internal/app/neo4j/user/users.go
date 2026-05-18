package user

import (
	"context"
	"strconv"

	"github.com/sitnikovik/ndbx/autograder/internal/app/neo4j/node/user"
	"github.com/sitnikovik/ndbx/autograder/internal/client/neo4j/graph"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

// neo4jClient implements for the Neo4j client
// to interact with the database.
type neo4jClient interface {
	// QueryNodes executes a Cypher query and returns the nodes.
	QueryNodes(
		ctx context.Context,
		query string,
		params map[string]any,
		keys ...string,
	) (graph.Nodes, error)
}

// Users represents the users stored in Neo4j database
// of the target application.
type Users struct {
	// db is the Neo4j client to interact with the database.
	db neo4jClient
}

// NewUsers creates a new Users instance.
func NewUsers(db neo4jClient) *Users {
	return &Users{
		db: db,
	}
}

// All returns users from the database by the given limit.
func (u *Users) All(ctx context.Context, limit int) (user.Users, error) {
	q := "MATCH (u:User) RETURN u.id"
	if limit > 0 {
		q += " LIMIT " + strconv.Itoa(limit)
	}
	nn, err := u.db.QueryNodes(ctx, q, nil, "u")
	if err != nil {
		return nil, errs.Wrap(err, "failed to get users from Neo4j")
	}
	return NodesToUsers(nn...), nil
}
