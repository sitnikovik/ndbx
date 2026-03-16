package body_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/users/post/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

func TestBody_MustBytes(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		b    body.Body
		want []byte
	}{
		{
			name: "ok",
			b: body.NewBody(
				user.NewUser(
					user.NewID("123"),
					"sams3p1ol",
					"Sam Sepiol",
				),
				"svp4_dvp4_s3cr3t_p4ssw0rd",
			),
			want: []byte(`{` +
				`"full_name":"Sam Sepiol",` +
				`"password":"svp4_dvp4_s3cr3t_p4ssw0rd",` +
				`"username":"sams3p1ol"` +
				`}`),
		},
		{
			name: "all empty",
			b: body.NewBody(
				user.NewUser(
					user.NewID(""),
					"",
					"",
				),
				"",
			),
			want: []byte(`{` +
				`"full_name":"",` +
				`"password":"",` +
				`"username":""` +
				`}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want,
				tt.b.MustBytes(),
			)
		})
	}
}
