package refresh_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/refresh"
)

func TestJob_Name(t *testing.T) {
	t.Parallel()
	assert.Equal(
		t,
		refresh.Name,
		refresh.NewJob().Name(),
	)
}

func TestJob_Description(t *testing.T) {
	t.Parallel()
	assert.Equal(
		t,
		refresh.Description,
		refresh.NewJob().Description(),
	)
}
