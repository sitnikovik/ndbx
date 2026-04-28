package redis_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/step"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/events/one/review/redis"
	redisfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/redis"
	eventfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/event"
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
					"Test title",
					"Test description",
				),
				redisfk.NewFakeClient(),
				eventfx.NewTestEvent(),
			),
			want: want{
				value: "Test title",
				panic: false,
			},
		},
		{
			name: "default fields",
			step: &impl.Step{},
			want: want{
				value: "",
				panic: false,
			},
		},
		{
			name: "nil",
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
					"Test title",
					"Test description",
				),
				redisfk.NewFakeClient(),
				eventfx.NewTestEvent(),
			),
			want: want{
				value: "Test description",
				panic: false,
			},
		},
		{
			name: "default fields",
			step: &impl.Step{},
			want: want{
				value: "",
				panic: false,
			},
		},
		{
			name: "nil",
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
