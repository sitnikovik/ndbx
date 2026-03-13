package doc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
)

func TestKVs_ToBsonD(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		KVs  doc.KVs
		want bson.D
	}{
		{
			name: "ok",
			KVs: doc.NewKVs(
				doc.NewKV("key1", "value1"),
				doc.NewKV("key2", "value2"),
			),
			want: bson.D{
				{Key: "key1", Value: "value1"},
				{Key: "key2", Value: "value2"},
			},
		},
		{
			name: "empty slice",
			KVs:  doc.NewKVs(),
			want: nil,
		},
		{
			name: "nil slice",
			KVs:  nil,
			want: nil,
		},
		{
			name: "single KV",
			KVs: doc.NewKVs(
				doc.NewKV("key", "value"),
			),
			want: bson.D{
				{Key: "key", Value: "value"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want,
				tt.KVs.ToBsonD(),
			)
		})
	}
}

func TestKVs_First(t *testing.T) {
	t.Parallel()
	type want struct {
		val doc.KV
	}
	tests := []struct {
		name string
		kk   doc.KVs
		want want
	}{
		{
			name: "ok",
			kk: doc.NewKVs(
				doc.NewKV("key1", "value1"),
				doc.NewKV("key2", "value2"),
			),
			want: want{
				val: doc.NewKV("key1", "value1"),
			},
		},
		{
			name: "empty list",
			kk:   doc.NewKVs(),
			want: want{
				val: doc.KV{},
			},
		},
		{
			name: "nil list",
			kk:   nil,
			want: want{
				val: doc.KV{},
			},
		},
		{
			name: "single KV",
			kk: doc.NewKVs(
				doc.NewKV("key", "value"),
			),
			want: want{
				val: doc.NewKV("key", "value"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.kk.First(),
			)
		})
	}
}

func TestKVs_Last(t *testing.T) {
	t.Parallel()
	type want struct {
		val doc.KV
	}
	tests := []struct {
		name string
		kk   doc.KVs
		want want
	}{
		{
			name: "ok",
			kk: doc.NewKVs(
				doc.NewKV("key1", "value1"),
				doc.NewKV("key2", "value2"),
			),
			want: want{
				val: doc.NewKV("key2", "value2"),
			},
		},
		{
			name: "empty list",
			kk:   doc.NewKVs(),
			want: want{
				val: doc.KV{},
			},
		},
		{
			name: "nil list",
			kk:   nil,
			want: want{
				val: doc.KV{},
			},
		},
		{
			name: "single KV",
			kk: doc.NewKVs(
				doc.NewKV("key", "value"),
			),
			want: want{
				val: doc.NewKV("key", "value"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.kk.Last(),
			)
		})
	}
}

func TestKVs_Has(t *testing.T) {
	t.Parallel()
	type args struct {
		key string
	}
	type want struct {
		ok bool
	}
	tests := []struct {
		name string
		kk   doc.KVs
		args args
		want want
	}{
		{
			name: "key exists",
			kk: doc.NewKVs(
				doc.NewKV("key1", "value1"),
				doc.NewKV("key2", "value2"),
			),
			args: args{
				key: "key1",
			},
			want: want{
				ok: true,
			},
		},
		{
			name: "key does not exist",
			kk: doc.NewKVs(
				doc.NewKV("key1", "value1"),
				doc.NewKV("key2", "value2"),
			),
			args: args{
				key: "key3",
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "empty list",
			kk:   doc.NewKVs(),
			args: args{
				key: "key",
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "nil list",
			kk:   nil,
			args: args{
				key: "key",
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "multiple KVs with same key",
			kk: doc.NewKVs(
				doc.NewKV("key", "value1"),
				doc.NewKV("key", "value2"),
			),
			args: args{
				key: "key",
			},
			want: want{
				ok: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.kk.Has(tt.args.key)
			if tt.want.ok {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}
