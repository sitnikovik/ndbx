package timex_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestMustParse(t *testing.T) {
	t.Parallel()
	type args struct {
		layout string
		value  string
	}
	tests := []struct {
		name      string
		args      args
		want      time.Time
		wantPanic bool
	}{
		{
			name: "ok rfc3339",
			args: args{
				layout: time.RFC3339,
				value:  "2024-06-23T12:43:00Z",
			},
			want: time.Date(2024, 6, 23, 12, 43, 0, 0, time.UTC),
		},
		{
			name: "ok custom layout",
			args: args{
				layout: "2006-01-02 15:04:05",
				value:  "2024-06-23 12:43:00",
			},
			want: time.Date(2024, 6, 23, 12, 43, 0, 0, time.UTC),
		},
		{
			name: "invalid layout",
			args: args{
				layout: "2024-06-23 12:43:00",
				value:  "2024-06-23T12:43:00Z",
			},
			wantPanic: true,
		},
		{
			name: "empty layout",
			args: args{
				layout: "",
				value:  "2024-06-23T12:43:00Z",
			},
			wantPanic: true,
		},
		{
			name: "invalid value",
			args: args{
				layout: time.RFC3339,
				value:  "23.06.2024 12:43",
			},
			wantPanic: true,
		},
		{
			name: "empty value",
			args: args{
				layout: time.RFC3339,
				value:  "",
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.wantPanic {
				assert.Panics(t, func() {
					_ = timex.MustParse(tt.args.layout, tt.args.value)
				})
				return
			}
			assert.Equal(
				t,
				tt.want,
				timex.MustParse(
					tt.args.layout,
					tt.args.value,
				),
			)
		})
	}
}
