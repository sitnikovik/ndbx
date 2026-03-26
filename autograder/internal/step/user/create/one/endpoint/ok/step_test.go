package ok_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	endpoint "github.com/sitnikovik/ndbx/autograder/internal/step/user/create/one/endpoint/ok"
	httpxfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/httpx"
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
				userfx.NewSamSepiol(),
			),
			want: endpoint.Name,
		},
		{
			name: "default fields",
			step: endpoint.NewStep(
				nil,
				"",
				userfx.NewSamSepiol(),
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
				userfx.NewSamSepiol(),
			),
			want: endpoint.Description,
		},
		{
			name: "default fields",
			step: endpoint.NewStep(
				nil,
				"",
				userfx.NewSamSepiol(),
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
