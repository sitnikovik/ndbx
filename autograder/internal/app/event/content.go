package event

// Content represents content info of the event.
type Content struct {
	// title is the title of the event.
	title string
	// desc is the description of the event.
	desc string
}

// NewContent creates a new Content instance.
func NewContent(
	title string,
	desc string,
) Content {
	return Content{
		title: title,
		desc:  desc,
	}
}

// Title returns the title of the event.
func (c Content) Title() string {
	return c.title
}

// Description returns the description of the event.
func (c Content) Description() string {
	return c.desc
}
