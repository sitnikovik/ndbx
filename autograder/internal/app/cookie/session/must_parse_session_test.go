package session_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cookie/session"
)

func TestMustParseSession(t *testing.T) {
	t.Parallel()
	type args struct {
		ckk []*http.Cookie
	}
	tests := []struct {
		name  string
		args  args
		want  session.Session
		panic bool
	}{
		{
			name: "ok",
			args: args{
				ckk: []*http.Cookie{
					{Name: session.Name, Value: "sad12efrfdsaxzsddsqa"},
				},
			},
			want: session.NewSession(
				&http.Cookie{
					Name:  session.Name,
					Value: "sad12efrfdsaxzsddsqa",
				},
			),
		},
		{
			name: "no session cookie",
			args: args{
				ckk: []*http.Cookie{
					{Name: "other_cookie", Value: "value"},
				},
			},
			panic: true,
		},
		{
			name: "empty session cookie value",
			args: args{
				ckk: []*http.Cookie{
					{Name: session.Name, Value: ""},
				},
			},
			want: session.NewSession(
				&http.Cookie{
					Name:  session.Name,
					Value: "",
				},
			),
			panic: false,
		},
		{
			name: "empty cookies",
			args: args{
				ckk: []*http.Cookie{},
			},
			panic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.panic {
				assert.Panics(t, func() {
					_ = session.MustParseSession(tt.args.ckk)
				})
				return
			}
			assert.Equal(
				t,
				tt.want,
				session.MustParseSession(tt.args.ckk),
			)
		})
	}
}
