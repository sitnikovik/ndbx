package shard_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/shard"
)

func TestParseShard(t *testing.T) {
	t.Parallel()
	type args struct {
		id string
		m  map[string]any
	}
	type want struct {
		val         shard.Shard
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
				id: "rs0",
				m: map[string]any{
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
			want: want{
				val: shard.NewShard(
					"rs0",
					shard.WithCount(512),
					shard.WithOk(true),
				),
				errContains: "",
			},
		},
		{
			name: "empty id",
			args: args{
				id: "",
				m:  nil,
			},
			want: want{
				val:         shard.Shard{},
				errContains: "empty id",
			},
		},
		{
			name: "nil map",
			args: args{
				id: "rs0",
				m:  nil,
			},
			want: want{
				val:         shard.Shard{},
				errContains: "empty map",
			},
		},
		{
			name: "count is int",
			args: args{
				id: "rs0",
				m: map[string]any{
					"ns":             "ndbx.events",
					"count":          int(512),
					"size":           int64(102400),
					"avgObjSize":     int64(200),
					"storageSize":    int64(65536),
					"capped":         false,
					"nindexes":       int32(3),
					"totalIndexSize": int64(49152),
					"ok":             float64(1),
				},
			},
			want: want{
				val:         shard.Shard{},
				errContains: "expect 'count' to be type of",
			},
		},
		{
			name: "count is int32",
			args: args{
				id: "rs0",
				m: map[string]any{
					"ns":             "ndbx.events",
					"count":          int32(512),
					"size":           int64(102400),
					"avgObjSize":     int64(200),
					"storageSize":    int64(65536),
					"capped":         false,
					"nindexes":       int32(3),
					"totalIndexSize": int64(49152),
					"ok":             float64(1),
				},
			},
			want: want{
				val: shard.NewShard(
					"rs0",
					shard.WithCount(512),
					shard.WithOk(true),
				),
				errContains: "",
			},
		},
		{
			name: "ok is bool",
			args: args{
				id: "rs0",
				m: map[string]any{
					"ns":             "ndbx.events",
					"count":          int64(512),
					"size":           int64(102400),
					"avgObjSize":     int64(200),
					"storageSize":    int64(65536),
					"capped":         false,
					"nindexes":       int32(3),
					"totalIndexSize": int64(49152),
					"ok":             false,
				},
			},
			want: want{
				val:         shard.Shard{},
				errContains: "expect 'ok' to be type of",
			},
		},
		{
			name: "ok is uint8",
			args: args{
				id: "rs0",
				m: map[string]any{
					"ns":             "ndbx.events",
					"count":          int64(512),
					"size":           int64(102400),
					"avgObjSize":     int64(200),
					"storageSize":    int64(65536),
					"capped":         false,
					"nindexes":       int32(3),
					"totalIndexSize": int64(49152),
					"ok":             uint8(1),
				},
			},
			want: want{
				val:         shard.Shard{},
				errContains: "expect 'ok' to be type of",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := shard.ParseShard(tt.args.id, tt.args.m)
			assert.Equal(t, tt.want.val, got)
			if tt.want.errContains != "" {
				assert.ErrorContains(t, err, tt.want.errContains)
			}
		})
	}
}

func TestShard_ID(t *testing.T) {
	t.Parallel()
	type want struct {
		val string
	}
	tests := []struct {
		name string
		s    shard.Shard
		want want
	}{
		{
			name: "ok",
			s:    shard.NewShard("rs0"),
			want: want{
				val: "rs0",
			},
		},
		{
			name: "empty",
			s:    shard.NewShard(""),
			want: want{
				val: "",
			},
		},
		{
			name: "space",
			s:    shard.NewShard(" "),
			want: want{
				val: " ",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.s.ID(),
			)
		})
	}
}

func TestShard_Count(t *testing.T) {
	t.Parallel()
	type want struct {
		val int
	}
	tests := []struct {
		name string
		s    shard.Shard
		want want
	}{
		{
			name: "positive",
			s: shard.NewShard(
				"rs0",
				shard.WithCount(1),
			),
			want: want{
				val: 1,
			},
		},
		{
			name: "zero",
			s: shard.NewShard(
				"rs0",
				shard.WithCount(0),
			),
			want: want{
				val: 0,
			},
		},
		{
			name: "negative",
			s: shard.NewShard(
				"rs0",
				shard.WithCount(-1),
			),
			want: want{
				val: -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.s.Count(),
			)
		})
	}
}

func TestShard_Ok(t *testing.T) {
	t.Parallel()
	type want struct {
		val bool
	}
	tests := []struct {
		name string
		s    shard.Shard
		want want
	}{
		{
			name: "true",
			s: shard.NewShard(
				"rs0",
				shard.WithOk(true),
			),
			want: want{
				val: true,
			},
		},
		{
			name: "false",
			s: shard.NewShard(
				"rs0",
				shard.WithOk(false),
			),
			want: want{
				val: false,
			},
		},
		{
			name: "all empty but ok",
			s: shard.NewShard(
				"",
				shard.WithCount(0),
				shard.WithOk(true),
			),
			want: want{
				val: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.s.Ok(),
			)
		})
	}
}
