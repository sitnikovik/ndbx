package doc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
)

func TestKV_ToBsonE(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		k    doc.KV
		want bson.E
	}{
		{
			name: "ok",
			k:    doc.NewKV("key", "value"),
			want: bson.E{Key: "key", Value: "value"},
		},
		{
			name: "empty key",
			k:    doc.NewKV("", "value"),
			want: bson.E{Key: "", Value: "value"},
		},
		{
			name: "nil value",
			k:    doc.NewKV("key", nil),
			want: bson.E{Key: "key", Value: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want,
				tt.k.ToBsonE(),
			)
		})
	}
}

func TestKV_Has(t *testing.T) {
	t.Parallel()
	type args struct {
		key string
	}
	type want struct {
		ok bool
	}
	tests := []struct {
		name string
		k    doc.KV
		args args
		want want
	}{
		{
			name: "exists",
			k:    doc.NewKV("key", "value"),
			args: args{
				key: "key",
			},
			want: want{
				ok: true,
			},
		},
		{
			name: "not exists",
			k:    doc.NewKV("key", "value"),
			args: args{
				key: "otherkey",
			},
			want: want{
				ok: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.k.Has(tt.args.key)
			if tt.want.ok {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}
