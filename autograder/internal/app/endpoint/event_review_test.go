package endpoint_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
)

func TestEndpoint_EventReview(t *testing.T) {
	t.Parallel()
	type args struct {
		eventID  string
		reviewID string
	}
	type want struct {
		val string
	}
	tests := []struct {
		name string
		e    endpoint.Endpoint
		args args
		want want
	}{

		{
			name: "ok",
			e:    endpoint.NewEndpoint("http://localhost"),
			args: args{
				eventID:  "event-id-123",
				reviewID: "review-id-567",
			},
			want: want{
				val: "http://localhost/events/event-id-123/reviews/review-id-567",
			},
		},
		{
			name: "empty event id",
			e:    endpoint.NewEndpoint("http://localhost"),
			args: args{
				eventID:  "",
				reviewID: "review-id-567",
			},
			want: want{
				val: "http://localhost/events//reviews/review-id-567",
			},
		},
		{
			name: "empty review id",
			e:    endpoint.NewEndpoint("http://localhost"),
			args: args{
				eventID:  "event-id-123",
				reviewID: "",
			},
			want: want{
				val: "http://localhost/events/event-id-123/reviews/",
			},
		},
		{
			name: "empty base URL",
			e:    endpoint.NewEndpoint(""),
			args: args{
				eventID:  "event-id-123",
				reviewID: "review-id-567",
			},
			want: want{
				val: "/events/event-id-123/reviews/review-id-567",
			},
		},
		{
			name: "empty base URL and ids",
			e:    endpoint.NewEndpoint(""),
			args: args{
				eventID:  "",
				reviewID: "",
			},
			want: want{
				val: "/events//reviews/",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.e.EventReview(
					tt.args.eventID,
					tt.args.reviewID,
				),
			)
		})
	}
}
