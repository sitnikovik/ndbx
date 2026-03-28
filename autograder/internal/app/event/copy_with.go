package event

// CopyWith creates a copy of the event with the given options applied.
func (e Event) CopyWith(opts ...Option) Event {
	cop := e
	for _, opt := range opts {
		opt(&cop)
	}
	return cop
}
