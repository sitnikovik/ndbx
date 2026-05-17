package user

import (
	neo4j "github.com/sitnikovik/ndbx/autograder/internal/app/neo4j/node/user"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/client/neo4j/graph"
)

// NodesToUsers converts graph nodes to Neo4j users objects.
//
// Panics if the node has incorrect field of value type.
func NodesToUsers(nn ...graph.Node) neo4j.Users {
	res := make([]neo4j.User, len(nn))
	for i, n := range nn {
		res[i] = neo4j.NewUser(
			user.NewID(
				n.Properties().
					ByName("id").
					Value().
					Base().
					MustString(),
			),
		)
	}
	return res
}
