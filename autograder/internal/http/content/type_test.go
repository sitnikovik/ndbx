package content_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/http/content"
)

func TestContentType_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		t    content.Type
		want string
	}{
		{
			name: "ApplicationJSON",
			t:    content.ApplicationJSON,
			want: "application/json",
		},
		{
			name: "custom content type",
			t:    content.Type("custom/content-type"),
			want: "custom/content-type",
		},
		{
			name: "empty content type",
			t:    content.Type(""),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.t.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
