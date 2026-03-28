package mongo

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/app/mongo/event/collection"
	appevdoc "github.com/sitnikovik/ndbx/autograder/internal/app/mongo/event/doc"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab4"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run inserts 1000 fake events and verifies
// they are distributed beetwen shards and shards are replicated.
func (s *Step) Run(
	ctx context.Context,
	_ step.Variables,
) error {
	err := s.mongo.Insert(
		ctx,
		collection.Name,
		appevdoc.
			FromEvents(
				lab4.FakeEvents(
					1000,
					[]user.ID{"123", "234", "356"},
				),
			).
			KVs()...,
	)
	if err != nil {
		return errs.WrapJoin(
			"failed to insert events",
			err,
			errs.ErrExternalDependencyFailed,
		)
	}
	shards, err := s.mongo.Shards(
		ctx,
		collection.Name,
	)
	if err != nil {
		return errs.WrapJoin(
			"failed to get shards",
			err,
			errs.ErrExternalDependencyFailed,
		)
	}
	err = numbers.AssertEqualOrGreater(
		lab4.MinShardCount,
		len(shards),
	)
	if err != nil {
		return errs.Wrap(
			err,
			"shards",
		)
	}
	for _, shard := range shards {
		id := shard.ID()
		if shard.Count() == 0 {
			return errs.Wrap(
				errs.ErrExpectationFailed,
				"expected shard '%s' to have more than 0 records",
				id,
			)
		}
		if !shard.Ok() {
			return errs.Wrap(
				errs.ErrExpectationFailed,
				"shard '%s' is not ok",
				id,
			)
		}
		hosts, err := s.mongo.HostsOfShard(ctx, id)
		if err != nil {
			return errs.WrapJoin(
				"failed to get hosts for shard",
				err,
				errs.ErrExternalDependencyFailed,
			)
		}
		err = numbers.AssertEqualOrGreater(
			lab4.MinReplicasCount,
			len(hosts),
		)
		if err != nil {
			return errs.Wrap(
				err,
				"count of replica for '%s'",
				id,
			)
		}
	}
	return nil
}
