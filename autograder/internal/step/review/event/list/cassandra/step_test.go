package cassandra_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/review/event/list/cassandra"
	"github.com/sitnikovik/ndbx/autograder/internal/step/review/event/list/cassandra/expectation"
	cassandrafk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/cassandra/client"
	eventfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/event"
	userfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestStep_Name(t *testing.T) {
	t.Parallel()
	type want struct {
		value string
		panic bool
	}
	tests := []struct {
		name string
		s    *impl.Step
		want want
	}{
		{
			name: "ok",
			s: impl.NewStep(
				step.NewDesc(
					"test title",
					"test desc",
				),
				cassandrafk.NewClient(),
				eventfx.NewBirthdayParty(
					event.NewDates(
						timex.MustRFC3339("2024-02-05T11:00:00Z"),
						timex.MustRFC3339("2024-02-05T18:00:00Z"),
					),
					timex.MustRFC3339("2024-02-01T11:00:00Z"),
					userfx.NewJohnDoe(),
				),
				expectation.NewExpectations(
					expectation.WithCount(1),
				),
			),
			want: want{
				value: "test title",
				panic: false,
			},
		},
		{
			name: "empty title",
			s: impl.NewStep(
				step.NewDesc(
					"",
					"test desc",
				),
				cassandrafk.NewClient(),
				eventfx.NewBirthdayParty(
					event.NewDates(
						timex.MustRFC3339("2024-02-05T11:00:00Z"),
						timex.MustRFC3339("2024-02-05T18:00:00Z"),
					),
					timex.MustRFC3339("2024-02-01T11:00:00Z"),
					userfx.NewJohnDoe(),
				),
				expectation.NewExpectations(
					expectation.WithCount(1),
				),
			),
			want: want{
				value: "",
				panic: false,
			},
		},
		{
			name: "default value",
			s:    nil,
			want: want{
				value: "",
				panic: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = tt.s.Name()
				})
				return
			}
			assert.Equal(
				t,
				tt.want.value,
				tt.s.Name(),
			)
		})
	}
}

func TestStep_Description(t *testing.T) {
	t.Parallel()
	type want struct {
		value string
		panic bool
	}
	tests := []struct {
		name string
		s    *impl.Step
		want want
	}{
		{
			name: "ok",
			s: impl.NewStep(
				step.NewDesc(
					"test title",
					"test desc",
				),
				cassandrafk.NewClient(),
				eventfx.NewBirthdayParty(
					event.NewDates(
						timex.MustRFC3339("2024-02-05T11:00:00Z"),
						timex.MustRFC3339("2024-02-05T18:00:00Z"),
					),
					timex.MustRFC3339("2024-02-01T11:00:00Z"),
					userfx.NewJohnDoe(),
				),
				expectation.NewExpectations(
					expectation.WithCount(1),
				),
			),
			want: want{
				value: "test desc",
				panic: false,
			},
		},
		{
			name: "empty title",
			s: impl.NewStep(
				step.NewDesc(
					"test title",
					"",
				),
				cassandrafk.NewClient(),
				eventfx.NewBirthdayParty(
					event.NewDates(
						timex.MustRFC3339("2024-02-05T11:00:00Z"),
						timex.MustRFC3339("2024-02-05T18:00:00Z"),
					),
					timex.MustRFC3339("2024-02-01T11:00:00Z"),
					userfx.NewJohnDoe(),
				),
				expectation.NewExpectations(
					expectation.WithCount(1),
				),
			),
			want: want{
				value: "",
				panic: false,
			},
		},
		{
			name: "default value",
			s:    nil,
			want: want{
				value: "",
				panic: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = tt.s.Name()
				})
				return
			}
			assert.Equal(
				t,
				tt.want.value,
				tt.s.Description(),
			)
		})
	}
}
