package event_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	impl "github.com/sitnikovik/ndbx/autograder/internal/app/redis/event"
)

func TestKey(t *testing.T) {
	t.Parallel()
	type args struct {
		id event.ID
	}
	type want struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "ok",
			args: args{
				id: event.NewID("123"),
			},
			want: want{
				val: "event:123",
			},
		},
		{
			name: "empty id",
			args: args{
				id: event.NewID(""),
			},
			want: want{
				val: "event:",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				impl.Key(
					tt.args.id,
				),
			)
		})
	}
}
