package relationship_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/sitnikovik/ndbx/autograder/internal/app/neo4j/relationship"
)

func TestType_String(t *testing.T) {
	t.Parallel()
	type want struct {
		value string
	}
	tests := []struct {
		name string
		t    impl.Type
		want want
	}{
		{
			name: "liked const",
			t:    impl.Liked,
			want: want{
				value: "LIKED",
			},
		},
		{
			name: "liked by a string",
			t:    impl.NewType("liked"),
			want: want{
				value: "liked",
			},
		},
		{
			name: "empty string",
			t:    impl.NewType(""),
			want: want{
				value: "",
			},
		},
		{
			name: "unknown type",
			t:    impl.NewType("UNKNOWN_TYPE"),
			want: want{
				value: "UNKNOWN_TYPE",
			},
		},
		{
			name: "whitespace",
			t:    impl.NewType(" "),
			want: want{
				value: " ",
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
