package reactions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/redis/event/reactions"
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
				val: "event:123:reactions",
			},
		},
		{
			name: "empty id",
			args: args{
				sfx: "",
			},
			want: want{
				val: "event::reactions",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				reactions.Key(
					tt.args.sfx,
				),
			)
		})
	}
}
