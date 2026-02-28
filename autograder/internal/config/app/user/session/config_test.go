package session_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/config/app/user/session"
)

func TestConfig_TTL(t *testing.T) {
	t.Parallel()
	type fields struct {
		ttl time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Duration
	}{
		{
			name: "1 minute TTL",
			fields: fields{
				ttl: time.Minute,
			},
			want: time.Minute,
		},
		{
			name:   "default fields",
			fields: fields{},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := session.NewConfig(
				tt.fields.ttl,
			)
			got := c.TTL()
			assert.Equal(t, tt.want, got)
		})
	}
}
