package mongo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/config/mongo"
)

func TestConfig_URI(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		c    mongo.Config
		want string
	}{
		{
			name: "ok",
			c: mongo.NewConfig(
				"testdb",
				"testuser",
				"testpass",
				"localhost",
				27017,
			),
			want: "mongodb://testuser:testpass@localhost:27017/testdb",
		},
		{
			name: "no auth",
			c:    mongo.NewConfig("testdb", "", "", "localhost", 27017),
			want: "mongodb://localhost:27017/testdb",
		},
		{
			name: "empty config",
			c:    mongo.NewConfig("", "", "", "", 0),
			want: "mongodb://:0/",
		},
		{
			name: "negative port",
			c:    mongo.NewConfig("testdb", "testuser", "testpass", "localhost", -1),
			want: "mongodb://testuser:testpass@localhost:-1/testdb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want,
				tt.c.URI(),
			)
		})
	}
}

func TestConfig_Database(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		c    mongo.Config
		want string
	}{
		{
			name: "ok",
			c: mongo.NewConfig(
				"testdb",
				"testuser",
				"testpass",
				"localhost",
				27017,
			),
			want: "testdb",
		},
		{
			name: "empty config",
			c:    mongo.NewConfig("", "", "", "", 0),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want,
				tt.c.Database(),
			)
		})
	}
}
