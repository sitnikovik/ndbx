package httpx_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/sitnikovik/ndbx/autograder/internal/client/httpx"
)

func TestNewClient(t *testing.T) {
	t.Parallel()
	t.Run("without options", func(t *testing.T) {
		t.Parallel()
		c := httpx.NewClient()
		require.NotNil(t, c)
		require.Equal(t, httpx.DefaultTimeout, c.Timeout())
	})
	t.Run("with custom timeout", func(t *testing.T) {
		t.Parallel()
		timeout := 10 * time.Second
		c := httpx.NewClient(httpx.WithTimeout(timeout))
		require.NotNil(t, c)
		require.Equal(t, timeout, c.Timeout())
	})
}
