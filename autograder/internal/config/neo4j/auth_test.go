package neo4j_test

import (
	"testing"

	impl "github.com/sitnikovik/ndbx/autograder/internal/config/neo4j"
	"github.com/stretchr/testify/assert"
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
