package endpoint_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/step"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/user/one/recommendations/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/step/user/one/recommendations/expect"
	httpxfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/httpx"
	userfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/user"
)

func TestStep_Name(t *testing.T) {
	t.Parallel()
	type want struct {
		value string
		panic bool
	}
	tests := []struct {
		name string
		step *impl.Step
		want want
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
				userfx.NewAlexSmith(),
				expect.NewExpectations(
					expect.WithNoEvents(),
				),
			),
			want: want{
				value: "Title",
				panic: false,
			},
		},
		{
			name: "nil step",
			step: nil,
			want: want{
				panic: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = tt.step.Name()
				})
				return
			}
			assert.Equal(
				t,
				tt.want.value,
				tt.step.Name(),
			)
		})
	}
}

func TestStep_Description(t *testing.T) {
	t.Parallel()
	type want struct {
		value string
		panic bool
	}
	tests := []struct {
		name string
		step *impl.Step
		want want
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
				userfx.NewAlexSmith(),
				expect.NewExpectations(
					expect.WithNoEvents(),
				),
			),
			want: want{
				value: "Description",
				panic: false,
			},
		},
		{
			name: "nil step",
			step: nil,
			want: want{
				panic: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = tt.step.Name()
				})
				return
			}
			assert.Equal(
				t,
				tt.want.value,
				tt.step.Description(),
			)
		})
	}
}
