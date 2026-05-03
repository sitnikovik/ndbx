package event

import (
	"encoding/json"
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/app/creation"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/rating"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

// MustParseJSON parses the JSON content of the review.
//
// Panics if parsing the JSON fails or if the required fields are missing.
func MustParseJSON(bb []byte) Review {
	var v struct {
		ID        string  `json:"id"`
		EventID   string  `json:"event_id"`
		Comment   string  `json:"comment"`
		CreatedAt string  `json:"created_at"`
		UpdatedAt string  `json:"updated_at"`
		CreatedBy string  `json:"created_by"`
		Rating    float64 `json:"rating"`
	}
	err := json.Unmarshal(bb, &v)
	if err != nil {
		panic(err)
	}
	var createdAt time.Time
	if v.CreatedAt != "" {
		createdAt = timex.MustRFC3339(v.CreatedAt)
	}
	var updatedAt time.Time
	if v.UpdatedAt != "" {
		updatedAt = timex.MustRFC3339(v.UpdatedAt)
	}
	return NewReview(
		v.ID,
		creation.NewStamp(
			creation.NewCreated(
				createdAt,
				user.NewIdentity(
					user.NewID(v.CreatedBy),
				),
			),
		),
		NewEvent(
			event.NewID(
				v.EventID,
			),
		),
		v.Comment,
		rating.NewRating(
			v.Rating,
		),
		WithUpdatedAt(updatedAt),
	)
}
