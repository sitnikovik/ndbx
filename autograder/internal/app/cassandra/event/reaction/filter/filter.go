package filter

import (
	"github.com/sitnikovik/ndbx/autograder/internal/app/creation"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/console"
)

// where implements WHERE clause.
type where interface {
	// Add adds a new WHERE condition with the given field and value.
	Add(field string, value any)
	// String returns WHERE clause as a string.
	String() string
	// Args returns arguments for the WHERE clause.
	Args() []any
}

// Filter represents filter for reactions.
type Filter struct {
	// where is the WHERE clause.
	where where
	// stamp is the creation stamp of the reaction.
	stamp creation.Stamp
	// eventID is the ID of the event.
	eventID event.ID
	// like is the value defines
	// whether it is like or dislike or something else.
	like Like
}

// NewFilter creates a new Filter instance.
func NewFilter(w where, opts ...Option) Filter {
	f := Filter{
		where: w,
	}
	for _, opt := range opts {
		opt(&f)
	}
	return f
}

// Empty defines whether the filter is empty.
func (f Filter) Empty() bool {
	return f.like.Empty() &&
		f.eventID.Empty() &&
		f.stamp.Created().By().ID().Empty()
}

// Where returns WHERE clause for the filter.
func (f Filter) Where() string {
	if v := f.like; !v.Empty() {
		f.where.Add("like_value", v.Value())
	}
	if v := f.eventID; !v.Empty() {
		f.where.Add("event_id", v.String())
	}
	if v := f.stamp.Created().By().ID(); !v.Empty() {
		f.where.Add("created_by", v.String())
	}
	return f.where.String()
}

// Args returns arguments for the WHERE clause.
func (f Filter) Args() []any {
	vv := f.where.Args()
	if len(vv) != 0 {
		console.Log("%v", vv)
		return vv
	}
	vv = make([]any, 0, 3)
	if v := f.like; !v.Empty() {
		vv = append(vv, v.Value())
	}
	if v := f.eventID; !v.Empty() {
		vv = append(vv, v.String())
	}
	if v := f.stamp.Created().By().ID(); !v.Empty() {
		vv = append(vv, v.String())
	}
	return vv
}
