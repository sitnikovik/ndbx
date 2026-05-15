package neo4j_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/sitnikovik/ndbx/autograder/internal/config/neo4j"
)

func TestAuth_Username(t *testing.T) {
	t.Parallel()
	type want struct {
		value string
	}
	tests := []struct {
		name string
		a    impl.Auth
		want want
	}{
		{
			name: "ok",
			a:    impl.NewAuth("user", "pass"),
			want: want{
				value: "user",
			},
		},
		{
			name: "all empty",
			a:    impl.NewAuth("", ""),
			want: want{
				value: "",
			},
		},
		{
			name: "empty username",
			a:    impl.NewAuth("", "pass"),
			want: want{
				value: "",
			},
		},
		{
			name: "empty password",
			a:    impl.NewAuth("user", ""),
			want: want{
				value: "user",
			},
		},
		{
			name: "default value",
			a:    impl.Auth{},
			want: want{
				value: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.a.Username(),
			)
		})
	}
}
func TestAuth_Password(t *testing.T) {
	t.Parallel()
	type want struct {
		value string
	}
	tests := []struct {
		name string
		a    impl.Auth
		want want
	}{
		{
			name: "ok",
			a:    impl.NewAuth("user", "pass"),
			want: want{
				value: "pass",
			},
		},
		{
			name: "all empty",
			a:    impl.NewAuth("", ""),
			want: want{
				value: "",
			},
		},
		{
			name: "empty username",
			a:    impl.NewAuth("", "pass"),
			want: want{
				value: "pass",
			},
		},
		{
			name: "empty password",
			a:    impl.NewAuth("user", ""),
			want: want{
				value: "",
			},
		},
		{
			name: "default value",
			a:    impl.Auth{},
			want: want{
				value: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.a.Password(),
			)
		})
	}
}

func TestAuth_Validate(t *testing.T) {
	t.Parallel()
	type want struct {
		errContains string
		errored     bool
	}
	tests := []struct {
		name string
		a    impl.Auth
		want want
	}{
		{
			name: "ok",
			a:    impl.NewAuth("user", "pass"),
			want: want{
				errored: false,
			},
		},
		{
			name: "empty username",
			a:    impl.NewAuth("", "pass"),
			want: want{
				errContains: "username",
				errored:     true,
			},
		},
		{
			name: "empty password",
			a:    impl.NewAuth("user", ""),
			want: want{
				errContains: "password",
				errored:     true,
			},
		},
		{
			name: "all empty",
			a:    impl.NewAuth("", ""),
			want: want{
				errContains: "",
				errored:     false,
			},
		},
		{
			name: "whitespace username",
			a:    impl.NewAuth(" ", ""),
			want: want{
				errContains: "password",
				errored:     true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.a.Validate()
			if tt.want.errored {
				assert.Error(t, err)
				assert.Contains(
					t,
					err.Error(),
					tt.want.errContains,
				)
			} else {
				assert.NoErrorf(
					t,
					err,
					"unexpected error: %v",
					err,
				)
			}
		})
	}
}

func TestAuth_Empty(t *testing.T) {
	t.Parallel()
	type want struct {
		value bool
	}
	tests := []struct {
		name string
		a    impl.Auth
		want want
	}{
		{
			name: "ok",
			a:    impl.NewAuth("user", "pass"),
			want: want{
				value: false,
			},
		},
		{
			name: "empty username",
			a:    impl.NewAuth("", "pass"),
			want: want{
				value: false,
			},
		},
		{
			name: "empty password",
			a:    impl.NewAuth("user", ""),
			want: want{
				value: false,
			},
		},
		{
			name: "all empty",
			a:    impl.NewAuth("", ""),
			want: want{
				value: true,
			},
		},
		{
			name: "whitespace username",
			a:    impl.NewAuth(" ", ""),
			want: want{
				value: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.a.Empty()
			if tt.want.value {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}
