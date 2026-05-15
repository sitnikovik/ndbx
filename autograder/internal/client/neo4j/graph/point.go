package graph

type Point struct {
	id string
}

func NewPoint(id string) Point {
	return Point{
		id,
	}
}
