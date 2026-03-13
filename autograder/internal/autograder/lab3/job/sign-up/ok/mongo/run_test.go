package mongo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/mongo/user/doc/key"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	mongoStep "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/sign-up/ok/mongo"
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
					mongofk.WithAllBy(
						func(
							_ context.Context,
							_ string,
							_ doc.KVs,
						) (doc.Documents, error) {
							return doc.NewDocuments(
								doc.NewDocument(
									"000000000000000000000000",
									doc.NewKV(key.FullName, "Sam Sepiol"),
									doc.NewKV(key.Username, "sams3piol"),
									doc.NewKV(key.Password, "svpa_dvpa_str0ng_h4sh"),
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
								doc.NewUniqueIndex("_id"),
								doc.NewUniqueIndex(key.Username),
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
							user.NewID(""),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.UserPassword,
						"password123",
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
							user.NewID("000000000000000000000000"),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.UserPassword,
						"password123",
					)
					return vars
				}(),
				err:   nil,
				panic: false,
			},
		},
		{
			name: "failed to fetch users",
			s: mongoStep.NewStep(
				mongofk.NewFakeClient(
					mongofk.WithAllBy(
						func(
							_ context.Context,
							_ string,
							_ doc.KVs,
						) (doc.Documents, error) {
							return nil, assert.AnError
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
							user.NewID(""),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.UserPassword,
						"password123",
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
							user.NewID(""),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.UserPassword,
						"password123",
					)
					return vars
				}(),
				err:   errs.ErrExternalDependencyFailed,
				panic: false,
			},
		},
		{
			name: "got many users",
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
									"000000000000000000000000",
									doc.NewKV(key.FullName, "Sam Sepiol"),
									doc.NewKV(key.Username, "sams3piol"),
									doc.NewKV(key.Password, "svpa_dvpa_str0ng_h4sh"),
								),
								doc.NewDocument(
									"000000000000000000000001",
									doc.NewKV(key.FullName, "Sam Sepiol"),
									doc.NewKV(key.Username, "sams3piol"),
									doc.NewKV(key.Password, "svpa_dvpa_str0ng_h4sh"),
								),
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
							user.NewID(""),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.UserPassword,
						"password123",
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
							user.NewID(""),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.UserPassword,
						"password123",
					)
					return vars
				}(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "password is not a string",
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
									"000000000000000000000000",
									doc.NewKV(key.FullName, "Sam Sepiol"),
									doc.NewKV(key.Username, "sams3piol"),
									doc.NewKV(key.Password, 12345),
								),
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
							user.NewID(""),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.UserPassword,
						"password123",
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
							user.NewID(""),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.UserPassword,
						"password123",
					)
					return vars
				}(),
				err:   nil,
				panic: true,
			},
		},
		{
			name: "password equals to provided password",
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
									"000000000000000000000000",
									doc.NewKV(key.FullName, "Sam Sepiol"),
									doc.NewKV(key.Username, "sams3piol"),
									doc.NewKV(key.Password, "password123"),
								),
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
							user.NewID(""),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.UserPassword,
						"password123",
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
							user.NewID(""),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.UserPassword,
						"password123",
					)
					return vars
				}(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "missed some input vars",
			s:    mongoStep.NewStep(mongofk.NewFakeClient()),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						variable.User,
						user.NewUser(
							user.NewID(""),
							"sams3p1ol",
							"Sam Sepiol",
						),
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
							user.NewID(""),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					return vars
				}(),
				err:   nil,
				panic: true,
			},
		},
		{
			name: "failed to get indexes",
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
									"000000000000000000000000",
									doc.NewKV(key.FullName, "Sam Sepiol"),
									doc.NewKV(key.Username, "sams3piol"),
									doc.NewKV(key.Password, "svpa_dvpa_str0ng_h4sh"),
								),
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
							user.NewID(""),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.UserPassword,
						"password123",
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
							user.NewID("000000000000000000000000"),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.UserPassword,
						"password123",
					)
					return vars
				}(),
				err:   errs.ErrExternalDependencyFailed,
				panic: false,
			},
		},
		{
			name: "has no index for username field",
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
									"000000000000000000000000",
									doc.NewKV(key.FullName, "Sam Sepiol"),
									doc.NewKV(key.Username, "sams3piol"),
									doc.NewKV(key.Password, "svpa_dvpa_str0ng_h4sh"),
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
								doc.NewUniqueIndex("_id"),
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
							user.NewID(""),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.UserPassword,
						"password123",
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
							user.NewID("000000000000000000000000"),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.UserPassword,
						"password123",
					)
					return vars
				}(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "has no unique index for username field",
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
									"000000000000000000000000",
									doc.NewKV(key.FullName, "Sam Sepiol"),
									doc.NewKV(key.Username, "sams3piol"),
									doc.NewKV(key.Password, "svpa_dvpa_str0ng_h4sh"),
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
								doc.NewUniqueIndex("_id"),
								doc.NewIndex(key.Username),
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
							user.NewID(""),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.UserPassword,
						"password123",
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
							user.NewID("000000000000000000000000"),
							"sams3p1ol",
							"Sam Sepiol",
						),
					)
					vars.Set(
						variable.UserPassword,
						"password123",
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
		})
	}
}
