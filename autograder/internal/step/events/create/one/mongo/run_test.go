package mongo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/events/create/one/mongo"
	mongofk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/mongo"
	eventfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/event"
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
		s    *impl.Step
		args args
		want want
	}{
		{
			name: "ok",
			s: impl.NewStep(
				mongofk.NewFakeClient(
					mongofk.WithInsert(
						func(
							_ context.Context,
							_ string,
							_ ...doc.KVs,
						) ([]string, error) {
							return []string{"1"}, nil
						},
					),
				),
				eventfx.NewTestEvent(),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						eventfx.
							NewTestEvent().
							Created().
							By().
							Hash(),
						"213edf",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						eventfx.
							NewTestEvent().
							Created().
							By().
							Hash(),
						"213edf",
					)
					vars.Set(
						eventfx.NewTestEvent().Hash(),
						"1",
					)
					return vars
				}(),
				err:   nil,
				panic: false,
			},
		},
		{
			name: "failed to insert",
			s: impl.NewStep(
				mongofk.NewFakeClient(
					mongofk.WithInsert(
						func(
							_ context.Context,
							_ string,
							_ ...doc.KVs,
						) ([]string, error) {
							return nil, assert.AnError
						},
					),
				),
				eventfx.NewTestEvent(),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						eventfx.
							NewTestEvent().
							Created().
							By().
							Hash(),
						"213edf",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						eventfx.
							NewTestEvent().
							Created().
							By().
							Hash(),
						"213edf",
					)
					return vars
				}(),
				err:   errs.ErrExternalDependencyFailed,
				panic: false,
			},
		},
		{
			name: "got empty ids",
			s: impl.NewStep(
				mongofk.NewFakeClient(
					mongofk.WithInsert(
						func(
							_ context.Context,
							_ string,
							_ ...doc.KVs,
						) ([]string, error) {
							return []string{}, nil
						},
					),
				),
				eventfx.NewTestEvent(),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						eventfx.
							NewTestEvent().
							Created().
							By().
							Hash(),
						"213edf",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						eventfx.
							NewTestEvent().
							Created().
							By().
							Hash(),
						"213edf",
					)
					return vars
				}(),
				err:   errs.ErrExternalDependencyFailed,
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
				assert.Equal(
					t,
					tt.want.vars,
					tt.args.vars,
				)
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
