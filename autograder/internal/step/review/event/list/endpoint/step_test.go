package endpoint_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/step"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/review/event/list/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/step/review/event/list/endpoint/expect"
	httpxfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/httpx"
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
					"Test Step",
					"Test description",
				),
				httpxfk.NewFakeClient(),
				eventFixture,
				"/localhost",
				expect.NewExpectations(
					expect.WithCount(0),
				),
			),
			want: want{
				value: "Test Step",
				panic: false,
			},
		},
		{
			name: "empty title",
			step: impl.NewStep(
				step.NewDesc(
					"",
					"Test description",
				),
				httpxfk.NewFakeClient(),
				eventFixture,
				"/localhost",
				expect.NewExpectations(
					expect.WithCount(0),
				),
			),
			want: want{
				value: "",
				panic: false,
			},
		},
		{
			name: "default value",
			step: nil,
			want: want{
				value: "",
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
					"Test Step",
					"Test description",
				),
				httpxfk.NewFakeClient(),
				eventFixture,
				"/localhost",
				expect.NewExpectations(
					expect.WithCount(0),
				),
			),
			want: want{
				value: "Test description",
				panic: false,
			},
		},
		{
			name: "empty description",
			step: impl.NewStep(
				step.NewDesc(
					"Test Step",
					"",
				),
				httpxfk.NewFakeClient(),
				eventFixture,
				"/localhost",
				expect.NewExpectations(
					expect.WithCount(0),
				),
			),
			want: want{
				value: "",
				panic: false,
			},
		},
		{
			name: "default value",
			step: nil,
			want: want{
				value: "",
				panic: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = tt.step.Description()
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
