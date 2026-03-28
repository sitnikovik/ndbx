package event_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
)

func TestLocation_Address(t *testing.T) {
	t.Parallel()
	type want struct {
		val string
	}
	tests := []struct {
		name string
		l    event.Location
		want want
	}{
		{
			name: "ok",
			l: event.NewLocation(
				"City, Country, Street, 123",
			),
			want: want{
				val: "City, Country, Street, 123",
			},
		},
		{
			name: "empty address",
			l:    event.NewLocation(""),
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
				tt.l.Address(),
			)
		})
	}
}

func TestLocation_City(t *testing.T) {
	t.Parallel()
	type want struct {
		val string
	}
	tests := []struct {
		name string
		l    event.Location
		want want
	}{
		{
			name: "ok",
			l: event.NewLocation(
				"City, Country, Street, 123",
				event.WithCity("New York"),
			),
			want: want{
				val: "New York",
			},
		},
		{
			name: "empty city but full address",
			l: event.NewLocation(
				"City, Country, Street, 123",
			),
			want: want{
				val: "",
			},
		},
		{
			name: "default",
			l:    event.NewLocation(""),
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
				tt.l.City(),
			)
		})
	}
}

func TestLocation_Empty(t *testing.T) {
	t.Parallel()
	type want struct {
		val bool
	}
	tests := []struct {
		name string
		l    event.Location
		want want
	}{
		{
			name: "city and address",
			l: event.NewLocation(
				"City, Country, Street, 123",
				event.WithCity("New York"),
			),
			want: want{
				val: false,
			},
		},
		{
			name: "only city",
			l: event.NewLocation(
				"",
				event.WithCity("New York"),
			),
			want: want{
				val: false,
			},
		},
		{
			name: "only address",
			l: event.NewLocation(
				"City, Country, Street, 123",
			),
			want: want{
				val: false,
			},
		},
		{
			name: "empty city and address",
			l: event.NewLocation(
				"",
				event.WithCity(""),
			),
			want: want{
				val: true,
			},
		},
		{
			name: "default value",
			l:    event.Location{},
			want: want{
				val: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.l.Empty()
			if tt.want.val {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}
