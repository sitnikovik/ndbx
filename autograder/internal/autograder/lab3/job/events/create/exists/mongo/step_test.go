package mongo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	step "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/events/create/exists/mongo"
	mongofk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/mongo"
)

func TestStep_Name(t *testing.T) {
	t.Parallel()
	type want struct {
		val string
	}
	tests := []struct {
		name string
		step *step.Step
		want want
	}{
		{
			name: "ok",
			step: step.NewStep(
				mongofk.NewFakeClient(),
			),
			want: want{
				val: step.Name,
			},
		},
		{
			name: "default fields",
			step: step.NewStep(
				nil,
			),
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
				tt.step.Name(),
			)
		})
	}
}

func TestStep_Description(t *testing.T) {
	t.Parallel()
	type want struct {
		val string
	}
	tests := []struct {
		name string
		step *step.Step
		want want
	}{
		{
			name: "ok",
			step: step.NewStep(
				mongofk.NewFakeClient(),
			),
			want: want{
				val: step.Description,
			},
		},
		{
			name: "default fields",
			step: step.NewStep(
				nil,
			),
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
				tt.step.Description(),
			)
		})
	}
}
