package autograder_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/autograder"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	stepfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/step"
)

func TestAutograder_Run(t *testing.T) {
	t.Parallel()
	type fields struct {
		jobs []autograder.Runner
	}
	type args struct {
		ctx  context.Context
		vars step.Variables
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "first step of three ones failed",
			fields: fields{
				jobs: []autograder.Runner{
					stepfk.NewFakeRunner(
						stepfk.WithErrRun(assert.AnError),
					),
					stepfk.NewFakeRunner(
						stepfk.WithErrRun(assert.AnError),
					),
					stepfk.NewFakeRunner(
						stepfk.WithOkRun(),
					),
				},
			},
			wantErr: assert.AnError,
		},
		{
			name: "all passed",
			fields: fields{
				jobs: []autograder.Runner{
					stepfk.NewFakeRunner(
						stepfk.WithOkRun(),
					),
					stepfk.NewFakeRunner(
						stepfk.WithOkRun(),
					),
					stepfk.NewFakeRunner(
						stepfk.WithOkRun(),
					),
				},
			},
			wantErr: nil,
		},
		{
			name: "empty list",
			fields: fields{
				jobs: []autograder.Runner{},
			},
			wantErr: errs.ErrNothingToRun,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.wantErr,
				autograder.
					NewAutograder(tt.fields.jobs...).
					Run(
						tt.args.ctx,
						tt.args.vars,
					),
			)
		})
	}
}
