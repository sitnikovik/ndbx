package step_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

func TestVars_With(t *testing.T) {
	t.Parallel()
	t.Run("with new vars", func(t *testing.T) {
		t.Parallel()
		v := step.NewVariables()
		got := v.With(
			step.NewVariable("var1", 123),
			step.NewVariable("var2", "abc"),
		)
		assert.NotSame(t, v, got)
		assert.Equal(t, 0, v.Len())
		assert.Equal(t, 2, got.Len())
	})
}

func TestVars_Get(t *testing.T) {
	t.Parallel()
	type args struct {
		name string
	}
	tests := []struct {
		name string
		v    step.Variables
		args args
		want step.Variable
		ok   bool
	}{
		{
			name: "existing variable",
			v: func() step.Variables {
				v := step.NewVariables()
				v = v.With(step.NewVariable("var1", 123))
				return v
			}(),
			args: args{name: "var1"},
			want: step.NewVariable("var1", 123),
			ok:   true,
		},
		{
			name: "non-existing variable",
			v: func() step.Variables {
				v := step.NewVariables()
				v = v.With(step.NewVariable("var1", 123))
				return v
			}(),
			args: args{name: "var2"},
			want: step.Variable{},
			ok:   false,
		},
		{
			name: "default value",
			v: func() step.Variables {
				v := step.NewVariables()
				return v
			}(),
			args: args{name: "var3"},
			want: step.NewVariable("", nil),
			ok:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, ok := tt.v.Get(tt.args.name)
			assert.Equal(t, tt.want, got)
			if tt.ok {
				assert.True(t, ok)
			} else {
				assert.False(t, ok)
			}
		})
	}
}

func TestVars_Len(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		v    step.Variables
		want int
	}{
		{
			name: "with values",
			v: func() step.Variables {
				v := step.NewVariables()
				v = v.With(
					step.NewVariable("var1", 123),
					step.NewVariable("var2", "abc"),
				)
				return v
			}(),
			want: 2,
		},
		{
			name: "empty variables",
			v: func() step.Variables {
				v := step.NewVariables()
				return v
			}(),
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.v.Len()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestVars_Empty(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		v    step.Variables
		want bool
	}{
		{
			name: "with values",
			v: func() step.Variables {
				v := step.NewVariables()
				v = v.With(
					step.NewVariable("var1", 123),
					step.NewVariable("var2", "abc"),
				)
				return v
			}(),
			want: false,
		},
		{
			name: "empty variables",
			v: func() step.Variables {
				v := step.NewVariables()
				return v
			}(),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.v.Empty()
			if tt.want {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}

func TestVars_MustGet(t *testing.T) {
	t.Parallel()
	type args struct {
		name string
	}
	tests := []struct {
		name string
		v    step.Variables
		args args
		want step.Variable
		ok   bool
	}{
		{
			name: "existing variable",
			v: func() step.Variables {
				v := step.NewVariables()
				v = v.With(step.NewVariable("var1", 123))
				return v
			}(),
			args: args{name: "var1"},
			want: step.NewVariable("var1", 123),
			ok:   true,
		},
		{
			name: "non-existing variable",
			v: func() step.Variables {
				v := step.NewVariables()
				v = v.With(step.NewVariable("var1", 123))
				return v
			}(),
			args: args{name: "var2"},
			want: step.Variable{},
			ok:   false,
		},
		{
			name: "default value",
			v: func() step.Variables {
				v := step.NewVariables()
				return v
			}(),
			args: args{name: "var3"},
			want: step.NewVariable("", nil),
			ok:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.ok {
				got := tt.v.MustGet(tt.args.name)
				assert.Equal(t, tt.want, got)
			} else {
				assert.Panics(t, func() {
					tt.v.MustGet(tt.args.name)
				})
			}
		})
	}
}

func TestVars_Set(t *testing.T) {
	t.Parallel()
	type args struct {
		name  string
		value any
	}
	tests := []struct {
		name string
		v    step.Variables
		args args
		want step.Variable
	}{
		{
			name: "set a new one",
			v:    step.NewVariables(),
			args: args{
				name:  "var1",
				value: 123,
			},
			want: step.NewVariable("var1", 123),
		},
		{
			name: "update an existing one",
			v: func() step.Variables {
				v := step.NewVariables()
				v.Set("var1", 123)
				return v
			}(),
			args: args{
				name:  "var1",
				value: 456,
			},
			want: step.NewVariable("var1", 456),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.v.Set(tt.args.name, tt.args.value)
			got, _ := tt.v.Get(tt.args.name)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestVars_Del(t *testing.T) {
	t.Parallel()
	type args struct {
		name string
	}
	type want struct {
		v step.Variables
	}
	tests := []struct {
		name string
		v    step.Variables
		args args
		want want
	}{
		{
			name: "existing variable",
			v: func() step.Variables {
				v := step.NewVariables()
				v.Set("var1", 123)
				return v
			}(),
			args: args{name: "var1"},
			want: want{
				v: step.NewVariables(),
			},
		},
		{
			name: "non-existing variable",
			v: func() step.Variables {
				v := step.NewVariables()
				v.Set("var1", 123)
				return v
			}(),
			args: args{name: "var2"},
			want: want{
				v: func() step.Variables {
					v := step.NewVariables()
					v.Set("var1", 123)
					return v
				}(),
			},
		},
		{
			name: "default value",
			v:    step.NewVariables(),
			args: args{name: "var3"},
			want: want{
				v: step.NewVariables(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.v.Del(tt.args.name)
			assert.Equal(
				t,
				tt.want.v,
				tt.v,
			)
		})
	}
}
