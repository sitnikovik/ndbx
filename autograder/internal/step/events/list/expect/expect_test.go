package expect_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event/reaction"
	"github.com/sitnikovik/ndbx/autograder/internal/app/reaction/count"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/events/list/expect"
	eventfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/event"
	userfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestExpectations_EventsRequired(t *testing.T) {
	t.Parallel()
	type want struct {
		value bool
	}
	tests := []struct {
		name string
		e    impl.Expectations
		want want
	}{
		{
			name: "with no events",
			e: impl.NewExpectations(
				impl.WithNoEvents(),
			),
			want: want{
				value: true,
			},
		},
		{
			name: "no events specified",
			e: impl.NewExpectations(
				impl.WithEvents(),
			),
			want: want{
				value: false,
			},
		},
		{
			name: "events specified",
			e: impl.NewExpectations(
				impl.WithEvents(
					eventfx.NewBirthdayParty(
						event.NewDates(
							timex.MustRFC3339("2026-03-31T15:00:00Z"),
							timex.MustRFC3339("2026-03-31T23:00:00Z"),
						),
						timex.MustRFC3339("2026-03-14T12:31:00Z"),
						userfx.NewSamwiseGamgee(),
					),
				),
			),
			want: want{
				value: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.e.EventsRequired()
			if tt.want.value {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}

func TestExpectations_Events(t *testing.T) {
	t.Parallel()
	type want struct {
		value []event.Event
	}
	tests := []struct {
		name string
		e    impl.Expectations
		want want
	}{
		{
			name: "with no events",
			e: impl.NewExpectations(
				impl.WithNoEvents(),
			),
			want: want{
				value: []event.Event{},
			},
		},
		{
			name: "no events specified",
			e: impl.NewExpectations(
				impl.WithEvents(),
			),
			want: want{
				value: nil,
			},
		},
		{
			name: "events specified",
			e: impl.NewExpectations(
				impl.WithEvents(
					eventfx.NewBirthdayParty(
						event.NewDates(
							timex.MustRFC3339("2026-03-31T15:00:00Z"),
							timex.MustRFC3339("2026-03-31T23:00:00Z"),
						),
						timex.MustRFC3339("2026-03-14T12:31:00Z"),
						userfx.NewSamwiseGamgee(),
					),
				),
			),
			want: want{
				value: []event.Event{
					eventfx.NewBirthdayParty(
						event.NewDates(
							timex.MustRFC3339("2026-03-31T15:00:00Z"),
							timex.MustRFC3339("2026-03-31T23:00:00Z"),
						),
						timex.MustRFC3339("2026-03-14T12:31:00Z"),
						userfx.NewSamwiseGamgee(),
					),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.e.Events(),
			)
		})
	}
}

func TestExpectations_ReactionsRequired(t *testing.T) {
	t.Parallel()
	type want struct {
		value bool
	}
	tests := []struct {
		name string
		e    impl.Expectations
		want want
	}{
		{
			name: "with no reactions",
			e: impl.NewExpectations(
				impl.WithNoReactions(),
			),
			want: want{
				value: true,
			},
		},
		{
			name: "no reactions specified",
			e: impl.NewExpectations(
				impl.WithReactions(),
			),
			want: want{
				value: false,
			},
		},
		{
			name: "reactions specified",
			e: impl.NewExpectations(
				impl.WithReactions(
					reaction.NewReactions(
						reaction.WithCounts(
							count.NewCounts(
								count.WithLikes(24),
								count.WithDislikes(3),
							),
						),
					),
				),
			),
			want: want{
				value: true,
			},
		},
		{
			name: "reactions specified in the events",
			e: impl.NewExpectations(
				impl.WithEvents(
					event.NewEvent(
						event.NewID("1"),
						event.NewContent(
							"test title",
							"test description",
						),
						event.NewLocation("test location"),
						event.NewCreated(
							timex.MustRFC3339("2024-01-01T00:00:00Z"),
							user.NewIdentity("test_user"),
						),
						event.NewDates(
							timex.MustRFC3339("2024-01-01T01:00:00Z"),
							timex.MustRFC3339("2024-01-01T02:00:00Z"),
						),
						event.WithLikes(24),
						event.WithDislikes(3),
					),
				),
			),
			want: want{
				value: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.e.ReactionsRequired()
			if tt.want.value {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}

func TestExpectations_Reactions(t *testing.T) {
	t.Parallel()
	type want struct {
		value []reaction.Reactions
	}
	tests := []struct {
		name string
		e    impl.Expectations
		want want
	}{
		{
			name: "with no reactions",
			e: impl.NewExpectations(
				impl.WithNoReactions(),
			),
			want: want{
				value: []reaction.Reactions{},
			},
		},
		{
			name: "no reactions specified",
			e: impl.NewExpectations(
				impl.WithReactions(),
			),
			want: want{
				value: nil,
			},
		},
		{
			name: "reactions specified",
			e: impl.NewExpectations(
				impl.WithReactions(
					reaction.NewReactions(
						reaction.WithCounts(
							count.NewCounts(
								count.WithLikes(24),
								count.WithDislikes(3),
							),
						),
					),
				),
			),
			want: want{
				value: []reaction.Reactions{
					reaction.NewReactions(
						reaction.WithCounts(
							count.NewCounts(
								count.WithLikes(24),
								count.WithDislikes(3),
							),
						),
					),
				},
			},
		},
		{
			name: "reactions specified in the events",
			e: impl.NewExpectations(
				impl.WithEvents(
					event.NewEvent(
						event.NewID("1"),
						event.NewContent(
							"test title",
							"test description",
						),
						event.NewLocation("test location"),
						event.NewCreated(
							timex.MustRFC3339("2024-01-01T00:00:00Z"),
							user.NewIdentity("test_user"),
						),
						event.NewDates(
							timex.MustRFC3339("2024-01-01T01:00:00Z"),
							timex.MustRFC3339("2024-01-01T02:00:00Z"),
						),
						event.WithLikes(24),
						event.WithDislikes(3),
					),
				),
			),
			want: want{
				value: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.e.Reactions(),
			)
		})
	}
}
