package event_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/money"
)

func TestCosts_Entry(t *testing.T) {
	t.Parallel()
	type want struct {
		val money.Money
	}
	tests := []struct {
		name string
		c    event.Costs
		want want
	}{
		{
			name: "ok",
			c: event.NewCosts(
				money.NewMoney(100, 50),
			),
			want: want{
				val: money.NewMoney(100, 50),
			},
		},
		{
			name: "zeros",
			c: event.NewCosts(
				money.NewMoney(0, 0),
			),
			want: want{
				val: money.NewMoney(0, 0),
			},
		},
		{
			name: "default value",
			c:    event.Costs{},
			want: want{
				val: money.NewMoney(0, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.c.Entry(),
			)
		})
	}
}
