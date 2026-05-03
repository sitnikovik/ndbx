package cassandra_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cassandra"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/review/event/list/cassandra"
	"github.com/sitnikovik/ndbx/autograder/internal/step/review/event/list/cassandra/expectation"
	cassandrafk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/cassandra"
	dbfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/cassandra/client"
	eventfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/event"
	userfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

var (
	// eventFixture is a fixture used in the tests cases.
	eventFixture = eventfx.NewBirthdayParty(
		event.NewDates(
			timex.MustRFC3339("2026-03-31T15:00:00Z"),
			timex.MustRFC3339("2026-03-31T23:00:00Z"),
		),
		timex.MustRFC3339("2026-03-14T12:31:00Z"),
		userfx.NewSamwiseGamgee(),
	)
	// vars is the step variables used in the tests cases.
	vars = func() step.Variables {
		vv := step.NewVariables()
		vv.Set(
			eventFixture.Hash(),
			"13298",
		)
		return vv
	}()
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
		s    *impl.Step
		args args
		want want
	}{
		{
			name: "ok",
			s: impl.NewStep(
				step.NewDesc(
					"test title",
					"test desc",
				),
				dbfk.NewClient(
					dbfk.WithSelect(
						func(
							_ context.Context,
							_ string,
							_ ...any,
						) (cassandra.Scanner, error) {
							return cassandrafk.NewIter(
								cassandrafk.NewRow(
									[]string{
										"id",
										"event_id",
										"rate",
										"comment",
										"created_at",
										"created_by",
										"updated_at",
									},
									[]any{
										"123",
										"43224",
										int8(4),
										"Great!",
										timex.MustRFC3339("2025-03-01T12:00:00Z"),
										"123213213",
										timex.MustRFC3339("2025-03-03T14:00:00Z"),
									},
								),
							), nil
						},
					),
				),
				eventFixture,
				expectation.NewExpectations(
					expectation.WithCount(
						1,
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: vars,
			},
			want: want{
				vars:  vars,
				err:   nil,
				panic: false,
			},
		},
		{
			name: "event id not set in vars",
			s: impl.NewStep(
				step.NewDesc(
					"test title",
					"test desc",
				),
				dbfk.NewClient(
					dbfk.WithSelect(
						func(
							_ context.Context,
							_ string,
							_ ...any,
						) (cassandra.Scanner, error) {
							return cassandrafk.NewIter(
								cassandrafk.NewRow(
									[]string{
										"id",
										"event_id",
										"rate",
										"comment",
										"created_at",
										"created_by",
										"updated_at",
									},
									[]any{
										"123",
										"43224",
										int8(4),
										"Great!",
										timex.MustRFC3339("2025-03-01T12:00:00Z"),
										"123213213",
										timex.MustRFC3339("2025-03-03T14:00:00Z"),
									},
								),
							), nil
						},
					),
				),
				eventFixture,
				expectation.NewExpectations(
					expectation.WithCount(
						1,
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
				panic: true,
			},
		},
		{
			name: "failed to select rows",
			s: impl.NewStep(
				step.NewDesc(
					"test title",
					"test desc",
				),
				dbfk.NewClient(
					dbfk.WithSelect(
						func(
							_ context.Context,
							_ string,
							_ ...any,
						) (cassandra.Scanner, error) {
							return nil, assert.AnError
						},
					),
				),
				eventFixture,
				expectation.NewExpectations(
					expectation.WithCount(1),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: vars,
			},
			want: want{
				vars:  vars,
				err:   errs.ErrExternalDependencyFailed,
				panic: false,
			},
		},
		{
			name: "ok",
			s: impl.NewStep(
				step.NewDesc(
					"test title",
					"test desc",
				),
				dbfk.NewClient(
					dbfk.WithSelect(
						func(
							_ context.Context,
							_ string,
							_ ...any,
						) (cassandra.Scanner, error) {
							return cassandrafk.NewIter(
								cassandrafk.NewRow(
									[]string{
										"id",
										"event_id",
										"rate",
										"comment",
										"created_at",
										"created_by",
										"updated_at",
									},
									[]any{
										"123",
										"43224",
										int8(4),
										"Great!",
										timex.MustRFC3339("2025-03-01T12:00:00Z"),
										"123213213",
										timex.MustRFC3339("2025-03-03T14:00:00Z"),
									},
								),
								cassandrafk.NewRow(
									[]string{
										"id",
										"event_id",
										"rate",
										"comment",
										"created_at",
										"created_by",
										"updated_at",
									},
									[]any{
										"124",
										"43224",
										int8(5),
										"Great!!!",
										timex.MustRFC3339("2025-03-01T13:00:00Z"),
										"123213213",
										timex.MustRFC3339("2025-03-03T15:00:00Z"),
									},
								),
							), nil
						},
					),
				),
				eventFixture,
				expectation.NewExpectations(
					expectation.WithCount(
						1,
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: vars,
			},
			want: want{
				vars:  vars,
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
