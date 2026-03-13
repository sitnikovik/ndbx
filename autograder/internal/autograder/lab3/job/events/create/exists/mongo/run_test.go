package mongo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	mongoStep "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/events/create/exists/mongo"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
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
		vars  step.Variables
		err   error
		panic bool
	}
	tests := []struct {
		name string
		s    *mongoStep.Step
		args args
		want want
	}{
		{
			name: "ok",
			s: mongoStep.NewStep(
				mongofk.NewFakeClient(
					mongofk.WithAllBy(
						func(
							_ context.Context,
							_ string,
							_ doc.KVs,
						) (doc.Documents, error) {
							return doc.NewDocuments(
								doc.NewDocument(
									"000000000000000000000001",
									NewOkEventDocKVFixture()...,
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
				err:   nil,
				panic: false,
			},
		},
		{
			name: "failed to get events",
			s: mongoStep.NewStep(
				mongofk.NewFakeClient(
					mongofk.WithAllBy(
						func(
							_ context.Context,
							_ string,
							_ doc.KVs,
						) (doc.Documents, error) {
							return doc.NewDocuments(), assert.AnError

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
			name: "got gt 1 events",
			s: mongoStep.NewStep(
				mongofk.NewFakeClient(
					mongofk.WithAllBy(
						func(
							_ context.Context,
							_ string,
							_ doc.KVs,
						) (doc.Documents, error) {
							return doc.NewDocuments(
								doc.NewDocument(
									"000000000000000000000001",
									NewOkEventDocKVFixture()...,
								),
								doc.NewDocument(
									"000000000000000000000002",
									NewOkEventDocKVFixture()...,
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
			name: "got no events",
			s: mongoStep.NewStep(
				mongofk.NewFakeClient(
					mongofk.WithAllBy(
						func(
							_ context.Context,
							_ string,
							_ doc.KVs,
						) (doc.Documents, error) {
							return doc.NewDocuments(), nil
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
