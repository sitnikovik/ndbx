package money_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/money"
)

func TestMoney_String(t *testing.T) {
	t.Parallel()
	type want struct {
		val string
	}
	tests := []struct {
		name string
		m    money.Money
		want want
	}{
		{
			name: "100.50",
			m:    money.NewMoney(100, 50),
			want: want{
				val: "100.50",
			},
		},
		{
			name: "zeros",
			m:    money.NewMoney(0, 0),
			want: want{
				val: "0.0",
			},
		},
		{
			name: "0.50",
			m:    money.NewMoney(0, 50),
			want: want{
				val: "0.50",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.m.String(),
			)
		})
	}
}

func TestMoney_Free(t *testing.T) {
	t.Parallel()
	type want struct {
		ok bool
	}
	tests := []struct {
		name string
		m    money.Money
		want want
	}{
		{
			name: "100.50",
			m:    money.NewMoney(100, 50),
			want: want{
				ok: false,
			},
		},
		{
			name: "100.00",
			m:    money.NewMoney(100, 0),
			want: want{
				ok: false,
			},
		},
		{
			name: "zeros",
			m:    money.NewMoney(0, 0),
			want: want{
				ok: true,
			},
		},
		{
			name: "0.50",
			m:    money.NewMoney(0, 50),
			want: want{
				ok: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.m.Free()
			if tt.want.ok {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}
