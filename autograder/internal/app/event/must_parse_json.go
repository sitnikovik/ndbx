package event

import (
	"encoding/json"
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

// MustParseJSON parses the JSON content of the event.
//
// Panics if parsing the JSON fails or if the required fields are missing.
func MustParseJSON(bb []byte) Event {
	var v struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Location    struct {
			Address string `json:"address"`
		} `json:"location"`
		CreatedAt  string `json:"created_at"`
		CreatedBy  string `json:"created_by"`
		StartedAt  string `json:"started_at"`
		FinishedAt string `json:"finished_at"`
		Reactions  struct {
			Likes    uint64 `json:"likes"`
			Dislikes uint64 `json:"dislikes"`
		} `json:"reactions"`
	}
	err := json.Unmarshal(bb, &v)
	if err != nil {
		panic(err)
	}
	var createdAt time.Time
	if v.CreatedAt != "" {
		createdAt = timex.MustRFC3339(v.CreatedAt)
	}
	var startedAt time.Time
	if v.StartedAt != "" {
		startedAt = timex.MustRFC3339(v.StartedAt)
	}
	var finishedAt time.Time
	if v.FinishedAt != "" {
		finishedAt = timex.MustRFC3339(v.FinishedAt)
	}
	return NewEvent(
		ID(v.ID),
		NewContent(
			v.Title,
			v.Description,
		),
		NewLocation(v.Location.Address),
		NewCreated(
			createdAt,
			user.NewIdentity(
				user.NewID(v.CreatedBy),
			),
		),
		NewDates(
			startedAt,
			finishedAt,
		),
		WithLikes(v.Reactions.Likes),
		WithDislikes(v.Reactions.Dislikes),
	)
}
