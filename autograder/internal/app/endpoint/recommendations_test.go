package endpoint_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
)

func TestEndpoint_Recommendations(t *testing.T) {
	t.Parallel()
	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		assert.Equal(
			t,
			"http://localhost:8080/recommendations",
			impl.
				NewEndpoint("http://localhost:8080").
				Recommendations(),
		)
	})
	t.Run("empty base url", func(t *testing.T) {
		t.Parallel()
		assert.Equal(
			t,
			"/recommendations",
			impl.
				NewEndpoint("").
				Recommendations(),
		)
	})
}
