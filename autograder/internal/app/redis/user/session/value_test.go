package session_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/redis/user/session"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

func TestValue_Validate(t *testing.T) {
	t.Parallel()
	type want struct {
		err error
	}
	tests := []struct {
		name string
		v    session.Value
		want want
	}{
		{
			name: "ok",
			v: session.NewValue(
				"1",
				map[string]string{
					session.CreatedAtField: "2024-06-01T00:00:00Z",
					session.UpdatedAtField: "2024-06-01T00:00:00Z",
					session.UserIDField:    "1",
				},
			),
			want: want{
				err: nil,
			},
		},
		{
			name: "missing sid",
			v: session.NewValue(
				"",
				map[string]string{
					session.UpdatedAtField: "2024-06-01T00:00:00Z",
					session.UserIDField:    "1",
				},
			),
			want: want{
				err: errs.ErrExpectationFailed,
			},
		},
		{
			name: "missing created_at field",
			v: session.NewValue(
				"1",
				map[string]string{
					session.UpdatedAtField: "2024-06-01T00:00:00Z",
					session.UserIDField:    "1",
				},
			),
			want: want{
				err: errs.ErrExpectationFailed,
			},
		},
		{
			name: "missing updated_at field",
			v: session.NewValue(
				"1",
				map[string]string{
					session.CreatedAtField: "2024-06-01T00:00:00Z",
					session.UserIDField:    "1",
				},
			),
			want: want{
				err: errs.ErrExpectationFailed,
			},
		},
		{
			name: "missing user_id field",
			v: session.NewValue(
				"1",
				map[string]string{
					session.CreatedAtField: "2024-06-01T00:00:00Z",
					session.UpdatedAtField: "2024-06-01T00:00:00Z",
				},
			),
			want: want{
				err: errs.ErrExpectationFailed,
			},
		},
		{
			name: "invalid created_at field",
			v: session.NewValue(
				"1",
				map[string]string{
					session.CreatedAtField: "invalid",
					session.UpdatedAtField: "2024-06-01T00:00:00Z",
					session.UserIDField:    "1",
				}),
			want: want{
				err: errs.ErrExpectationFailed,
			},
		},
		{
			name: "invalid updated_at field",
			v: session.NewValue(
				"1",
				map[string]string{
					session.CreatedAtField: "2024-06-01T00:00:00Z",
					session.UpdatedAtField: "invalid",
					session.UserIDField:    "1",
				}),
			want: want{
				err: errs.ErrExpectationFailed,
			},
		},
		{
			name: "updated_at is before created_at",
			v: session.NewValue(
				"1",
				map[string]string{
					session.CreatedAtField: "2024-06-01T01:00:00Z",
					session.UpdatedAtField: "2024-06-01T00:59:59Z",
					session.UserIDField:    "1",
				}),
			want: want{
				err: errs.ErrExpectationFailed,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.ErrorIs(
				t,
				tt.v.Validate(),
				tt.want.err,
			)
		})
	}
}
