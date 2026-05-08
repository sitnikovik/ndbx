package tag_test

import (
	"testing"

	impl "github.com/sitnikovik/ndbx/autograder/internal/app/event/tag"
	"github.com/stretchr/testify/assert"
)

func TestTag_String(t *testing.T) {
	t.Parallel()
	type want struct {
		value string
	}
	tests := []struct {
		name string
		t    impl.Tag
		want want
	}{
		{
			name: "concert",
			t:    impl.Concert,
			want: want{
				value: "concert",
			},
		},
		{
			name: "exhibition",
			t:    impl.Exhibition,
			want: want{
				value: "exhibition",
			},
		},
		{
			name: "culture",
			t:    impl.Culture,
			want: want{
				value: "culture",
			},
		},
		{
			name: "theater",
			t:    impl.Theater,
			want: want{
				value: "theater",
			},
		},
		{
			name: "sport",
			t:    impl.Sport,
			want: want{
				value: "sport",
			},
		},
		{
			name: "food",
			t:    impl.Food,
			want: want{
				value: "food",
			},
		},
		{
			name: "education",
			t:    impl.Education,
			want: want{
				value: "education",
			},
		},
		{
			name: "technology",
			t:    impl.Technology,
			want: want{
				value: "technology",
			},
		},
		{
			name: "custom",
			t:    impl.Tag("foo"),
			want: want{
				value: "foo",
			},
		},
		{
			name: "custom string concert",
			t:    impl.Tag("concert"),
			want: want{
				value: "concert",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.t.String(),
			)
		})
	}
}
