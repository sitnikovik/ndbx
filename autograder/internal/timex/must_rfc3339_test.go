package timex_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestMustRFC3339(t *testing.T) {
	t.Parallel()
	type args struct {
		tim string
	}
	type want struct {
		val   time.Time
		panic bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "rfc3339",
			args: args{
				tim: "2024-06-23T12:43:00Z",
			},
			want: want{
				val:   time.Date(2024, 6, 23, 12, 43, 0, 0, time.UTC),
				panic: false,
			},
		},
		{
			name: "datetime",
			args: args{
				tim: "2024-06-23 12:43:00",
			},
			want: want{
				val:   time.Time{},
				panic: true,
			},
		},
		{
			name: "empty string",
			args: args{
				tim: "",
			},
			want: want{
				val:   time.Time{},
				panic: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = timex.MustRFC3339(tt.args.tim)
				})
				return
			}
			assert.Equal(
				t,
				tt.want.val,
				timex.MustRFC3339(tt.args.tim),
			)
		})
	}
}
