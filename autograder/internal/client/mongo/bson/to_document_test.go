package bson_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"

	bsoni "github.com/sitnikovik/ndbx/autograder/internal/client/mongo/bson"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
)

func TestMList_ToDocuments(t *testing.T) {
	t.Parallel()
	type want struct {
		val   doc.Documents
		panic bool
	}
	tests := []struct {
		name string
		m    bsoni.MList
		want want
	}{
		{
			name: "ok",
			m: bsoni.NewBsonMList(
				bson.M{
					"_id": func() bson.ObjectID {
						id, err := bson.ObjectIDFromHex("000000000000000000000000")
						require.NoError(t, err)
						return id
					}(),
					"key": "value",
				},
				bson.M{
					"_id": func() bson.ObjectID {
						id, err := bson.ObjectIDFromHex("000000000000000000000001")
						require.NoError(t, err)
						return id
					}(),
					"key": "value",
				},
			),
			want: want{
				val: doc.NewDocuments(
					doc.NewDocument(
						"000000000000000000000000",
						doc.NewKV("key", "value"),
					),
					doc.NewDocument(
						"000000000000000000000001",
						doc.NewKV("key", "value"),
					),
				),
				panic: false,
			},
		},
		{
			name: "empty list",
			m:    bsoni.NewBsonMList(),
			want: want{
				val:   nil,
				panic: false,
			},
		},
		{
			name: "nil list",
			m:    nil,
			want: want{
				val:   nil,
				panic: false,
			},
		},
		{
			name: "invalid _id",
			m: bsoni.NewBsonMList(
				bson.M{
					"_id": "invalid_id",
					"key": "value",
				},
			),
			want: want{
				val:   nil,
				panic: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = tt.m.ToDocuments()
				})
				return
			}
			assert.Equal(
				t,
				tt.want.val,
				tt.m.ToDocuments(),
			)
		})
	}
}
