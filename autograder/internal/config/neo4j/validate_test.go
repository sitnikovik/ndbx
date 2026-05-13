package neo4j_test

import (
	"testing"

	impl "github.com/sitnikovik/ndbx/autograder/internal/config/neo4j"
	"github.com/stretchr/testify/assert"
)

func TestConfig_Validate(t *testing.T) {
	t.Parallel()
	type want struct {
		errContains string
		errored     bool
	}
	tests := []struct {
		name string
		c    impl.Config
		want want
	}{
		{
			name: "ok",
			c: impl.NewConfig(
				impl.NewConnection("neo4j://localhost:7687"),
				impl.NewAuth("user", "password"),
			),
			want: want{
				errContains: "",
				errored:     false,
			},
		},
		{
			name: "invalid url",
			c: impl.NewConfig(
				impl.NewConnection("invalid"),
				impl.NewAuth("user", "password"),
			),
			want: want{
				errContains: "",
				errored:     false,
			},
		},
		{
			name: "empty connection",
			c: impl.NewConfig(
				impl.NewConnection(""),
				impl.NewAuth("user", "password"),
			),
			want: want{
				errContains: "connection",
				errored:     true,
			},
		},
		{
			name: "empty auth",
			c: impl.NewConfig(
				impl.NewConnection("neo4j://localhost:7687"),
				impl.NewAuth("", ""),
			),
			want: want{
				errContains: "auth",
				errored:     true,
			},
		},
		{
			name: "empty username",
			c: impl.NewConfig(
				impl.NewConnection("neo4j://localhost:7687"),
				impl.NewAuth("", "password"),
			),
			want: want{
				errContains: "auth",
				errored:     true,
			},
		},
		{
			name: "empty password",
			c: impl.NewConfig(
				impl.NewConnection("neo4j://localhost:7687"),
				impl.NewAuth("user", ""),
			),
			want: want{
				errContains: "auth",
				errored:     true,
			},
		},
		{
			name: "all empty",
			c: impl.NewConfig(
				impl.NewConnection(""),
				impl.NewAuth("", ""),
			),
			want: want{
				errContains: "connection",
				errored:     true,
			},
		},
		{
			name: "default value",
			c:    impl.Config{},
			want: want{
				errContains: "connection",
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
					err)
			}
		})
	}
}
