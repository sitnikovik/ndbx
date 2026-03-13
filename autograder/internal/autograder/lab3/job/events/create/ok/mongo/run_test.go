package mongo_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/mongo/event/doc/key"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3"
	mongoStep "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/events/create/ok/mongo"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
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
					mongofk.WithByID(
						func(
							_ context.Context,
							_ string,
							_ string,
						) (doc.Document, error) {
							return doc.NewDocument(
								"000000000000000000000001",
								NewOkEventDocKVFixture()...,
							), nil
						},
					),
					mongofk.WithIndexes(
						func(
							_ context.Context,
							_ string,
						) (doc.Indexes, error) {
							return doc.NewIndexes(
								doc.NewUniqueIndex("_id"),
								doc.NewUniqueIndex(key.Title),
								doc.NewIndex(key.Title, key.CreatedBy),
								doc.NewIndex(key.CreatedBy),
							), nil
						},
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.User,
						user.NewUser(
							user.NewID("000000000000000000000123"),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.EventID,
						"000000000000000000000001",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.User,
						user.NewUser(
							user.NewID("000000000000000000000123"),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.EventID,
						"000000000000000000000001",
					)
					return vars
				}(),
				err:   nil,
				panic: false,
			},
		},
		{
			name: "failed to get event",
			s: mongoStep.NewStep(
				mongofk.NewFakeClient(
					mongofk.WithByID(
						func(
							_ context.Context,
							_ string,
							_ string,
						) (doc.Document, error) {
							return doc.Document{}, assert.AnError
						},
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.User,
						user.NewUser(
							user.NewID("000000000000000000000123"),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.EventID,
						"000000000000000000000001",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.User,
						user.NewUser(
							user.NewID("000000000000000000000123"),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.EventID,
						"000000000000000000000001",
					)
					return vars
				}(),
				err:   errs.ErrExternalDependencyFailed,
				panic: false,
			},
		},
		{
			name: "author id does not match",
			s: mongoStep.NewStep(
				mongofk.NewFakeClient(
					mongofk.WithByID(
						func(
							_ context.Context,
							_ string,
							_ string,
						) (doc.Document, error) {
							ent := lab3.NewTestEvent()
							return doc.NewDocument(
								"000000000000000000000001",
								doc.NewKV(
									key.Title,
									ent.
										Content().
										Title(),
								),
								doc.NewKV(
									key.Description,
									ent.
										Content().
										Description(),
								),
								doc.NewKV(
									key.Location,
									ent.Location().Address(),
								),
								doc.NewKV(
									key.CreatedBy,
									user.
										NewID("000000000000000000000999").
										String(),
								),
								doc.NewKV(
									key.CreatedAt,
									ent.
										Created().
										At().
										Format(time.RFC3339),
								),
								doc.NewKV(
									key.StartedAt,
									ent.
										Dates().
										StartedAt().
										Format(time.RFC3339),
								),
								doc.NewKV(
									key.StartedAt,
									ent.
										Dates().
										StartedAt().
										Format(time.RFC3339),
								),
								doc.NewKV(
									key.FinishedAt,
									ent.
										Dates().
										FinishedAt().
										Format(time.RFC3339),
								),
							), nil
						},
					),
					mongofk.WithIndexes(
						func(
							_ context.Context,
							_ string,
						) (doc.Indexes, error) {
							return doc.NewIndexes(
								doc.NewUniqueIndex(key.Title),
								doc.NewIndex(key.Title, key.CreatedBy),
								doc.NewIndex(key.CreatedBy),
							), nil
						},
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.User,
						user.NewUser(
							user.NewID("000000000000000000000123"),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.EventID,
						"000000000000000000000001",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.User,
						user.NewUser(
							user.NewID("000000000000000000000123"),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.EventID,
						"000000000000000000000001",
					)
					return vars
				}(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "failed to get indexes",
			s: mongoStep.NewStep(
				mongofk.NewFakeClient(
					mongofk.WithByID(
						func(
							_ context.Context,
							_ string,
							_ string,
						) (doc.Document, error) {
							return doc.NewDocument(
								"000000000000000000000001",
								NewOkEventDocKVFixture()...,
							), nil
						},
					),
					mongofk.WithIndexes(
						func(
							_ context.Context,
							_ string,
						) (doc.Indexes, error) {
							return doc.NewIndexes(), assert.AnError
						},
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.User,
						user.NewUser(
							user.NewID("000000000000000000000123"),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.EventID,
						"000000000000000000000001",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.User,
						user.NewUser(
							user.NewID("000000000000000000000123"),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.EventID,
						"000000000000000000000001",
					)
					return vars
				}(),
				err:   errs.ErrExternalDependencyFailed,
				panic: false,
			},
		},
		{
			name: "has no index",
			s: mongoStep.NewStep(
				mongofk.NewFakeClient(
					mongofk.WithByID(
						func(
							_ context.Context,
							_ string,
							_ string,
						) (doc.Document, error) {
							return doc.NewDocument(
								"000000000000000000000001",
								NewOkEventDocKVFixture()...,
							), nil
						},
					),
					mongofk.WithIndexes(
						func(
							_ context.Context,
							_ string,
						) (doc.Indexes, error) {
							return doc.NewIndexes(), nil
						},
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.User,
						user.NewUser(
							user.NewID("000000000000000000000123"),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.EventID,
						"000000000000000000000001",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.User,
						user.NewUser(
							user.NewID("000000000000000000000123"),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.EventID,
						"000000000000000000000001",
					)
					return vars
				}(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "title idx is not unique",
			s: mongoStep.NewStep(
				mongofk.NewFakeClient(
					mongofk.WithByID(
						func(
							_ context.Context,
							_ string,
							_ string,
						) (doc.Document, error) {
							return doc.NewDocument(
								"000000000000000000000001",
								NewOkEventDocKVFixture()...,
							), nil
						},
					),
					mongofk.WithIndexes(
						func(
							_ context.Context,
							_ string,
						) (doc.Indexes, error) {
							return doc.NewIndexes(
								doc.NewUniqueIndex("_id"),
								doc.NewIndex(key.Title),
								doc.NewIndex(key.Title, key.CreatedBy),
								doc.NewIndex(key.CreatedBy),
							), nil
						},
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.User,
						user.NewUser(
							user.NewID("000000000000000000000123"),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.EventID,
						"000000000000000000000001",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.User,
						user.NewUser(
							user.NewID("000000000000000000000123"),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.EventID,
						"000000000000000000000001",
					)
					return vars
				}(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "has not idx title n created_by",
			s: mongoStep.NewStep(
				mongofk.NewFakeClient(
					mongofk.WithByID(
						func(
							_ context.Context,
							_ string,
							_ string,
						) (doc.Document, error) {
							return doc.NewDocument(
								"000000000000000000000001",
								NewOkEventDocKVFixture()...,
							), nil
						},
					),
					mongofk.WithIndexes(
						func(
							_ context.Context,
							_ string,
						) (doc.Indexes, error) {
							return doc.NewIndexes(
								doc.NewUniqueIndex("_id"),
								doc.NewUniqueIndex(key.Title),
								doc.NewIndex(key.CreatedBy),
							), nil
						},
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.User,
						user.NewUser(
							user.NewID("000000000000000000000123"),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.EventID,
						"000000000000000000000001",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.User,
						user.NewUser(
							user.NewID("000000000000000000000123"),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.EventID,
						"000000000000000000000001",
					)
					return vars
				}(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "has not idx for created_by only",
			s: mongoStep.NewStep(
				mongofk.NewFakeClient(
					mongofk.WithByID(
						func(
							_ context.Context,
							_ string,
							_ string,
						) (doc.Document, error) {
							return doc.NewDocument(
								"000000000000000000000001",
								NewOkEventDocKVFixture()...,
							), nil
						},
					),
					mongofk.WithIndexes(
						func(
							_ context.Context,
							_ string,
						) (doc.Indexes, error) {
							return doc.NewIndexes(
								doc.NewUniqueIndex("_id"),
								doc.NewUniqueIndex(key.Title),
								doc.NewIndex(key.Title, key.CreatedBy),
							), nil
						},
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.User,
						user.NewUser(
							user.NewID("000000000000000000000123"),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.EventID,
						"000000000000000000000001",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.User,
						user.NewUser(
							user.NewID("000000000000000000000123"),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.EventID,
						"000000000000000000000001",
					)
					return vars
				}(),
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
