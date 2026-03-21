package doc

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/money"
	"github.com/sitnikovik/ndbx/autograder/internal/app/mongo/event/doc/key"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/console"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

// ToEvent converts the EventDocument struct to an Event struct.
//
// Panics if the MongoDB document does not contain the required fields or if the field types are incorrect.
func (d EventDocument) ToEvent() event.Event {
	var (
		err        error
		title      string
		desc       string
		addr       string
		loc        event.Location
		createdAt  time.Time
		createdBy  user.Identity
		startedAt  time.Time
		finishedAt time.Time
		price      uint64
	)
	for _, kv := range d.orig.KVs() {
		console.Log(
			"key: %s, value: %v, type: %T",
			kv.Key(),
			kv.Value(),
			kv.Value(),
		)
	}
	for _, kv := range d.orig.KVs() {
		switch kv.Key() {
		case key.Title:
			if v, ok := kv.Value().(string); ok {
				title = v
			}
		case key.Description:
			if v, ok := kv.Value().(string); ok {
				desc = v
			}
		case key.Location:
			var opts []event.LocationOption
			if v, ok := kv.Value().(string); ok {
				var jsn struct {
					Address string `json:"address"`
					City    string `json:"city"`
				}
				err := json.Unmarshal([]byte(v), &jsn)
				if err == nil {
					if jsn.Address != "" {
						addr = jsn.Address
					}
					if jsn.City != "" {
						opts = append(opts, event.WithCity(jsn.City))
					}
				} else {
					addr = v
				}
				loc = event.NewLocation(addr, opts...)
			}
		case key.CreatedAt:
			if v, ok := kv.Value().(string); ok {
				createdAt = timex.MustParse(time.RFC3339, v)
			}
		case key.CreatedBy:
			if v, ok := kv.Value().(string); ok {
				createdBy = user.NewIdentity(user.NewID(v))
			}
		case key.StartedAt:
			if v, ok := kv.Value().(string); ok {
				startedAt = timex.MustParse(time.RFC3339, v)
			}
		case key.FinishedAt:
			if v, ok := kv.Value().(string); ok {
				finishedAt = timex.MustParse(time.RFC3339, v)
			}
		case key.Price:
			if v, ok := kv.Value().(string); ok {
				price, err = strconv.ParseUint(v, 10, 64)
				if err != nil {
					panic(fmt.Sprintf("failed to parse price: %v", err))
				}
			}
		}
	}
	evnt := event.NewEvent(
		event.NewID(d.orig.ID()),
		event.NewContent(
			title,
			desc,
		),
		loc,
		event.NewCreated(
			createdAt,
			createdBy,
		),
		event.NewDates(
			startedAt,
			finishedAt,
		),
		event.WithCosts(
			event.NewCosts(
				money.NewMoney(price, 0),
			),
		),
	)
	return evnt
}
