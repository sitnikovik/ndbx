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

func TestCosts_Equals(t *testing.T) {
	t.Parallel()
	type args struct {
		other event.Costs
	}
	type want struct {
		value bool
	}
	tests := []struct {
		name string
		c    event.Costs
		args args
		want want
	}{
		{
			name: "same entry",
			c: event.NewCosts(
				money.NewMoney(100, 50),
			),
			args: args{
				other: event.NewCosts(
					money.NewMoney(100, 50),
				),
			},
			want: want{
				value: true,
			},
		},
		{
			name: "diff entry",
			c: event.NewCosts(
				money.NewMoney(100, 50),
			),
			args: args{
				other: event.NewCosts(
					money.NewMoney(100, 00),
				),
			},
			want: want{
				value: false,
			},
		},
		{
			name: "default value",
			c:    event.Costs{},
			args: args{
				other: event.NewCosts(
					money.NewMoney(100, 00),
				),
			},
			want: want{
				value: false,
			},
		},
		{
			name: "default arg",
			c: event.NewCosts(
				money.NewMoney(100, 00),
			),
			args: args{
				other: event.Costs{},
			},
			want: want{
				value: false,
			},
		},
		{
			name: "default value",
			c:    event.Costs{},
			args: args{
				other: event.NewCosts(
					money.NewMoney(100, 00),
				),
			},
			want: want{
				value: false,
			},
		},
		{
			name: "default arg and value",
			c:    event.Costs{},
			args: args{
				other: event.Costs{},
			},
			want: want{
				value: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.c.Equals(tt.args.other)
			if tt.want.value {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}
