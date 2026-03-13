package endpoint_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/events/create/exists/endpoint"
	httpxfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/httpx"
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
			),
			want: endpoint.Name,
		},
		{
			name: "default fields",
			step: endpoint.NewStep(
				nil,
				"",
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
			),
			want: endpoint.Description,
		},
		{
			name: "default fields",
			step: endpoint.NewStep(
				nil,
				"",
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
