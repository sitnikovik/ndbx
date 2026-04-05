package event

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event/category"
)

// Content represents content info of the event.
type Content struct {
	// title is the title of the event.
	title string
	// desc is the description of the event.
	desc string
	// cat is the category of the event.
	cat category.Type
}

// ContentOption represents a functional option
// for configuring the content of the event.
type ContentOption func(c *Content)

// WithCategory sets the category type of the event.
func WithCategory(cat category.Type) ContentOption {
	return func(c *Content) {
		c.cat = cat
	}
}

// NewContent creates a new Content instance.
func NewContent(
	title string,
	desc string,
	opts ...ContentOption,
) Content {
	c := Content{
		title: title,
		desc:  desc,
	}
	for _, opt := range opts {
		opt(&c)
	}
	return c
}

// Title returns the title of the event.
func (c Content) Title() string {
	return c.title
}

// TitleHash returns the hash of the title of the event.
func (c Content) TitleHash() string {
	v := c.Title()
	if v == "" {
		return ""
	}
	hash := md5.Sum([]byte(v))
	return hex.EncodeToString(hash[:])
}

// Description returns the description of the event.
func (c Content) Description() string {
	return c.desc
}

// Category returns the category of the event.
func (c Content) Category() category.Type {
	return c.cat
}
