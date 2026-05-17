package user_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/config/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/config/app/user/recommendation"
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

func TestConfig_Recommendations(t *testing.T) {
	t.Parallel()
	type want struct {
		cfg recommendation.Config
	}
	tests := []struct {
		name string
		c    user.Config
		want want
	}{
		{
			name: "ok",
			c: user.NewConfig(
				session.NewConfig(5*time.Second),
				user.WithRecommendations(
					recommendation.NewConfig(),
				),
			),
			want: want{
				cfg: recommendation.NewConfig(),
			},
		},
		{
			name: "default value",
			c:    user.Config{},
			want: want{
				cfg: recommendation.NewConfig(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.cfg,
				tt.c.Recommendations(),
			)
		})
	}
}
