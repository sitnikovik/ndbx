package endpoint_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
)

func TestEndpoint_Session(t *testing.T) {
	t.Parallel()
	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		assert.Equal(
			t,
			"http://localhost:8080/session",
			endpoint.
				NewEndpoint("http://localhost:8080").
				Session(),
		)
	})
	t.Run("empty base url", func(t *testing.T) {
		t.Parallel()
		assert.Equal(
			t,
			"/session",
			endpoint.
				NewEndpoint("").
				Session(),
		)
	})
}
