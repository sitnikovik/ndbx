package neo4j_test

import (
	"testing"

	impl "github.com/sitnikovik/ndbx/autograder/internal/config/neo4j"
	"github.com/stretchr/testify/assert"
)

func TestConnection_URL(t *testing.T) {
	t.Parallel()
	type want struct {
		value string
	}
	tests := []struct {
		name string
		c    impl.Connection
		want want
	}{
		{
			name: "ok",
			c:    impl.NewConnection("neo4j://localhost:7687"),
			want: want{
				value: "neo4j://localhost:7687",
			},
		},
		{
			name: "empty",
			c:    impl.NewConnection(""),
			want: want{
				value: "",
			},
		},
		{
			name: "invalid",
			c:    impl.NewConnection("invalid"),
			want: want{
				value: "invalid",
			},
		},
		{
			name: "default value",
			c:    impl.Connection{},
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
				tt.c.URL(),
			)
		})
	}
}

func TestConnection_Validate(t *testing.T) {
	t.Parallel()
	type want struct {
		errContains string
		errored     bool
	}
	tests := []struct {
		name string
		c    impl.Connection
		want want
	}{
		{
			name: "ok",
			c:    impl.NewConnection("neo4j://localhost:7687"),
			want: want{
				errContains: "",
				errored:     false,
			},
		},
		{
			name: "empty",
			c:    impl.NewConnection(""),
			want: want{
				errContains: "url",
				errored:     true,
			},
		},
		{
			name: "invalid",
			c:    impl.NewConnection("invalid"),
			want: want{
				errContains: "",
				errored:     false,
			},
		},
		{
			name: "default value",
			c:    impl.Connection{},
			want: want{
				errContains: "url",
				errored:     true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.c.Validate()
			if tt.want.errored {
				assert.Error(t, err)
				assert.Contains(
					t,
					err.Error(),
					tt.want.errContains,
				)
			} else {
				assert.NoErrorf(
					t,
					err,
					"unexpected error: %v",
					err,
				)
			}
		})
	}
}
