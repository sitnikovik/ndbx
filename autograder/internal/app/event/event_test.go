package event_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/money"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestEvent_Content(t *testing.T) {
	t.Parallel()
	type want struct {
		val event.Content
	}
	tests := []struct {
		name string
		e    event.Event
		want want
	}{
		{
			name: "ok",
			e: event.NewEvent(
				event.NewID("1"),
				event.NewContent(
					"Title",
					"description",
				),
				event.NewLocation(
					"Main street, 13",
				),
				event.NewCreated(
					time.Time{},
					user.NewIdentity(
						user.NewID("123"),
					),
				),
				event.NewDates(
					time.Time{},
					time.Time{},
				),
			),
			want: want{
				val: event.NewContent(
					"Title",
					"description",
				),
			},
		},
		{
			name: "default value",
			e:    event.Event{},
			want: want{
				val: event.Content{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.e.Content(),
			)
		})
	}
}

func TestEvent_Location(t *testing.T) {
	t.Parallel()
	type want struct {
		val event.Location
	}
	tests := []struct {
		name string
		e    event.Event
		want want
	}{
		{
			name: "ok",
			e: event.NewEvent(
				event.NewID("1"),
				event.NewContent(
					"Title",
					"description",
				),
				event.NewLocation(
					"Main street, 13",
				),
				event.NewCreated(
					time.Time{},
					user.NewIdentity(
						user.NewID("123"),
					),
				),
				event.NewDates(
					time.Time{},
					time.Time{},
				),
			),
			want: want{
				val: event.NewLocation(
					"Main street, 13",
				),
			},
		},
		{
			name: "default value",
			e:    event.Event{},
			want: want{
				val: event.Location{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.e.Location(),
			)
		})
	}
}

func TestEvent_Dates(t *testing.T) {
	t.Parallel()
	type want struct {
		val event.Dates
	}
	tests := []struct {
		name string
		e    event.Event
		want want
	}{
		{
			name: "ok",
			e: event.NewEvent(
				event.NewID("1"),
				event.NewContent(
					"Title",
					"description",
				),
				event.NewLocation(
					"Main street, 13",
				),
				event.NewCreated(
					time.Time{},
					user.NewIdentity(
						user.NewID("123"),
					),
				),
				event.NewDates(
					timex.MustParse(time.RFC3339, "2025-01-01T12:00:00Z"),
					timex.MustParse(time.RFC3339, "2025-01-01T13:00:00Z"),
				),
			),
			want: want{
				val: event.NewDates(
					timex.MustParse(time.RFC3339, "2025-01-01T12:00:00Z"),
					timex.MustParse(time.RFC3339, "2025-01-01T13:00:00Z"),
				),
			},
		},
		{
			name: "default value",
			e:    event.Event{},
			want: want{
				val: event.Dates{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.e.Dates(),
			)
		})
	}
}

func TestEvent_Quantity(t *testing.T) {
	t.Parallel()
	type want struct {
		val event.Quantity
	}
	tests := []struct {
		name string
		e    event.Event
		want want
	}{
		{
			name: "ok",
			e: event.NewEvent(
				event.NewID("1"),
				event.NewContent(
					"Title",
					"description",
				),
				event.NewLocation(
					"Main street, 13",
				),
				event.NewCreated(
					time.Time{},
					user.NewIdentity(
						user.NewID("123"),
					),
				),
				event.NewDates(
					timex.MustParse(time.RFC3339, "2025-01-01T12:00:00Z"),
					timex.MustParse(time.RFC3339, "2025-01-01T13:00:00Z"),
				),
				event.WithQuantity(event.NewQuantity(0, 10)),
			),
			want: want{
				val: event.NewQuantity(0, 10),
			},
		},
		{
			name: "default value",
			e:    event.Event{},
			want: want{
				val: event.Quantity{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.e.Quantity(),
			)
		})
	}
}

func TestEvent_ID(t *testing.T) {
	t.Parallel()
	type want struct {
		val event.ID
	}
	tests := []struct {
		name string
		e    event.Event
		want want
	}{
		{
			name: "ok",
			e: event.NewEvent(
				event.NewID("1"),
				event.NewContent(
					"Title",
					"description",
				),
				event.NewLocation(
					"Main street, 13",
				),
				event.NewCreated(
					time.Time{},
					user.NewIdentity(
						user.NewID("123"),
					),
				),
				event.NewDates(
					timex.MustParse(time.RFC3339, "2025-01-01T12:00:00Z"),
					timex.MustParse(time.RFC3339, "2025-01-01T13:00:00Z"),
				),
				event.WithQuantity(event.NewQuantity(0, 10)),
			),
			want: want{
				val: event.NewID("1"),
			},
		},
		{
			name: "default value",
			e:    event.Event{},
			want: want{
				val: event.NewID(""),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.e.ID(),
			)
		})
	}
}

func TestEvent_Costs(t *testing.T) {
	t.Parallel()
	type want struct {
		val event.Costs
	}
	tests := []struct {
		name string
		e    event.Event
		want want
	}{
		{
			name: "ok",
			e: event.NewEvent(
				event.NewID("1"),
				event.NewContent(
					"Title",
					"description",
				),
				event.NewLocation(
					"Main street, 13",
				),
				event.NewCreated(
					time.Time{},
					user.NewIdentity(
						user.NewID("123"),
					),
				),
				event.NewDates(
					timex.MustParse(time.RFC3339, "2025-01-01T12:00:00Z"),
					timex.MustParse(time.RFC3339, "2025-01-01T13:00:00Z"),
				),
				event.WithCosts(
					event.NewCosts(
						money.NewMoney(100, 0),
					),
				),
			),
			want: want{
				val: event.NewCosts(
					money.NewMoney(100, 0),
				),
			},
		},
		{
			name: "zero costs",
			e: event.NewEvent(
				event.NewID("1"),
				event.NewContent(
					"Title",
					"description",
				),
				event.NewLocation(
					"Main street, 13",
				),
				event.NewCreated(
					time.Time{},
					user.NewIdentity(
						user.NewID("123"),
					),
				),
				event.NewDates(
					timex.MustParse(time.RFC3339, "2025-01-01T12:00:00Z"),
					timex.MustParse(time.RFC3339, "2025-01-01T13:00:00Z"),
				),
				event.WithCosts(
					event.NewCosts(
						money.NewMoney(0, 0),
					),
				),
			),
			want: want{
				val: event.NewCosts(
					money.NewMoney(0, 0),
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.e.Costs(),
			)
		})
	}
}

func TestEvent_Hash(t *testing.T) {
	t.Parallel()
	type want struct {
		val string
	}
	tests := []struct {
		name string
		e    event.Event
		want want
	}{
		{
			name: "ok",
			e: event.NewEvent(
				event.NewID("1"),
				event.NewContent(
					"Title",
					"description",
				),
				event.NewLocation(
					"Main street, 13",
				),
				event.NewCreated(
					time.Time{},
					user.NewIdentity(
						user.NewID("123"),
					),
				),
				event.NewDates(
					timex.MustParse(time.RFC3339, "2025-01-01T12:00:00Z"),
					timex.MustParse(time.RFC3339, "2025-01-01T13:00:00Z"),
				),
				event.WithCosts(
					event.NewCosts(
						money.NewMoney(100, 0),
					),
				),
			),
			want: want{
				val: "f2b6c37565467bfb8e4ae05c4117ce61",
			},
		},
		{
			name: "default value",
			e:    event.Event{},
			want: want{
				val: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.e.Hash(),
			)
		})
	}
}
