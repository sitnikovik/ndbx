package mongo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab4/job/events/create/ok/mongo"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/shard"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	mongofk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/mongo"
)

func TestStep_Run(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx  context.Context
		vars step.Variables
	}
	type want struct {
		err   error
		vars  step.Variables
		panic bool
	}
	tests := []struct {
		name string
		s    *mongo.Step
		args args
		want want
	}{
		{
			name: "ok",
			s: mongo.NewStep(
				mongofk.NewFakeClient(
					mongofk.WithInsert(
						func(
							_ context.Context,
							_ string,
							_ ...doc.KVs,
						) error {
							return nil
						},
					),
					mongofk.WithShards(
						func(
							_ context.Context,
							_ string,
						) (shard.Shards, error) {
							return shard.NewShards(
								shard.NewShard(
									"rs0",
									shard.WithCount(512),
									shard.WithOk(true),
								),
								shard.NewShard(
									"rs1",
									shard.WithCount(488),
									shard.WithOk(true),
								),
							), nil
						},
					),
					mongofk.WithHostsOfShard(
						func(
							_ context.Context,
							_ string,
						) ([]string, error) {
							return []string{
								"mongo:27017",
								"mongo:27018",
								"mongo:27019",
							}, nil
						},
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				vars:  step.NewVariables(),
				err:   nil,
				panic: false,
			},
		},
		{
			name: "failed to insert events",
			s: mongo.NewStep(
				mongofk.NewFakeClient(
					mongofk.WithInsert(
						func(
							_ context.Context,
							_ string,
							_ ...doc.KVs,
						) error {
							return assert.AnError
						},
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				vars:  step.NewVariables(),
				err:   errs.ErrExternalDependencyFailed,
				panic: false,
			},
		},
		{
			name: "failed to get shards",
			s: mongo.NewStep(
				mongofk.NewFakeClient(
					mongofk.WithInsert(
						func(
							_ context.Context,
							_ string,
							_ ...doc.KVs,
						) error {
							return nil
						},
					),
					mongofk.WithShards(
						func(
							_ context.Context,
							_ string,
						) (shard.Shards, error) {
							return nil, assert.AnError
						},
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				vars:  step.NewVariables(),
				err:   errs.ErrExternalDependencyFailed,
				panic: false,
			},
		},
		{
			name: "got less shards than expected",
			s: mongo.NewStep(
				mongofk.NewFakeClient(
					mongofk.WithInsert(
						func(
							_ context.Context,
							_ string,
							_ ...doc.KVs,
						) error {
							return nil
						},
					),
					mongofk.WithShards(
						func(
							_ context.Context,
							_ string,
						) (shard.Shards, error) {
							return shard.NewShards(
								shard.NewShard(
									"rs0",
									shard.WithOk(true),
								),
							), nil
						},
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				vars:  step.NewVariables(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "one of shards has no records",
			s: mongo.NewStep(
				mongofk.NewFakeClient(
					mongofk.WithInsert(
						func(
							_ context.Context,
							_ string,
							_ ...doc.KVs,
						) error {
							return nil
						},
					),
					mongofk.WithShards(
						func(
							_ context.Context,
							_ string,
						) (shard.Shards, error) {
							return shard.NewShards(
								shard.NewShard(
									"rs0",
									shard.WithCount(1000),
									shard.WithOk(true),
								),
								shard.NewShard(
									"rs1",
									shard.WithCount(0),
									shard.WithOk(true),
								),
							), nil
						},
					),
					mongofk.WithHostsOfShard(
						func(
							_ context.Context,
							_ string,
						) ([]string, error) {
							return []string{
								"mongo:27017",
								"mongo:27018",
								"mongo:27019",
							}, nil
						},
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				vars:  step.NewVariables(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "one of shards is not ok",
			s: mongo.NewStep(
				mongofk.NewFakeClient(
					mongofk.WithInsert(
						func(
							_ context.Context,
							_ string,
							_ ...doc.KVs,
						) error {
							return nil
						},
					),
					mongofk.WithShards(
						func(
							_ context.Context,
							_ string,
						) (shard.Shards, error) {
							return shard.NewShards(
								shard.NewShard(
									"rs0",
									shard.WithCount(512),
									shard.WithOk(true),
								),
								shard.NewShard(
									"rs1",
									shard.WithCount(488),
									shard.WithOk(false),
								),
							), nil
						},
					),
					mongofk.WithHostsOfShard(
						func(
							_ context.Context,
							_ string,
						) ([]string, error) {
							return []string{
								"mongo:27017",
								"mongo:27018",
								"mongo:27019",
							}, nil
						},
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				vars:  step.NewVariables(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "failed to get hosts for shard",
			s: mongo.NewStep(
				mongofk.NewFakeClient(
					mongofk.WithInsert(
						func(
							_ context.Context,
							_ string,
							_ ...doc.KVs,
						) error {
							return nil
						},
					),
					mongofk.WithShards(
						func(
							_ context.Context,
							_ string,
						) (shard.Shards, error) {
							return shard.NewShards(
								shard.NewShard(
									"rs0",
									shard.WithCount(512),
									shard.WithOk(true),
								),
								shard.NewShard(
									"rs1",
									shard.WithCount(488),
									shard.WithOk(true),
								),
							), nil
						},
					),
					mongofk.WithHostsOfShard(
						func(
							_ context.Context,
							_ string,
						) ([]string, error) {
							return nil, assert.AnError
						},
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				vars:  step.NewVariables(),
				err:   errs.ErrExternalDependencyFailed,
				panic: false,
			},
		},
		{
			name: "unexpected count of replicas for shard",
			s: mongo.NewStep(
				mongofk.NewFakeClient(
					mongofk.WithInsert(
						func(
							_ context.Context,
							_ string,
							_ ...doc.KVs,
						) error {
							return nil
						},
					),
					mongofk.WithShards(
						func(
							_ context.Context,
							_ string,
						) (shard.Shards, error) {
							return shard.NewShards(
								shard.NewShard(
									"rs0",
									shard.WithCount(512),
									shard.WithOk(true),
								),
								shard.NewShard(
									"rs1",
									shard.WithCount(488),
									shard.WithOk(true),
								),
							), nil
						},
					),
					mongofk.WithHostsOfShard(
						func(
							_ context.Context,
							_ string,
						) ([]string, error) {
							return []string{
								"mongo:27017",
								"mongo:27018",
							}, nil
						},
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				vars:  step.NewVariables(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = tt.s.Run(
						tt.args.ctx,
						tt.args.vars,
					)
				})
				return
			}
			assert.ErrorIs(
				t,
				tt.s.Run(
					tt.args.ctx,
					tt.args.vars,
				),
				tt.want.err,
			)
			assert.Equal(
				t,
				tt.want.vars,
				tt.args.vars,
			)
		})
	}
}
