package consistency

import (
	"fmt"

	"github.com/gocql/gocql"
)

// Consistency defines Cassandra consistency level.
//
// It is a number of replicas that must respond to a request
// for it to be considered successful.
type Consistency uint8

const (
	// Any allows a write to succeed if written to at least one node,
	// including hinting nodes. Not usable for reads.
	Any Consistency = 1
	// One requires a response from at least one replica.
	One Consistency = 2
	// Two requires responses from at least two replicas.
	Two Consistency = 3
	// Three requires responses from at least three replicas.
	Three Consistency = 4
	// Quorum requires a majority of replicas to respond.
	// Formula: floor(total_replicas / 2) + 1.
	Quorum Consistency = 5
	// All requires a response from all replicas.
	// Provides the highest consistency but lowest availability.
	All Consistency = 6
	// LocalQuorum requires a quorum of replicas within the local datacenter.
	// Ideal for multi-datacenter clusters to avoid cross-DC latency.
	LocalQuorum Consistency = 7
	// EachQuorum requires a quorum in each datacenter.
	// Used during writes when consistency is needed across all DCs.
	EachQuorum Consistency = 8
	// LocalOne requires a response from one replica in the local datacenter.
	// Often used for reads in multi-DC setups.
	LocalOne Consistency = 9
)

// MustParseConsistency parses string and
// returns the consistency level or panics if not parsed.
func MustParseConsistency(s string) Consistency {
	c, err := ParseConsistency(s)
	if err != nil {
		panic(fmt.Sprintf(
			"failed to parse Apache Cassandra consistency level: %v",
			err,
		))
	}
	return c
}

// ParseConsistency parses string and returns the consistency level
// or returns an error if not parsed.
func ParseConsistency(s string) (Consistency, error) {
	switch s {
	case Any.String():
		return Any, nil
	case One.String():
		return One, nil
	case Two.String():
		return Two, nil
	case Three.String():
		return Three, nil
	case Quorum.String():
		return Quorum, nil
	case All.String():
		return All, nil
	case LocalQuorum.String():
		return LocalQuorum, nil
	case EachQuorum.String():
		return EachQuorum, nil
	case LocalOne.String():
		return LocalOne, nil
	default:
		return 0, fmt.Errorf("invalid consistency '%q'", s)
	}
}

// String retunrs string representation of Consistency.
func (c Consistency) String() string {
	switch c {
	case Any:
		return "ANY"
	case One:
		return "ONE"
	case Two:
		return "TWO"
	case Three:
		return "THREE"
	case Quorum:
		return "QUORUM"
	case All:
		return "ALL"
	case LocalQuorum:
		return "LOCAL_QUORUM"
	case EachQuorum:
		return "EACH_QUORUM"
	case LocalOne:
		return "LOCAL_ONE"
	default:
		return ""
	}
}

// ToCQL returns gocql representation of Consistency.
func (c Consistency) ToCQL() gocql.Consistency {
	switch c {
	case Any:
		return gocql.Any
	case One:
		return gocql.One
	case Two:
		return gocql.Two
	case Three:
		return gocql.Three
	case Quorum:
		return gocql.Quorum
	case All:
		return gocql.All
	case LocalQuorum:
		return gocql.LocalQuorum
	case EachQuorum:
		return gocql.EachQuorum
	case LocalOne:
		return gocql.LocalOne
	default:
		return gocql.Quorum
	}
}
