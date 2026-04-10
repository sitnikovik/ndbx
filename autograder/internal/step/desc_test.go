package step_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/sitnikovik/ndbx/autograder/internal/step"
)

func TestDesc_Title(t *testing.T) {
	t.Parallel()
	type want struct {
		value string
	}
	tests := []struct {
		name string
		d    impl.Desc
		want want
	}{
		{
			name: "all set",
			d: impl.NewDesc(
				"title",
				"description",
			),
			want: want{
				value: "title",
			},
		},
		{
			name: "only title",
			d: impl.NewDesc(
				"title",
				"",
			),
			want: want{
				value: "title",
			},
		},
		{
			name: "only description",
			d: impl.NewDesc(
				"",
				"description",
			),
			want: want{
				value: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.d.Title(),
			)
		})
	}
}

func TestDesc_Description(t *testing.T) {
	t.Parallel()
	type want struct {
		value string
	}
	tests := []struct {
		name string
		d    impl.Desc
		want want
	}{
		{
			name: "all set",
			d: impl.NewDesc(
				"title",
				"description",
			),
			want: want{
				value: "description",
			},
		},
		{
			name: "only title",
			d: impl.NewDesc(
				"title",
				"",
			),
			want: want{
				value: "",
			},
		},
		{
			name: "only description",
			d: impl.NewDesc(
				"",
				"description",
			),
			want: want{
				value: "description",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.d.Description(),
			)
		})
	}
}
