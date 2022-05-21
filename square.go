package square

import "math"

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (side Square) End() Point {
return Point{x: side.start.x + int(side.a), y: side.start.y + int(side.a)}
}

func (square Square) Area() uint {
	return uint(math.Pow(float64(square.a), 2))
}

func (perimeter Square) Perimeter() uint {
	return 4 * perimeter.a
}
