package user_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/config/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/config/app/user/session"
)

func TestConfig_Session(t *testing.T) {
	t.Parallel()
	type fields struct {
		session session.Config
	}
	tests := []struct {
		name   string
		fields fields
		want   session.Config
	}{
		{
			name: "user session config",
			fields: fields{
				session: session.NewConfig(
					5 * time.Second,
				),
			},
			want: session.NewConfig(
				5 * time.Second,
			),
		},
		{
			name:   "default fields",
			fields: fields{},
			want:   session.Config{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := user.NewConfig(
				tt.fields.session,
			)
			got := c.Session()
			assert.Equal(t, tt.want, got)
		})
	}
}
