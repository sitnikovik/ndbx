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
