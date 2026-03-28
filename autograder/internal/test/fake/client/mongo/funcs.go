package mongo

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/shard"
)

// funcs holds function fields for mocking the behavior of the FakeClient methods.
type funcs struct {
	// allBy is a function that will be used to mock the behavior of the AllBy method.
	allBy func(
		ctx context.Context,
		collection string,
		by doc.KVs,
	) (doc.Documents, error)
	// oneBy is a function that will be used to mock the behavior of the OneBy method.
	oneBy func(
		ctx context.Context,
		collection string,
		by doc.KVs,
	) (doc.Document, error)
	// byID is a function that will be used to mock the behavior of the ByID method.
	byID func(
		ctx context.Context,
		collection string,
		id string,
	) (doc.Document, error)
	// indexes is a function that will be used to mock the behavior of the Indexes method.
	indexes func(
		ctx context.Context,
		collection string,
	) (doc.Indexes, error)
	// insert is a function that will be used to mock the behavior of the Insert method.
	insert func(
		ctx context.Context,
		collection string,
		kvs ...doc.KVs,
	) ([]string, error)
	// hostsOfShard is a function that will be used to mock the behavior of the HostsOfShard method.
	hostsOfShard func(
		ctx context.Context,
		id string,
	) ([]string, error)
	// shards is a function that will be used to mock the behavior of the Shards method.
	shards func(
		ctx context.Context,
		collection string,
	) (shard.Shards, error)
}
