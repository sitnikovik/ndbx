package doc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
)

func TestDocuments_IDs(t *testing.T) {
	t.Parallel()
	type want struct {
		val []string
	}
	tests := []struct {
		name string
		dd   doc.Documents
		want want
	}{
		{
			name: "ok",
			dd: doc.NewDocuments(
				doc.NewDocument(
					"id1",
					doc.NewKV("key", "value"),
				),
				doc.NewDocument(
					"id2",
					doc.NewKV("key", "value"),
				),
			),
			want: want{
				val: []string{"id1", "id2"},
			},
		},
		{
			name: "empty list",
			dd:   doc.NewDocuments(),
			want: want{
				val: nil,
			},
		},
		{
			name: "nil list",
			dd:   nil,
			want: want{
				val: nil,
			},
		},
		{
			name: "single document",
			dd: doc.NewDocuments(
				doc.NewDocument(
					"id",
					doc.NewKV("key", "value"),
				),
			),
			want: want{
				val: []string{"id"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.dd.IDs(),
			)
		})
	}
}

func TestDocuments_First(t *testing.T) {
	t.Parallel()
	type want struct {
		val doc.Document
	}
	tests := []struct {
		name string
		dd   doc.Documents
		want want
	}{
		{
			name: "ok",
			dd: doc.NewDocuments(
				doc.NewDocument(
					"id1",
					doc.NewKV("key", "value"),
				),
				doc.NewDocument(
					"id2",
					doc.NewKV("key", "value"),
				),
			),
			want: want{
				val: doc.NewDocument(
					"id1",
					doc.NewKV("key", "value"),
				),
			},
		},
		{
			name: "empty list",
			dd:   doc.NewDocuments(),
			want: want{
				val: doc.Document{},
			},
		},
		{
			name: "nil list",
			dd:   nil,
			want: want{
				val: doc.Document{},
			},
		},
		{
			name: "single document",
			dd: doc.NewDocuments(
				doc.NewDocument(
					"id",
					doc.NewKV("key", "value"),
				),
			),
			want: want{
				val: doc.NewDocument(
					"id",
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
				tt.dd.First(),
			)
		})
	}
}

func TestDocuments_Last(t *testing.T) {
	t.Parallel()
	type want struct {
		val doc.Document
	}
	tests := []struct {
		name string
		dd   doc.Documents
		want want
	}{
		{
			name: "ok",
			dd: doc.NewDocuments(
				doc.NewDocument(
					"id1",
					doc.NewKV("key", "value"),
				),
				doc.NewDocument(
					"id2",
					doc.NewKV("key", "value"),
				),
			),
			want: want{
				val: doc.NewDocument(
					"id2",
					doc.NewKV("key", "value"),
				),
			},
		},
		{
			name: "empty list",
			dd:   doc.NewDocuments(),
			want: want{
				val: doc.Document{},
			},
		},
		{
			name: "nil list",
			dd:   nil,
			want: want{
				val: doc.Document{},
			},
		},
		{
			name: "single document",
			dd: doc.NewDocuments(
				doc.NewDocument(
					"id",
					doc.NewKV("key", "value"),
				),
			),
			want: want{
				val: doc.NewDocument(
					"id",
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
				tt.dd.Last(),
			)
		})
	}
}
