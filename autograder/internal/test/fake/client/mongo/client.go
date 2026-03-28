package mongo

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/shard"
)

// FakeClient is a fake implementation of the MongoDB client used for testing purposes.
type FakeClient struct {
	// funcs is a struct that contains the functions that will be used to mock the behavior of the MongoDB client.
	funcs funcs
}

// NewFakeClient creates a new instance of FakeClient with the provided options.
func NewFakeClient(opts ...Option) *FakeClient {
	fc := new(FakeClient)
	for _, opt := range opts {
		opt(fc)
	}
	return fc
}

// AllBy simulates a MongoDB query that retrieves documents
// from a collection based on specified key-value pairs.
func (fc *FakeClient) AllBy(
	ctx context.Context,
	collection string,
	by doc.KVs,
) (doc.Documents, error) {
	if fc.funcs.allBy == nil {
		panic("not specified behavior for AllBy method")
	}
	return fc.funcs.allBy(ctx, collection, by)
}

// OneBy simulates a MongoDB query that retrieves a single document
// from a collection based on the specified ID.
func (fc *FakeClient) OneBy(
	ctx context.Context,
	collection string,
	by doc.KVs,
) (doc.Document, error) {
	if fc.funcs.oneBy == nil {
		panic("not specified behavior for OneBy method")
	}
	return fc.funcs.oneBy(ctx, collection, by)
}

// ByID simulates a MongoDB query that retrieves a single document
// from a collection based on the specified ID.
func (fc *FakeClient) ByID(
	ctx context.Context,
	collection string,
	id string,
) (doc.Document, error) {
	if fc.funcs.byID == nil {
		panic("not specified behavior for ByID method")
	}
	return fc.funcs.byID(ctx, collection, id)
}

// Indexes simulates a MongoDB query that retrieves the list of indexes
// for a specified collection.
func (fc *FakeClient) Indexes(
	ctx context.Context,
	collection string,
) (doc.Indexes, error) {
	if fc.funcs.indexes == nil {
		panic("not specified behavior for Indexes method")
	}
	return fc.funcs.indexes(ctx, collection)
}

// Insert inserts the list of documents into the specified collection.
// Returns a slice of inserted document IDs and an error if any.
func (fc *FakeClient) Insert(
	ctx context.Context,
	collection string,
	kvs ...doc.KVs,
) ([]string, error) {
	if fc.funcs.insert == nil {
		panic("not specified behavior for Insert method")
	}
	return fc.funcs.insert(ctx, collection, kvs...)
}

// HostsOfShard returns a list of hosts where the shard is running.
func (fc *FakeClient) HostsOfShard(
	ctx context.Context,
	id string,
) ([]string, error) {
	if fc.funcs.hostsOfShard == nil {
		panic("not specified behavior for HostsOfShard method")
	}
	return fc.funcs.hostsOfShard(ctx, id)
}

// Shards retrieves the shard
// information for the specified collection.
func (fc *FakeClient) Shards(
	ctx context.Context,
	collection string,
) (shard.Shards, error) {
	if fc.funcs.insert == nil {
		panic("not specified behavior for Shards method")
	}
	return fc.funcs.shards(ctx, collection)
}
