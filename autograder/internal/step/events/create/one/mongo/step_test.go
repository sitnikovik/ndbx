package mongo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/sitnikovik/ndbx/autograder/internal/step/events/create/one/mongo"
	mongofk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/mongo"
	eventfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/event"
)

func TestStep_Name(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		step      *impl.Step
		want      string
		wantPanic bool
	}{
		{
			name: "ok",
			step: impl.NewStep(
				mongofk.NewFakeClient(),
				eventfx.NewTestEvent(),
			),
			want: impl.Name,
		},
		{
			name: "default fields",
			step: &impl.Step{},
			want: impl.Name,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want,
				tt.step.Name(),
			)
		})
	}
}

func TestStep_Description(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		step      *impl.Step
		want      string
		wantPanic bool
	}{
		{
			name: "ok",
			step: impl.NewStep(
				mongofk.NewFakeClient(),
				eventfx.NewTestEvent(),
			),
			want: impl.Description,
		},
		{
			name: "default fields",
			step: &impl.Step{},
			want: impl.Description,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want,
				tt.step.Description(),
			)
		})
	}
}
