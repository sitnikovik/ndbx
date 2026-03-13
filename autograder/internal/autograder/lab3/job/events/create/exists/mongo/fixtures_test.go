package mongo_test

import (
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/app/mongo/event/doc/key"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
)

// NewOkEventDocKVFixture creates a new OK event document key-value fixture.
func NewOkEventDocKVFixture() []doc.KV {
	ent := lab3.NewTestEvent()
	return []doc.KV{
		doc.NewKV(
			key.Title,
			ent.
				Content().
				Title(),
		),
		doc.NewKV(
			key.Description,
			ent.
				Content().
				Description(),
		),
		doc.NewKV(
			key.Location,
			ent.Location().Address(),
		),
		doc.NewKV(
			key.CreatedBy,
			ent.
				Created().
				By().
				ID().
				String(),
		),
		doc.NewKV(
			key.CreatedAt,
			ent.
				Created().
				At().
				Format(time.RFC3339),
		),
		doc.NewKV(
			key.StartedAt,
			ent.
				Dates().
				StartedAt().
				Format(time.RFC3339),
		),
		doc.NewKV(
			key.StartedAt,
			ent.
				Dates().
				StartedAt().
				Format(time.RFC3339),
		),
		doc.NewKV(
			key.FinishedAt,
			ent.
				Dates().
				FinishedAt().
				Format(time.RFC3339),
		),
	}
}
