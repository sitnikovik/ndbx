package doc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
)

func TestDocument_KVs(t *testing.T) {
	t.Parallel()
	type want struct {
		val doc.KVs
	}
	tests := []struct {
		name string
		d    doc.Document
		want want
	}{
		{
			name: "ok",
			d: doc.NewDocument(
				"id",
				doc.NewKV("key1", "value1"),
				doc.NewKV("key2", "value2"),
			),
			want: want{
				val: doc.NewKVs(
					doc.NewKV("key1", "value1"),
					doc.NewKV("key2", "value2"),
				),
			},
		},
		{
			name: "empty KVs",
			d:    doc.NewDocument("id"),
			want: want{
				val: nil,
			},
		},
		{
			name: "single KV",
			d: doc.NewDocument(
				"id",
				doc.NewKV("key", "value"),
			),
			want: want{
				val: doc.NewKVs(
					doc.NewKV("key", "value"),
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.d.KVs(),
			)
		})
	}
}

func TestDocument_ID(t *testing.T) {
	t.Parallel()
	type want struct {
		val string
	}
	tests := []struct {
		name string
		d    doc.Document
		want want
	}{
		{
			name: "ok",
			d: doc.NewDocument(
				"id",
				doc.NewKV("key", "value"),
			),
			want: want{
				val: "id",
			},
		},
		{
			name: "empty id",
			d: doc.NewDocument(
				"",
				doc.NewKV("key", "value"),
			),
			want: want{
				val: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.d.ID(),
			)
		})
	}
}
