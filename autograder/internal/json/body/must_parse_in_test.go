package body_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/json/body"
)

func TestBody_MustParseIn(t *testing.T) {
	t.Parallel()
	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		data := `{"name":"test","value":42}`
		var v struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		}
		body := body.NewBody(strings.NewReader(data))
		body.MustParseIn(&v)
		assert.Equal(t, "test", v.Name)
		assert.Equal(t, 42, v.Value)
	})
	t.Run("invalid JSON", func(t *testing.T) {
		t.Parallel()
		data := `{"name":"test","value":42` // Missing closing brace
		var v struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		}
		body := body.NewBody(strings.NewReader(data))
		assert.Panics(t, func() {
			body.MustParseIn(&v)
		})
	})
	t.Run("type mismatch", func(t *testing.T) {
		t.Parallel()
		data := `{"name":"test","value":"42"}`
		var v struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		}
		body := body.NewBody(strings.NewReader(data))
		assert.Panics(t, func() {
			body.MustParseIn(&v)
		})
	})
	t.Run("empty json", func(t *testing.T) {
		t.Parallel()
		data := `{}`
		var v struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		}
		body := body.NewBody(strings.NewReader(data))
		body.MustParseIn(&v)
		assert.Equal(t, "", v.Name)
		assert.Equal(t, 0, v.Value)
	})
	t.Run("empty string", func(t *testing.T) {
		t.Parallel()
		data := ``
		var v struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		}
		body := body.NewBody(strings.NewReader(data))
		assert.Panics(t, func() {
			body.MustParseIn(&v)
		})
	})
}
