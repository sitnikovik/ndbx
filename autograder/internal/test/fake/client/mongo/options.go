package mongo

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/shard"
)

// Option defines a functional option for configuring the FakeClient.
type Option func(*FakeClient)

// WithAllBy sets the function that will be used to mock the behavior of the AllBy method.
func WithAllBy(
	fn func(
		ctx context.Context,
		collection string,
		by doc.KVs,
	) (doc.Documents, error),
) Option {
	return func(fc *FakeClient) {
		fc.funcs.allBy = fn
	}
}

// WithOneBy sets the function that will be used to mock the behavior of the OneBy method.
func WithOneBy(
	fn func(
		ctx context.Context,
		collection string,
		by doc.KVs,
	) (doc.Document, error),
) Option {
	return func(fc *FakeClient) {
		fc.funcs.oneBy = fn
	}
}

// WithByID sets the function that will be used to mock the behavior of the ByID method.
func WithByID(
	fn func(
		ctx context.Context,
		collection string,
		id string,
	) (doc.Document, error),
) Option {
	return func(fc *FakeClient) {
		fc.funcs.byID = fn
	}
}

// WithIndexes sets the function that will be used to mock the behavior of the Indexes method.
func WithIndexes(
	fn func(
		ctx context.Context,
		collection string,
	) (doc.Indexes, error),
) Option {
	return func(fc *FakeClient) {
		fc.funcs.indexes = fn
	}
}

// WithInsert sets the function that will be used to mock the behavior of the Insert method.
func WithInsert(
	fn func(
		ctx context.Context,
		collection string,
		kvs ...doc.KVs,
	) ([]string, error),
) Option {
	return func(fc *FakeClient) {
		fc.funcs.insert = fn
	}
}

// WithHostsOfShard sets the function that will be used
// to mock the behavior of the HostsOfShard method.
func WithHostsOfShard(
	fn func(
		ctx context.Context,
		id string,
	) ([]string, error),
) Option {
	return func(fc *FakeClient) {
		fc.funcs.hostsOfShard = fn
	}
}

// WithShards sets the function that will be used
// to mock the behavior of the Shard method.
func WithShards(
	fn func(
		ctx context.Context,
		collection string,
	) (shard.Shards, error),
) Option {
	return func(fc *FakeClient) {
		fc.funcs.shards = fn
	}
}
