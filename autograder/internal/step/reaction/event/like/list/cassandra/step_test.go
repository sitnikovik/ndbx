package cassandra_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/sitnikovik/ndbx/autograder/internal/step/reaction/event/like/list/cassandra"
	cassandrafk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/cassandra/client"
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
				cassandrafk.NewClient(),
				eventfx.NewTestEvent(),
				1,
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
				cassandrafk.NewClient(),
				eventfx.NewTestEvent(),
				1,
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
