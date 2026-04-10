package env_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/env"
)

func TestGet(t *testing.T) {
	t.Parallel()
	t.Run("not set", func(t *testing.T) {
		t.Parallel()
		assert.Empty(t, env.Get("TEST_ENV_VAR_NOT_SET"))
	})
	t.Run("set", func(t *testing.T) {
		t.Parallel()
		assert.NotEmpty(t, env.Get("PATH"))
	})
}

func TestMustGet(t *testing.T) {
	t.Parallel()
	t.Run("not set", func(t *testing.T) {
		t.Parallel()
		assert.Panics(t, func() {
			env.MustGet("TEST_ENV_VAR_NOT_SET")
		})
	})
	t.Run("set", func(t *testing.T) {
		t.Parallel()
		assert.NotEmpty(t, env.MustGet("PATH"))
	})
}

func TestValue_Empty(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		v    env.Value
		want bool
	}{
		{
			name: "simple string",
			v:    env.NewValue("test"),
			want: false,
		},
		{
			name: "empty",
			v:    env.NewValue(""),
			want: true,
		},
		{
			name: "space",
			v:    env.NewValue(" "),
			want: false,
		},
		{
			name: "zero string",
			v:    env.NewValue("0"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want {
				assert.True(t, tt.v.Empty())
			} else {
				assert.False(t, tt.v.Empty())
			}
		})
	}
}

func TestValue_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		v    env.Value
		want string
	}{
		{
			name: "simple string",
			v:    env.NewValue("test"),
			want: "test",
		},
		{
			name: "empty",
			v:    env.NewValue(""),
			want: "",
		},
		{
			name: "space",
			v:    env.NewValue(" "),
			want: " ",
		},
		{
			name: "zero string",
			v:    env.NewValue("0"),
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.v.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestValue_Int(t *testing.T) {
	t.Parallel()
	type want struct {
		value int
	}
	tests := []struct {
		name string
		v    env.Value
		want want
	}{
		{
			name: "zero int",
			v:    env.NewValue("0"),
			want: want{
				value: 0,
			},
		},
		{
			name: "empty string",
			v:    env.NewValue(""),
			want: want{
				value: 0,
			},
		},
		{
			name: "space string",
			v:    env.NewValue(""),
			want: want{
				value: 0,
			},
		},
		{
			name: "word",
			v:    env.NewValue("hello"),
			want: want{
				value: 0,
			},
		},
		{
			name: "max int64",
			v: env.NewValue(
				fmt.Sprint(
					math.MaxInt64,
				),
			),
			want: want{
				value: math.MaxInt64,
			},
		},
		{
			name: "max int64 negative",
			v: env.NewValue(
				fmt.Sprint(
					-1 * math.MaxInt64,
				),
			),
			want: want{
				value: -1 * math.MaxInt64,
			},
		},
		{
			name: "max int32",
			v: env.NewValue(
				fmt.Sprint(
					math.MaxInt32,
				),
			),
			want: want{
				value: math.MaxInt32,
			},
		},
		{
			name: "max int32 negative",
			v: env.NewValue(
				fmt.Sprint(
					-1 * math.MaxInt32,
				),
			),
			want: want{
				value: -1 * math.MaxInt32,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.v.Int(),
			)
		})
	}
}
