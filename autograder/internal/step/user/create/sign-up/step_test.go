package endpoint_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/user/create/sign-up"
	httpxfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/httpx"
	mongofk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/mongo"
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
				mongofk.NewFakeClient(),
				"/localhost",
				userfx.NewJohnDoe(),
				"wqe13",
			),
			want: impl.Name,
		},
		{
			name: "default fields",
			step: impl.NewStep(
				nil,
				nil,
				"",
				user.User{},
				"",
			),
			want: impl.Name,
		},
		{
			name: "default value",
			step: nil,
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
				mongofk.NewFakeClient(),
				"/localhost",
				userfx.NewJohnDoe(),
				"wqe13",
			),
			want: impl.Description,
		},
		{
			name: "default fields",
			step: impl.NewStep(
				nil,
				nil,
				"",
				user.User{},
				"",
			),
			want: impl.Description,
		},
		{
			name: "default value",
			step: nil,
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
