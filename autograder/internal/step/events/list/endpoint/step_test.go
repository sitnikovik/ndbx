package endpoint_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	rq "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/get/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/events/list/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/step/events/list/endpoint/expect"
	httpxfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/httpx"
	eventfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/event"
	userfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
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
					"Test Step",
					"Test description",
				),
				httpxfk.NewFakeClient(),
				"/localhost",
				rq.NewBody(),
				expect.NewExpectations(
					expect.WithEvents(
						eventfx.NewBirthdayParty(
							event.NewDates(
								timex.MustRFC3339("2025-02-03T11:00:00Z"),
								timex.MustRFC3339("2025-02-03T23:00:00Z"),
							),
							timex.MustRFC3339("2025-02-01T10:00:00Z"),
							userfx.NewJohnDoe(),
						),
					),
				),
			),
			want: "Test Step",
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
				rq.NewBody(),
				expect.NewExpectations(
					expect.WithEvents(),
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
					"Test Step",
					"Test description",
				),
				httpxfk.NewFakeClient(),
				"/localhost",
				rq.NewBody(),
				expect.NewExpectations(
					expect.WithEvents(
						eventfx.NewBirthdayParty(
							event.NewDates(
								timex.MustRFC3339("2025-02-03T11:00:00Z"),
								timex.MustRFC3339("2025-02-03T23:00:00Z"),
							),
							timex.MustRFC3339("2025-02-01T10:00:00Z"),
							userfx.NewJohnDoe(),
						),
					),
				),
			),
			want: "Test description",
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
				rq.NewBody(),
				expect.NewExpectations(
					expect.WithEvents(),
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
