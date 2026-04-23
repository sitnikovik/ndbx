package endpoint_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/reviews/events/create/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/rating"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/review/event/create/endpoint"
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
				descFixture,
				httpxfk.NewFakeClient(),
				"/localhost",
				eventFixture,
				body.NewBody(
					body.WithComment("test review"),
					body.WithRating(rating.Five),
				),
			),
			want: want{
				value: descFixture.Title(),
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
				descFixture,
				httpxfk.NewFakeClient(),
				"/localhost",
				eventFixture,
				body.NewBody(
					body.WithComment("test review"),
					body.WithRating(rating.Five),
				),
			),
			want: want{
				value: descFixture.Description(),
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
