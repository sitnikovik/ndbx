package event_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
)

func TestContent_Title(t *testing.T) {
	t.Parallel()
	type want struct {
		val string
	}
	tests := []struct {
		name string
		c    event.Content
		want want
	}{
		{
			name: "ok",
			c: event.NewContent(
				"Test Event",
				"This is a test event.",
			),
			want: want{
				val: "Test Event",
			},
		},
		{
			name: "empty title",
			c: event.NewContent(
				"",
				"This is a test event.",
			),
			want: want{
				val: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.c.Title(),
			)
		})
	}
}

func TestContent_Description(t *testing.T) {
	t.Parallel()
	type want struct {
		val string
	}
	tests := []struct {
		name string
		c    event.Content
		want want
	}{
		{
			name: "ok",
			c: event.NewContent(
				"Test Event",
				"This is a test event.",
			),
			want: want{
				val: "This is a test event.",
			},
		},
		{
			name: "empty description",
			c: event.NewContent(
				"Test Event",
				"",
			),
			want: want{
				val: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.c.Description(),
			)
		})
	}
}
