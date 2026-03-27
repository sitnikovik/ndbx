package bson

import (
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// ToMap converts bson value to map or returns error if conversion is not possible.
func ToMap(v any) (map[string]any, error) {
	switch t := any(v).(type) {
	case bson.M:
		return MToMap(t), nil
	case bson.D:
		return DToMap(t), nil
	default:
		return nil, fmt.Errorf("unexpected type %T", v)
	}
}

// MToMap converts bson.M to map.
func MToMap(m bson.M) map[string]any {
	return map[string]any(m)
}

// DToMap converts bson.D to map.
func DToMap(d bson.D) map[string]any {
	m := make(map[string]any, len(d))
	for _, e := range d {
		m[e.Key] = e.Value
	}
	return m
}
