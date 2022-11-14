package vector

import "fmt"

type Num interface {
	int | int64 | float64
	comparable
}

// can't use constant int as param (it's a value, not a type),
// to specify the dimensions of the vector, so I hard coded it as 3.
type Point[T Num] [3]T

func New3[T Num](v1, v2, v3 T) Point[T] {
	return Point[T]{v1, v2, v3}
}

func New[T Num](v1, v2 T) Point[T] {
	return Point[T]{v1, v2, 0}
}

// SelfAdd adds one point to another
func (v *Point[T]) SelfAdd(v2 Point[T]) error {
	if len(v) != len(v2) {
		return fmt.Errorf("vectors must be same length")
	}
	for i := range v {
		v[i] += v2[i]
	}
	return nil
}

// Add adds one point to another
func (v *Point[T]) Add(v2 Point[T]) Point[T] {
	v3 := *v
	for i := range v {
		v3[i] += v2[i]
	}
	return v3
}

func Add[T Num](v1, v2 *Point[T]) *Point[T] {
	v := v1
	err := v.SelfAdd(*v2)
	if err != nil {
		return nil
	}
	return v
}

type Line[T Num] [2]Point[T]

// NewLine returns a 2D line
func NewLine[T Num](x1, y1, x2, y2 T) Line[T] {
	return Line[T]{New(x1, y1), New(x2, y2)}
}

// NewLine3 returns a 3D line
func NewLine3[T Num](x1, y1, z1, x2, y2, z2 T) Line[T] {
	return Line[T]{New3(x1, y1, z1), New3(x2, y2, z2)}
}

// Points returns all the points on line
func (l *Line[T]) Points() []Point[T] {
	x, y := l[0][0], l[0][1]
	d := Point[T]{l[1][0] - x, l[1][1] - y, 0}
	var le [2]T

	out := make([]Point[T], 0, 10)

	for i := 0; i < 2; i++ {
		if d[i] < 0 {
			le[i] = -d[i]
			d[i] = -1
		} else if d[i] > 0 {
			le[i] = d[i]
			d[i] = 1
		}
	}
	i := 0
	if d[i] == 0 {
		i = 1
	}
	for x1, y1 := x, y; le[i] >= 0; x1, y1, le[i] = x1+d[0], y1+d[1], le[i]-1 {
		out = append(out, New(x1, y1))
	}
	return out
}

// IsRightAngled returns true if the line is right angled, assuming compass directions
// only
func (l *Line[T]) IsRightAngled() bool {
	return l[0][0] == l[1][0] || l[0][1] == l[1][1]
}

// Can't do this. :cry:
// var CompassDirections[N Num] = []Point[N]{
// 	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
// }

func CompassDirections[N Num]() []Point[N] {
	return []Point[N]{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}
}

var CompassDirectionsInt = CompassDirections[int]()
var CompassDirectionsFloat64 = CompassDirections[float64]()
