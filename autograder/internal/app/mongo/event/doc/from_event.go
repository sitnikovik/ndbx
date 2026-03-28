package doc

import (
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/mongo/event/doc/key"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
)

// FromEvent converts the Event to an EventDocument related to MongoDB.
func FromEvent(e event.Event) EventDocument {
	kvs := make(doc.KVs, 0, 9)
	if v := e.Content().Title(); v != "" {
		kvs = append(kvs, doc.NewKV(
			key.Title,
			v,
		))
	}
	if v := e.Content().Description(); v != "" {
		kvs = append(kvs, doc.NewKV(
			key.Description,
			v,
		))
	}
	if v := e.Location(); !v.Empty() {
		m := make(map[string]string, 2)
		if v := v.Address(); v != "" {
			m["address"] = v
		}
		if v := v.City(); v != "" {
			m["city"] = v
		}
		kvs = append(kvs, doc.NewKV(
			key.Location,
			m,
		))
	}
	if v := e.Created().At(); !v.IsZero() {
		kvs = append(kvs, doc.NewKV(
			key.CreatedAt,
			v.Format(time.RFC3339),
		))
	}
	if v := e.Created().By().ID(); !v.Empty() {
		kvs = append(kvs, doc.NewKV(
			key.CreatedBy,
			v.String(),
		))
	}
	if v := e.Dates().StartedAt(); !v.IsZero() {
		kvs = append(kvs, doc.NewKV(
			key.StartedAt,
			v.Format(time.RFC3339),
		))
	}
	if v := e.Dates().FinishedAt(); !v.IsZero() {
		kvs = append(kvs, doc.NewKV(
			key.FinishedAt,
			v.Format(time.RFC3339),
		))
	}
	kvs = append(kvs, doc.NewKV(
		key.Price,
		e.Costs().Entry().Units(),
	))
	kvs = append(kvs, doc.NewKV(
		key.Category,
		e.Content().Category().String(),
	))
	return NewEventDocument(
		doc.NewDocument(
			e.ID().String(),
			kvs...,
		),
	)
}
