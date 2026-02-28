package app_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/config/app"
	"github.com/sitnikovik/ndbx/autograder/internal/config/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/config/app/user/session"
)

func TestConfig_User(t *testing.T) {
	t.Parallel()
	type fields struct {
		user user.Config
		host string
		port int
	}
	tests := []struct {
		name   string
		fields fields
		want   user.Config
	}{
		{
			name: "user config",
			fields: fields{
				user: user.NewConfig(
					session.NewConfig(
						1 * time.Second,
					),
				),
			},
			want: user.NewConfig(
				session.NewConfig(
					1 * time.Second,
				),
			),
		},
		{
			name:   "default fields",
			fields: fields{},
			want:   user.Config{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := app.NewConfig(
				tt.fields.user,
				tt.fields.host,
				tt.fields.port,
			)
			got := c.User()
			assert.Equal(t, tt.want, got)
		})
	}
}
