package event_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
)

func TestQuantity_Max(t *testing.T) {
	t.Parallel()
	type want struct {
		val int
	}
	tests := []struct {
		name string
		q    event.Quantity
		want want
	}{
		{
			name: "ok",
			q:    event.NewQuantity(5, 10),
			want: want{
				val: 10,
			},
		},
		{
			name: "zero max",
			q:    event.NewQuantity(5, 0),
			want: want{
				val: 0,
			},
		},
		{
			name: "max lt min",
			q:    event.NewQuantity(10, 5),
			want: want{
				val: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.q.Max(),
			)
		})
	}
}
func TestQuantity_Min(t *testing.T) {
	t.Parallel()
	type want struct {
		val int
	}
	tests := []struct {
		name string
		q    event.Quantity
		want want
	}{
		{
			name: "ok",
			q:    event.NewQuantity(5, 10),
			want: want{
				val: 5,
			},
		},
		{
			name: "zero min",
			q:    event.NewQuantity(0, 10),
			want: want{
				val: 0,
			},
		},
		{
			name: "max gt min",
			q:    event.NewQuantity(5, 10),
			want: want{
				val: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.q.Min(),
			)
		})
	}
}
