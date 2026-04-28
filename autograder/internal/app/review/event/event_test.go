package event_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	impl "github.com/sitnikovik/ndbx/autograder/internal/app/review/event"
)

func TestEvent_ID(t *testing.T) {
	t.Parallel()
	type want struct {
		value event.ID
	}
	tests := []struct {
		name string
		e    impl.Event
		want want
	}{
		{
			name: "ok",
			e: impl.NewEvent(
				event.NewID("123rews"),
			),
			want: want{
				value: event.NewID("123rews"),
			},
		},
		{
			name: "default value",
			e:    impl.Event{},
			want: want{
				value: event.NewID(""),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.e.ID(),
			)
		})
	}
}
