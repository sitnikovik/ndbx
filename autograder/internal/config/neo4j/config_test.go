package neo4j_test

import (
	"testing"

	impl "github.com/sitnikovik/ndbx/autograder/internal/config/neo4j"
	"github.com/stretchr/testify/assert"
)

func TestConfig_Connection(t *testing.T) {
	t.Parallel()
	type want struct {
		value impl.Connection
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
				value: impl.NewConnection("neo4j://localhost:7687"),
			},
		},
		{
			name: "empty connection",
			c: impl.NewConfig(
				impl.NewConnection(""),
				impl.NewAuth("user", "password"),
			),
			want: want{
				value: impl.NewConnection(""),
			},
		},
		{
			name: "default value",
			c:    impl.Config{},
			want: want{
				value: impl.NewConnection(""),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.c.Connection(),
			)
		})
	}
}

func TestConfig_Auth(t *testing.T) {
	t.Parallel()
	type want struct {
		value impl.Auth
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
				value: impl.NewAuth("user", "password"),
			},
		},
		{
			name: "empty auth",
			c: impl.NewConfig(
				impl.NewConnection("neo4j://localhost:7687"),
				impl.NewAuth("", ""),
			),
			want: want{
				value: impl.NewAuth("", ""),
			},
		},
		{
			name: "empty username",
			c: impl.NewConfig(
				impl.NewConnection("neo4j://localhost:7687"),
				impl.NewAuth("", "password"),
			),
			want: want{
				value: impl.NewAuth("", "password"),
			},
		},
		{
			name: "empty password",
			c: impl.NewConfig(
				impl.NewConnection("neo4j://localhost:7687"),
				impl.NewAuth("username", ""),
			),
			want: want{
				value: impl.NewAuth("username", ""),
			},
		},
		{
			name: "default value",
			c:    impl.Config{},
			want: want{
				value: impl.NewAuth("", ""),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.c.Auth(),
			)
		})
	}
}
