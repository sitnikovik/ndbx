package step_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	stepfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/step"
)

func TestList_Run(t *testing.T) {
	t.Parallel()
	type fields struct {
		steps []step.Runner
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
				steps: []step.Runner{
					stepfk.NewFakeRunner(
						stepfk.WithErrRun(assert.AnError),
					),
					stepfk.NewFakeRunner(
						stepfk.WithOkRun(),
					),
					stepfk.NewFakeRunner(
						stepfk.WithOkRun(),
					),
				},
			},
			wantErr: assert.AnError,
		},
		{
			name: "second step of three ones failed",
			fields: fields{
				steps: []step.Runner{
					stepfk.NewFakeRunner(
						stepfk.WithOkRun(),
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
			name: "all steps passed",
			fields: fields{
				steps: []step.Runner{
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
				steps: []step.Runner{},
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
				step.
					NewList(tt.fields.steps...).
					Run(
						tt.args.ctx,
						tt.args.vars,
					),
			)
		})
	}
}

func TestList_Name(t *testing.T) {
	t.Parallel()
	type want struct {
		val string
	}
	tests := []struct {
		name string
		l    *step.List
		want want
	}{
		{
			name: "empty list",
			l:    step.NewList(),
			want: want{
				val: step.Name,
			},
		},
		{
			name: "has steps",
			l: step.NewList(
				stepfk.NewFakeRunner(
					stepfk.WithErrRun(assert.AnError),
				),
				stepfk.NewFakeRunner(
					stepfk.WithOkRun(),
				),
				stepfk.NewFakeRunner(
					stepfk.WithOkRun(),
				),
			),
			want: want{
				val: step.Name,
			},
		},
		{
			name: "default value",
			l:    nil,
			want: want{
				val: step.Name,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.l.Name(),
			)
		})
	}
}

func TestList_Description(t *testing.T) {
	t.Parallel()
	type want struct {
		val string
	}
	tests := []struct {
		name string
		l    *step.List
		want want
	}{
		{
			name: "empty list",
			l:    step.NewList(),
			want: want{
				val: step.Description,
			},
		},
		{
			name: "has steps",
			l: step.NewList(
				stepfk.NewFakeRunner(
					stepfk.WithErrRun(assert.AnError),
				),
				stepfk.NewFakeRunner(
					stepfk.WithOkRun(),
				),
				stepfk.NewFakeRunner(
					stepfk.WithOkRun(),
				),
			),
			want: want{
				val: step.Description,
			},
		},
		{
			name: "default value",
			l:    nil,
			want: want{
				val: step.Description,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.l.Description(),
			)
		})
	}
}
