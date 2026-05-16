package event_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	impl "github.com/sitnikovik/ndbx/autograder/internal/app/neo4j/node/event"
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
				event.NewID("65e9c0b1a2b3c4d5e6f7a8b9"),
				"The Hobbit",
			),
			want: want{
				value: event.NewID("65e9c0b1a2b3c4d5e6f7a8b9"),
			},
		},
		{
			name: "empty id",
			e: impl.NewEvent(
				event.NewID(""),
				"The Hobbit",
			),
			want: want{
				value: event.NewID(""),
			},
		},
		{
			name: "whitespace id",
			e: impl.NewEvent(
				event.NewID("   "),
				"The Hobbit",
			),
			want: want{
				value: event.NewID("   "),
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

func TestEvent_Title(t *testing.T) {
	t.Parallel()
	type want struct {
		value string
	}
	tests := []struct {
		name string
		e    impl.Event
		want want
	}{
		{
			name: "ok",
			e: impl.NewEvent(
				event.NewID("65e9c0b1a2b3c4d5e6f7a8b9"),
				"The Hobbit",
			),
			want: want{
				value: "The Hobbit",
			},
		},
		{
			name: "empty title",
			e: impl.NewEvent(
				event.NewID("65e9c0b1a2b3c4d5e6f7a8b9"),
				"",
			),
			want: want{
				value: "",
			},
		},
		{
			name: "whitespace title",
			e: impl.NewEvent(
				event.NewID("65e9c0b1a2b3c4d5e6f7a8b9"),
				"   ",
			),
			want: want{
				value: "   ",
			},
		},
		{
			name: "default value",
			e:    impl.Event{},
			want: want{
				value: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.e.Title(),
			)
		})
	}
}
