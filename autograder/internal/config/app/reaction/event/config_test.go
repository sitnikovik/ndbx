package event_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	impl "github.com/sitnikovik/ndbx/autograder/internal/config/app/reaction/event"
)

func TestConfig_TTL(t *testing.T) {
	t.Parallel()
	type want struct {
		val time.Duration
	}
	tests := []struct {
		name string
		c    impl.Config
		want want
	}{
		{
			name: "ok",
			c:    impl.NewConfig(1 * time.Second),
			want: want{
				val: 1 * time.Second,
			},
		},
		{
			name: "default value",
			c:    impl.Config{},
			want: want{
				val: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.c.TTL(),
			)
		})
	}
}
