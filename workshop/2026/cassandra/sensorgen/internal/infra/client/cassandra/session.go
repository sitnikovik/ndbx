package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

func QuorumSession(
	ks string,
	p int,
	c9y string,
	hh ...string,
) (*gocql.Session, error) {
	clu := gocql.NewCluster(hh...)
	clu.Port = p
	clu.Keyspace = ks
	clu.Consistency = gocql.ParseConsistency(c9y)
	clu.DisableInitialHostLookup = true
	session, err := clu.CreateSession()
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create Cassandra session: %w",
			err,
		)
	}
	return session, nil
}
