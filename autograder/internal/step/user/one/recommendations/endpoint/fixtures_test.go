package endpoint_test

import (
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
	respxpct "github.com/sitnikovik/ndbx/autograder/internal/expect/http/response/expectation"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// NewDescFx returns a new step desc fixture.
func NewDescFx() step.Desc {
	return step.NewDesc("Title", "Description")
}

// NewResponseXpctFx returns a new response expectation fixture.
func NewResponseXpctFx() respxpct.Expectations {
	return respxpct.NewExpectations(
		respxpct.WithAsserts(
			response.AssertOKStatus,
			response.AssertNotEmptyContent,
		),
	)
}
