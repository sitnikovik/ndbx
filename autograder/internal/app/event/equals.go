package event

// Equals defines if the event equals to the provided one.
//
// Do not compares the objects with reviews, reactions and tags.
func (e Event) Equals(other Event) bool {
	return e.id == other.id &&
		e.created.Equals(other.created) &&
		e.content.Equals(other.content) &&
		e.loc.Equals(other.loc) &&
		e.dates.Equals(other.dates) &&
		e.costs.Equals(other.costs)
}
