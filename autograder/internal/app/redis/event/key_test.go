package event_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/sitnikovik/ndbx/autograder/internal/app/redis/event"
)

func TestKey(t *testing.T) {
	t.Parallel()
	type args struct {
		sfx string
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
				sfx: "123",
			},
			want: want{
				val: "event:123",
			},
		},
		{
			name: "empty id",
			args: args{
				sfx: "",
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
					tt.args.sfx,
				),
			)
		})
	}
}
