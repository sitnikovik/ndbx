package endpoint_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	rq "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/get/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/step/user/one/events/endpoint"
	httpxfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/httpx"
	eventfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/event"
	userfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/user"
)

func TestStep_Name(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		step      *endpoint.Step
		want      string
		wantPanic bool
	}{
		{
			name: "ok",
			step: endpoint.NewStep(
				httpxfk.NewFakeClient(),
				"/localhost",
				userfx.NewAlexSmith(),
				rq.NewBody(),
				[]event.Event{
					eventfx.NewTestEvent(),
				},
			),
			want: endpoint.Name,
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
		step      *endpoint.Step
		want      string
		wantPanic bool
	}{
		{
			name: "ok",
			step: endpoint.NewStep(
				httpxfk.NewFakeClient(),
				"/localhost",
				userfx.NewAlexSmith(),
				rq.NewBody(),
				[]event.Event{
					eventfx.NewTestEvent(),
				},
			),
			want: endpoint.Description,
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
