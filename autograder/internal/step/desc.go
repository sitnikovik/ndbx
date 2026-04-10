package step

// Desc represents the description details of the step to run.
type Desc struct {
	// title is the title of the step.
	title string
	// description is the short brief description of what the step does.
	description string
}

// NewDesc creates a new Desc instance.
func NewDesc(
	title string,
	description string,
) Desc {
	return Desc{
		title:       title,
		description: description,
	}
}

// Title returns the title of the step.
func (d Desc) Title() string {
	return d.title
}

// Description returns a brief description of what the step does.
func (d Desc) Description() string {
	return d.description
}
