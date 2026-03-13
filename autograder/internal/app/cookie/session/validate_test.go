package session_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cookie/session"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

func TestSession_Validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		s               session.Session
		wantErr         error
		wantErrContains string
	}{
		{
			name: "ok",
			s: session.NewSession(
				&http.Cookie{
					Name:     session.Name,
					Path:     "/",
					Value:    "0123456789abcdef0123456789abcdef",
					MaxAge:   3600,
					HttpOnly: true,
					Secure:   true,
				},
			),
			wantErr:         nil,
			wantErrContains: "",
		},
		{
			name: "empty session value",
			s: session.NewSession(
				&http.Cookie{
					Name:     session.Name,
					Path:     "/",
					Value:    "",
					MaxAge:   3600,
					HttpOnly: true,
					Secure:   true,
				},
			),
			wantErr:         errs.ErrExpectationFailed,
			wantErrContains: "must be at least 32 characters long",
		},
		{
			name: "session value is spaces",
			s: session.NewSession(
				&http.Cookie{
					Name:     session.Name,
					Path:     "/",
					Value:    "                                ",
					MaxAge:   3600,
					HttpOnly: true,
					Secure:   true,
				},
			),
			wantErr:         errs.ErrExpectationFailed,
			wantErrContains: "must be a hexadecimal string",
		},
		{
			name: "missing http only flag",
			s: session.NewSession(
				&http.Cookie{
					Name:     session.Name,
					Path:     "/",
					Value:    "0123456789abcdef0123456789abcdef",
					MaxAge:   3600,
					HttpOnly: false,
					Secure:   true,
				},
			),
			wantErr:         errs.ErrMissedCookie,
			wantErrContains: "cookie to have http only flag set",
		},
		{
			name: "missing MaxAge flag",
			s: session.NewSession(
				&http.Cookie{
					Name:     session.Name,
					Path:     "/",
					Value:    "0123456789abcdef0123456789abcdef",
					MaxAge:   0,
					HttpOnly: true,
					Secure:   true,
				},
			),
			wantErr:         errs.ErrMissedCookie,
			wantErrContains: "cookie to have MaxAge flag set",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.s.Validate()
			assert.ErrorIs(t, err, tt.wantErr)
			if tt.wantErrContains != "" {
				assert.ErrorContains(t, err, tt.wantErrContains)
			}
		})
	}
}
