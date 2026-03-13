package cookie_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/http/response/cookie"
)

func TestCookies_Has(t *testing.T) {
	t.Parallel()
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		cookies *cookie.Cookies
		args    args
		want    bool
	}{
		{
			name: "cookie exists",
			cookies: cookie.NewCookies([]*http.Cookie{
				{Name: "session_id", Value: "abc123"},
				{Name: "user_id", Value: "42"},
			}),
			args: args{name: "session_id"},
			want: true,
		},
		{
			name: "cookie does not exist",
			cookies: cookie.NewCookies([]*http.Cookie{
				{Name: "session_id", Value: "abc123"},
				{Name: "user_id", Value: "42"},
			}),
			args: args{name: "auth_token"},
			want: false,
		},
		{
			name: "empty cookie name",
			cookies: cookie.NewCookies([]*http.Cookie{
				{Name: "session_id", Value: "abc123"},
				{Name: "user_id", Value: "42"},
			}),
			args: args{name: ""},
			want: false,
		},
		{
			name:    "no cookies",
			cookies: cookie.NewCookies([]*http.Cookie{}),
			args:    args{name: "session_id"},
			want:    false,
		},
		{
			name:    "nil cookies",
			cookies: cookie.NewCookies(nil),
			args:    args{name: "session_id"},
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.cookies.Has(tt.args.name)
			if tt.want {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}

func TestCookies_MustGet(t *testing.T) {
	t.Parallel()
	type args struct {
		name string
	}
	tests := []struct {
		name      string
		cookies   *cookie.Cookies
		args      args
		want      *http.Cookie
		wantPanic bool
	}{
		{
			name: "cookie exists",
			cookies: cookie.NewCookies([]*http.Cookie{
				{
					Name:  "session_id",
					Value: "abc123",
				},
				{
					Name:  "user_id",
					Value: "42",
				},
			}),
			args: args{
				name: "session_id",
			},
			want: &http.Cookie{
				Name:  "session_id",
				Value: "abc123",
			},
		},
		{
			name: "cookie does not exist",
			cookies: cookie.NewCookies([]*http.Cookie{
				{
					Name:  "session_id",
					Value: "abc123",
				},
				{
					Name:  "user_id",
					Value: "42",
				},
			}),
			args: args{
				name: "auth_token",
			},
			wantPanic: true,
		},
		{
			name: "empty cookie value",
			cookies: cookie.NewCookies([]*http.Cookie{
				{
					Name:  "session_id",
					Value: "",
				},
				{
					Name:  "user_id",
					Value: "42",
				},
			}),
			args: args{
				name: "session_id",
			},
			want: &http.Cookie{
				Name:  "session_id",
				Value: "",
			},
			wantPanic: false,
		},
		{
			name: "empty cookie name",
			cookies: cookie.NewCookies([]*http.Cookie{
				{
					Name:  "session_id",
					Value: "abc123",
				},
				{
					Name:  "user_id",
					Value: "42",
				},
			}),
			args: args{
				name: "",
			},
			wantPanic: true,
		},
		{
			name:    "no cookies",
			cookies: cookie.NewCookies([]*http.Cookie{}),
			args: args{
				name: "session_id",
			},
			wantPanic: true,
		},
		{
			name:    "nil cookies",
			cookies: cookie.NewCookies(nil),
			args: args{
				name: "session_id",
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.wantPanic {
				assert.Panics(t, func() {
					_ = tt.cookies.MustGet(tt.args.name)
				})
				return
			}
			assert.Equal(
				t,
				tt.want,
				tt.cookies.MustGet(
					tt.args.name,
				),
			)
		})
	}
}
