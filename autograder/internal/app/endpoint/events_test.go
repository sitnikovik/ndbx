package endpoint_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
)

func TestEndpoint_Events(t *testing.T) {
	t.Parallel()
	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		assert.Equal(
			t,
			"http://localhost:8080/events",
			endpoint.
				NewEndpoint("http://localhost:8080").
				Events(),
		)
	})
	t.Run("empty base url", func(t *testing.T) {
		t.Parallel()
		assert.Equal(
			t,
			"/events",
			endpoint.
				NewEndpoint("").
				Events(),
		)
	})
}
