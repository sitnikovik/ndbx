package event_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/creation"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/rating"
	impl "github.com/sitnikovik/ndbx/autograder/internal/app/review/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

var reviewFx = impl.NewReview(
	"123",
	creation.NewStamp(
		creation.NewCreated(
			timex.MustRFC3339("2025-03-11T11:00:00Z"),
			user.NewIdentity(user.NewID("69587u4y")),
		),
	),
	impl.NewEvent(
		event.NewID("123rews"),
	),
	"Great!",
	rating.Five,
	impl.WithUpdatedAt(
		timex.MustRFC3339("2025-03-14T13:00:00Z"),
	),
)

func TestReview_Created(t *testing.T) {
	t.Parallel()
	type want struct {
		value creation.Created
	}
	tests := []struct {
		name string
		r    impl.Review
		want want
	}{
		{
			name: "ok",
			r:    reviewFx,
			want: want{
				value: creation.NewCreated(
					timex.MustRFC3339("2025-03-11T11:00:00Z"),
					user.NewIdentity(user.NewID("69587u4y")),
				),
			},
		},
		{
			name: "default value",
			r:    impl.Review{},
			want: want{
				value: creation.Created{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.r.Created(),
			)
		})
	}
}

func TestReview_UpdatedAt(t *testing.T) {
	t.Parallel()
	type want struct {
		value time.Time
	}
	tests := []struct {
		name string
		r    impl.Review
		want want
	}{
		{
			name: "ok",
			r:    reviewFx,
			want: want{
				value: timex.MustRFC3339("2025-03-14T13:00:00Z"),
			},
		},
		{
			name: "default value",
			r:    impl.Review{},
			want: want{
				value: time.Time{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.r.UpdatedAt(),
			)
		})
	}
}

func TestReview_Comment(t *testing.T) {
	t.Parallel()
	type want struct {
		value string
	}
	tests := []struct {
		name string
		r    impl.Review
		want want
	}{
		{
			name: "ok",
			r:    reviewFx,
			want: want{
				value: "Great!",
			},
		},
		{
			name: "default value",
			r:    impl.Review{},
			want: want{
				value: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.r.Comment(),
			)
		})
	}
}

func TestReview_Event(t *testing.T) {
	t.Parallel()
	type want struct {
		value impl.Event
	}
	tests := []struct {
		name string
		r    impl.Review
		want want
	}{
		{
			name: "ok",
			r:    reviewFx,
			want: want{
				value: impl.NewEvent(
					event.NewID("123rews"),
				),
			},
		},
		{
			name: "default value",
			r:    impl.Review{},
			want: want{
				value: impl.Event{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.r.Event(),
			)
		})
	}
}

func TestReview_ID(t *testing.T) {
	t.Parallel()
	type want struct {
		value string
	}
	tests := []struct {
		name string
		r    impl.Review
		want want
	}{
		{
			name: "ok",
			r:    reviewFx,
			want: want{
				value: "123",
			},
		},
		{
			name: "default value",
			r:    impl.Review{},
			want: want{
				value: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.r.ID(),
			)
		})
	}
}

func TestReview_Rating(t *testing.T) {
	t.Parallel()
	type want struct {
		value rating.Rating
	}
	tests := []struct {
		name string
		r    impl.Review
		want want
	}{
		{
			name: "ok",
			r:    reviewFx,
			want: want{
				value: rating.Five,
			},
		},
		{
			name: "default value",
			r:    impl.Review{},
			want: want{
				value: rating.None,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.r.Rating(),
			)
		})
	}
}
