package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

func TestID_Empty(t *testing.T) {
	t.Parallel()
	type want struct {
		ok bool
	}
	tests := []struct {
		name string
		id   user.ID
		want want
	}{
		{
			name: "not empty",
			id:   user.NewID("2134"),
			want: want{
				ok: false,
			},
		},
		{
			name: "zero int as string",
			id:   user.NewID("0"),
			want: want{
				ok: false,
			},
		},
		{
			name: "space",
			id:   user.NewID(" "),
			want: want{
				ok: false,
			},
		},
		{
			name: "empty string",
			id:   user.NewID(""),
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
