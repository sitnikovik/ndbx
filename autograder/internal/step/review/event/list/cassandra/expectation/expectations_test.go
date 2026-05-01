package expectation_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/sitnikovik/ndbx/autograder/internal/step/review/event/list/cassandra/expectation"
)

func TestExpectations_Count(t *testing.T) {
	t.Parallel()
	type want struct {
		value int
	}
	tests := []struct {
		name string
		e    impl.Expectations
		opt  impl.Option
		want want
	}{
		{
			name: "ok",
			e: impl.NewExpectations(
				impl.WithCount(1),
			),
			want: want{
				value: 1,
			},
		},
		{
			name: "zero set",
			e: impl.NewExpectations(
				impl.WithCount(0),
			),
			want: want{
				value: 0,
			},
		},
		{
			name: "negative set",
			e: impl.NewExpectations(
				impl.WithCount(-1),
			),
			want: want{
				value: -1,
			},
		},
		{
			name: "default value",
			e:    impl.Expectations{},
			want: want{
				value: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.e.Count(),
			)
		})
	}
}
