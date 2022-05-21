package square

import "math"

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (side Square) End() uint {
return uint(math.Sqrt(2) * (float64(side.a)))
}

func (side Square) Area() uint {
	return uint(math.Pow(float64(square.a), 2))
}

func (perimeter Square) Perimeter() uint {
	return 4 * perimeter.a
}
