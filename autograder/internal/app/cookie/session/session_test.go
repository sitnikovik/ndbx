package session_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cookie/session"
)

func TestSession_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		s    session.Session
		want string
	}{
		{
			name: "ok",
			s: session.NewSession(
				&http.Cookie{
					Value: "sad12efrfdsaxzsddsqa",
				},
			),
			want: "sad12efrfdsaxzsddsqa",
		},
		{
			name: "empty",
			s: session.NewSession(
				&http.Cookie{
					Value: "",
				},
			),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want,
				tt.s.String(),
			)
		})
	}
}

func TestSession_Expired(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		s    session.Session
		want bool
	}{
		{
			name: "not expired",
			s: session.NewSession(
				&http.Cookie{
					MaxAge: 3600,
				},
			),
			want: false,
		},
		{
			name: "expired",
			s: session.NewSession(
				&http.Cookie{
					MaxAge: 0,
				},
			),
			want: true,
		},
		{
			name: "expired with negative MaxAge",
			s: session.NewSession(
				&http.Cookie{
					MaxAge: -1,
				},
			),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want {
				assert.True(t, tt.s.Expired())
			} else {
				assert.False(t, tt.s.Expired())
			}
		})
	}
}
