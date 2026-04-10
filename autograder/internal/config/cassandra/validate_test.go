package cassandra_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/client/cassandra/consistency"
	impl "github.com/sitnikovik/ndbx/autograder/internal/config/cassandra"
)

func TestConfig_Validate(t *testing.T) {
	t.Parallel()
	type want struct {
		errored bool
	}
	tests := []struct {
		name string
		c    impl.Config
		want want
	}{
		{
			name: "ok",
			c: impl.NewConfig(
				impl.NewConnection(
					[]string{"localhost"},
					9042,
				),
				impl.NewAuth("usr", "pwd"),
				impl.NewDatabase(
					"testkeyspace",
					consistency.Quorum,
				),
			),
			want: want{
				errored: false,
			},
		},
		{
			name: "ok without auth",
			c: impl.NewConfig(
				impl.NewConnection(
					[]string{"localhost"},
					9042,
				),
				impl.NewAuth("", ""),
				impl.NewDatabase(
					"testkeyspace",
					consistency.Quorum,
				),
			),
			want: want{
				errored: false,
			},
		},
		{
			name: "empty hosts",
			c: impl.NewConfig(
				impl.NewConnection(
					[]string{},
					9042,
				),
				impl.NewAuth("", ""),
				impl.NewDatabase(
					"testkeyspace",
					consistency.Quorum,
				),
			),
			want: want{
				errored: true,
			},
		},
		{
			name: "invalid port",
			c: impl.NewConfig(
				impl.NewConnection(
					[]string{"localhost"},
					0,
				),
				impl.NewAuth("usr", "pwd"),
				impl.NewDatabase(
					"testkeyspace",
					consistency.Quorum,
				),
			),
			want: want{
				errored: true,
			},
		},
		{
			name: "invalid port",
			c: impl.NewConfig(
				impl.NewConnection(
					[]string{"localhost"},
					99999999,
				),
				impl.NewAuth("usr", "pwd"),
				impl.NewDatabase(
					"testkeyspace",
					consistency.Quorum,
				),
			),
			want: want{
				errored: true,
			},
		},
		{
			name: "empty keyspace",
			c: impl.NewConfig(
				impl.NewConnection(
					[]string{"localhost"},
					9042,
				),
				impl.NewAuth("usr", "pwd"),
				impl.NewDatabase(
					"",
					consistency.Quorum,
				),
			),
			want: want{
				errored: true,
			},
		},
		{
			name: "unknown consistency level",
			c: impl.NewConfig(
				impl.NewConnection(
					[]string{"localhost"},
					9042,
				),
				impl.NewAuth("usr", "pwd"),
				impl.NewDatabase(
					"testkeyspace",
					consistency.Consistency(213),
				),
			),
			want: want{
				errored: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.c.Validate()
			if tt.want.errored {
				assert.Error(t, err)
			} else {
				assert.NoErrorf(
					t,
					err,
					"unexpected error: %v",
					err,
				)
			}
		})
	}
}
