package bson_test

import (
	"testing"

	"go.mongodb.org/mongo-driver/v2/bson"

	bsoni "github.com/sitnikovik/ndbx/autograder/internal/client/mongo/bson"
)

func TestM_ToIndex_BsonD(t *testing.T) {
	tests := []struct {
		name     string
		m        bsoni.M
		wantKeys []string
		wantUniq bool
	}{
		{
			name: "simple index with bson.D key",
			m: bsoni.M{
				"v":    2,
				"key":  bson.D{{Key: "username", Value: 1}},
				"name": "username_1",
			},
			wantKeys: []string{"username"},
			wantUniq: false,
		},
		{
			name: "unique index with bson.D key",
			m: bsoni.M{
				"v":      2,
				"key":    bson.D{{Key: "username", Value: 1}},
				"name":   "username_1",
				"unique": true,
			},
			wantKeys: []string{"username"},
			wantUniq: true,
		},
		{
			name: "compound index with bson.D key",
			m: bsoni.M{
				"v":    2,
				"key":  bson.D{{Key: "field1", Value: 1}, {Key: "field2", Value: -1}},
				"name": "field1_1_field2_-1",
			},
			wantKeys: []string{"field1", "field2"},
			wantUniq: false,
		},
		{
			name: "simple index with map key",
			m: bsoni.M{
				"v":    2,
				"key":  bsoni.M{"username": 1},
				"name": "username_1",
			},
			wantKeys: []string{"username"},
			wantUniq: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.ToIndex()
			if len(got) != 1 {
				t.Fatalf("expected 1 index, got %d", len(got))
			}
			idx := got[0]

			// Check if index has all expected keys
			if !idx.HasAllFor(tt.wantKeys...) {
				t.Errorf("index doesn't have all expected keys %v", tt.wantKeys)
			}

			// Check unique flag
			if idx.Unique() != tt.wantUniq {
				t.Errorf("expected unique=%v, got %v", tt.wantUniq, idx.Unique())
			}
		})
	}
}

func TestMList_ToIndexes(t *testing.T) {
	ml := bsoni.MList{
		bsoni.M{
			"v":    2,
			"key":  bson.D{{Key: "_id", Value: 1}},
			"name": "_id_",
		},
		bsoni.M{
			"v":      2,
			"key":    bson.D{{Key: "username", Value: 1}},
			"name":   "username_1",
			"unique": true,
		},
	}

	indexes := ml.ToIndexes()
	if len(indexes) != 2 {
		t.Fatalf("expected 2 indexes, got %d", len(indexes))
	}

	// Check if we can find username index
	usernameIdx := indexes.For("username")
	if usernameIdx.Empty() {
		t.Error("username index not found")
	}
	if !usernameIdx.Unique() {
		t.Error("username index should be unique")
	}
}
