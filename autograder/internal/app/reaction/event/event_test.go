package event_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	impl "github.com/sitnikovik/ndbx/autograder/internal/app/reaction/event"
)

func TestEvent_ID(t *testing.T) {
	t.Parallel()
	type want struct {
		val event.ID
	}
	tests := []struct {
		name string
		e    impl.Event
		want want
	}{
		{
			name: "ok",
			e: impl.NewEvent(
				event.NewID("1"),
			),
			want: want{
				val: event.NewID("1"),
			},
		},
		{
			name: "default value",
			e:    impl.Event{},
			want: want{
				val: event.NewID(""),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.e.ID(),
			)
		})
	}
}
