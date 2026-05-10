package endpoint_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response/expectation"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/events/create/one/endpoint"
	httpxfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/httpx"
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
				step.NewDesc(
					"Title",
					"Description",
				),
				httpxfk.NewFakeClient(),
				"/localhost",
				eventfx.NewTestEvent(),
				expectation.NewExpectations(
					expectation.WithAsserts(),
				),
			),
			want: "Title",
		},
		{
			name: "default fields",
			step: impl.NewStep(
				step.NewDesc(
					"",
					"",
				),
				nil,
				"",
				eventfx.NewTestEvent(),
				expectation.NewExpectations(
					expectation.WithAsserts(),
				),
			),
			want: "",
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
				step.NewDesc(
					"Title",
					"Description",
				),
				httpxfk.NewFakeClient(),
				"/localhost",
				eventfx.NewTestEvent(),
				expectation.NewExpectations(
					expectation.WithAsserts(),
				),
			),
			want: "Description",
		},
		{
			name: "default fields",
			step: impl.NewStep(
				step.NewDesc(
					"",
					"",
				),
				nil,
				"",
				eventfx.NewTestEvent(),
				expectation.NewExpectations(
					expectation.WithAsserts(),
				),
			),
			want: "",
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
