package graph

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/sitnikovik/ndbx/autograder/pkg/anyv"
)

// Value wraps a raw graph value with typed accessors.
type Value struct {
	raw any
}

// NewValue creates a graph value wrapper.
func NewValue(v any) Value {
	return Value{
		v,
	}
}

// Base returns the underlying generic value wrapper.
func (v Value) Base() anyv.Value {
	return anyv.NewValue(v.raw)
}

// MustNode returns the value as a node or panics.
func (v Value) MustNode() Node {
	x, ok := v.AsNode()
	if !ok {
		panic("value is not a node")
	}
	return x
}

// AsNode returns the value as a node when possible.
func (v Value) AsNode() (Node, bool) {
	x, ok := v.raw.(*neo4j.Node)
	if !ok {
		return Node{}, false
	}
	return NewNode(
		x.GetElementId(),
		PropertiesFromMap(x.GetProperties()),
	), true
}

// MustRelationship returns the value as a relationship or panics.
func (v Value) MustRelationship() Relationship {
	x, ok := v.AsRelationship()
	if !ok {
		panic("value is not a relationship")
	}
	return x
}

// AsRelationship returns the value as a relationship when possible.
func (v Value) AsRelationship() (Relationship, bool) {
	x, ok := v.raw.(*neo4j.Relationship)
	if !ok {
		return Relationship{}, false
	}
	return NewRelationship(
		x.GetElementId(),
		x.Type,
		NewPoint(x.StartElementId),
		NewPoint(x.EndElementId),
		PropertiesFromMap(x.GetProperties()),
	), true
}

// MustPath returns the value as a path or panics.
func (v Value) MustPath() Path {
	x, ok := v.AsPath()
	if !ok {
		panic("value is not a path")
	}
	return x
}

// AsPath returns the value as a path when possible.
func (v Value) AsPath() (Path, bool) {
	x, ok := v.raw.(*neo4j.Path)
	if !ok {
		return Path{}, false
	}
	nodes := make([]Node, 0, len(x.Nodes))
	rels := make([]Relationship, 0, len(x.Relationships))
	for _, n := range x.Nodes {
		nodes = append(nodes, NewNode(
			n.GetElementId(),
			PropertiesFromMap(n.GetProperties()),
		))
	}
	for _, r := range x.Relationships {
		rels = append(rels, NewRelationship(
			r.GetElementId(),
			r.Type,
			NewPoint(r.StartElementId),
			NewPoint(r.EndElementId),
			PropertiesFromMap(r.GetProperties()),
		))
	}
	return NewPath(
		nodes,
		rels,
	), true
}
