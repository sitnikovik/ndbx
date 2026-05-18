package recommendation_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	impl "github.com/sitnikovik/ndbx/autograder/internal/config/app/user/recommendation"
	"github.com/sitnikovik/ndbx/autograder/internal/config/app/user/recommendation/event"
)

func TestConfig_Events(t *testing.T) {
	t.Parallel()
	type want struct {
		cfg event.Config
	}
	tests := []struct {
		name string
		c    impl.Config
		want want
	}{
		{
			name: "ok",
			c: impl.NewConfig(
				impl.WithEvent(
					event.NewConfig(
						1 * time.Minute,
					),
				),
			),
			want: want{
				cfg: event.NewConfig(
					1 * time.Minute,
				),
			},
		},
		{
			name: "default value",
			c:    impl.Config{},
			want: want{
				cfg: event.NewConfig(0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.cfg,
				tt.c.Events(),
			)
		})
	}
}
