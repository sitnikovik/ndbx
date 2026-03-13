package bson_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"

	bsoni "github.com/sitnikovik/ndbx/autograder/internal/client/mongo/bson"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
)

func TestM_ToKV(t *testing.T) {
	t.Parallel()
	type want struct {
		val doc.KVs
	}
	tests := []struct {
		name string
		m    bsoni.M
		want want
	}{
		{
			name: "ok",
			m: bsoni.NewBsonM(
				bson.M{
					"_id": func() bson.ObjectID {
						id, err := bson.ObjectIDFromHex("000000000000000000000000")
						require.NoError(t, err)
						return id
					}(),
					"key1": "value1",
				},
			),
			want: want{
				val: doc.NewKVs(
					doc.NewKV("key1", "value1"),
				),
			},
		},
		{
			name: "empty bson",
			m:    bsoni.NewBsonM(bson.M{}),
			want: want{
				val: nil,
			},
		},
		{
			name: "bson with only _id",
			m: bsoni.NewBsonM(
				bson.M{
					"_id": func() bson.ObjectID {
						id, err := bson.ObjectIDFromHex("000000000000000000000000")
						require.NoError(t, err)
						return id
					}(),
				},
			),
			want: want{
				val: doc.KVs{},
			},
		},
		{
			name: "bson with nested bson.D containing _id ObjectID",
			m: bsoni.NewBsonM(
				bson.M{
					"_id": func() bson.ObjectID {
						id, err := bson.ObjectIDFromHex("000000000000000000000001")
						require.NoError(t, err)
						return id
					}(),
					"created_by": bson.D{
						{Key: "_id", Value: func() bson.ObjectID {
							id, err := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
							require.NoError(t, err)
							return id
						}()},
					},
					"title": "test event",
				},
			),
			want: want{
				val: doc.NewKVs(
					doc.NewKV("created_by", "507f1f77bcf86cd799439011"),
					doc.NewKV("title", "test event"),
				),
			},
		},
		{
			name: "bson with nested bson.D containing _id string",
			m: bsoni.NewBsonM(
				bson.M{
					"_id": func() bson.ObjectID {
						id, err := bson.ObjectIDFromHex("000000000000000000000001")
						require.NoError(t, err)
						return id
					}(),
					"user_ref": bson.D{
						{Key: "_id", Value: "user123"},
					},
				},
			),
			want: want{
				val: doc.NewKVs(
					doc.NewKV("user_ref", "user123"),
				),
			},
		},
		{
			name: "bson with nested bson.D with multiple fields (not normalized)",
			m: bsoni.NewBsonM(
				bson.M{
					"_id": func() bson.ObjectID {
						id, err := bson.ObjectIDFromHex("000000000000000000000001")
						require.NoError(t, err)
						return id
					}(),
					"metadata": bson.D{
						{Key: "_id", Value: "id123"},
						{Key: "name", Value: "test"},
					},
				},
			),
			want: want{
				val: doc.NewKVs(
					doc.NewKV("metadata", bson.D{
						{Key: "_id", Value: "id123"},
						{Key: "name", Value: "test"},
					}),
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
				tt.m.ToKV(),
			)
		})
	}
}
