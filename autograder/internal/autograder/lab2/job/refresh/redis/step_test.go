package redis_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/refresh/redis"
	redisfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/redis"
)

func TestStep_Name(t *testing.T) {
	t.Parallel()
	assert.Equal(
		t,
		redis.Name,
		redis.
			NewStep(redisfk.NewFakeClient()).
			Name(),
	)
}

func TestStep_Description(t *testing.T) {
	t.Parallel()
	assert.Equal(
		t,
		redis.Description,
		redis.
			NewStep(redisfk.NewFakeClient()).
			Description(),
	)
}
