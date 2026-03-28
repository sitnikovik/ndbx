package shard_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/shard"
)

func TestParseShards(t *testing.T) {
	t.Parallel()
	type args struct {
		mm map[string]map[string]any
	}
	type want struct {
		val         shard.Shards
		errContains string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "ok",
			args: args{
				mm: map[string]map[string]any{
					"rs0": {
						"ns":             "ndbx.events",
						"count":          int64(512),
						"size":           int64(102400),
						"avgObjSize":     int64(200),
						"storageSize":    int64(65536),
						"capped":         false,
						"nindexes":       int32(3),
						"totalIndexSize": int64(49152),
						"ok":             float64(1),
					},
				},
			},
			want: want{
				val: shard.NewShards(
					shard.NewShard(
						"rs0",
						shard.WithCount(512),
						shard.WithOk(true),
					),
				),
				errContains: "",
			},
		},
		{
			name: "nil map",
			args: args{
				mm: nil,
			},
			want: want{
				val:         nil,
				errContains: "empty map",
			},
		},
		{
			name: "nil map in key",
			args: args{
				mm: map[string]map[string]any{
					"rs0": nil,
				},
			},
			want: want{
				val:         nil,
				errContains: "error for 'rs0': empty map",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := shard.ParseShards(tt.args.mm)
			assert.Equal(t, tt.want.val, got)
			if tt.want.errContains != "" {
				assert.ErrorContains(t, err, tt.want.errContains)
			}
		})
	}
}
