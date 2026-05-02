package body_test

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/creation"
	impl "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/reviews/events/list/resp/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/rating"
	review "github.com/sitnikovik/ndbx/autograder/internal/app/review/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestMustParseBody(t *testing.T) {
	t.Parallel()
	type args struct {
		body io.ReadCloser
	}
	type want struct {
		val   impl.Body
		panic bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "ok",
			args: args{
				body: io.NopCloser(
					strings.NewReader(`{` +
						`"events": [
						{` +
						`"id": "56e2c0b3a2b4c1a5e6f7f8b3",` +
						`"event_id": "12e9c0b1a2b3c3d5e6f7a8b7",` +
						`"comment": "Great!",` +
						`"created_at": "2026-03-14T14:59:32+03:00",` +
						`"created_by": "65e9c0b1a2b3c4d5e6f7a8b9",` +
						`"rating": 5,` +
						`"updated_at": "2026-03-14T14:59:32+03:00"` +
						`}` +
						`],` +
						`"count": 1` +
						`}`,
					),
				),
			},
			want: want{
				val: impl.NewBody(
					[]review.Review{
						review.NewReview(
							"56e2c0b3a2b4c1a5e6f7f8b3",
							creation.NewStamp(
								creation.NewCreated(
									timex.MustRFC3339("2026-03-14T14:59:32+03:00"),
									user.NewIdentity(
										user.NewID("65e9c0b1a2b3c4d5e6f7a8b9"),
									),
								),
							),
							review.NewEvent(
								event.NewID("12e9c0b1a2b3c3d5e6f7a8b7"),
							),
							"Great!",
							rating.NewRating(5),
							review.WithUpdatedAt(
								timex.MustRFC3339("2026-03-14T14:59:32+03:00"),
							),
						),
					},
					1,
				),
				panic: false,
			},
		},
		{
			name: "empty events but count is 1",
			args: args{
				body: io.NopCloser(
					strings.NewReader(`{` +
						`"events": [],` +
						`"count": 1` +
						`}`,
					),
				),
			},
			want: want{
				val: impl.NewBody(
					nil,
					1,
				),
				panic: false,
			},
		},
		{
			name: "invalid json",
			args: args{
				body: io.NopCloser(
					strings.NewReader(`not json`),
				),
			},
			want: want{
				val: impl.NewBody(
					nil,
					1,
				),
				panic: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = impl.MustParseBody(tt.args.body)
				})
				return
			}
			assert.Equal(
				t,
				tt.want.val,
				impl.MustParseBody(tt.args.body),
			)
		})
	}
}
