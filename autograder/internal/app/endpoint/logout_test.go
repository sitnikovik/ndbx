package endpoint_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
)

func TestEndpoint_Logout(t *testing.T) {
	t.Parallel()
	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		assert.Equal(
			t,
			"http://localhost:8080/auth/logout",
			endpoint.
				NewEndpoint("http://localhost:8080").
				Logout(),
		)
	})
	t.Run("empty base url", func(t *testing.T) {
		t.Parallel()
		assert.Equal(
			t,
			"/auth/logout",
			endpoint.
				NewEndpoint("").
				Logout(),
		)
	})
}
