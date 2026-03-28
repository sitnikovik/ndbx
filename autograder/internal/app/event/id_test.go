package event_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
)

func TestID_Empty(t *testing.T) {
	t.Parallel()
	type want struct {
		ok bool
	}
	tests := []struct {
		name string
		id   event.ID
		want want
	}{
		{
			name: "not empty",
			id:   event.NewID("2134"),
			want: want{
				ok: false,
			},
		},
		{
			name: "zero int as string",
			id:   event.NewID("0"),
			want: want{
				ok: false,
			},
		},
		{
			name: "space",
			id:   event.NewID(" "),
			want: want{
				ok: false,
			},
		},
		{
			name: "empty string",
			id:   event.NewID(""),
			want: want{
				ok: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.id.Empty()
			if tt.want.ok {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}
