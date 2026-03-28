package endpoint_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/users/list/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/user/list/by/endpoint"
	httpxfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/httpx"
	userfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/user"
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
				httpxfk.NewFakeClient(),
				"/localhost",
				body.NewBody(),
				[]user.User{
					userfx.NewSamSepiol(),
				},
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
				httpxfk.NewFakeClient(),
				"/localhost",
				body.NewBody(),
				[]user.User{
					userfx.NewSamSepiol(),
				},
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
