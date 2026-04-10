package event_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	impl "github.com/sitnikovik/ndbx/autograder/internal/config/app/event"
	reaction "github.com/sitnikovik/ndbx/autograder/internal/config/app/reaction/event"
)

func TestConfig_Reactions(t *testing.T) {
	t.Parallel()
	type want struct {
		val reaction.Config
	}
	tests := []struct {
		name string
		c    impl.Config
		want want
	}{
		{
			name: "ok",
			c: impl.NewConfig(
				reaction.NewConfig(
					10 * time.Second,
				),
			),
			want: want{
				val: reaction.NewConfig(
					10 * time.Second,
				),
			},
		},
		{
			name: "default value",
			c:    impl.Config{},
			want: want{
				val: reaction.NewConfig(0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.c.Reactions(),
			)
		})
	}
}
