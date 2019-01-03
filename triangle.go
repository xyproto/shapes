package shapes

import (
	"github.com/xyproto/num"
)

type Tri struct {
	p1 *Point
	p2 *Point
	p3 *Point
}

func NewTriangle(x1, y1, x2, y2, x3, y3 int) *Tri {
	return &Tri{NewPoint(x1, y1), NewPoint(x2, y2), NewPoint(x3, y3)}
}

func NewTrianglef(x1, y1, x2, y2, x3, y3 float64) *Tri {
	return &Tri{NewPointf(x1, y1), NewPointf(x2, y2), NewPointf(x3, y3)}
}

func Absi(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Absf(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func (t *Tri) Center() *Point {
	// Dividing by 3 will never divide by zero
	x, _ := num.DivInt(num.Add(num.Add(t.p1.x, t.p2.x), t.p3.x), 3)
	y, _ := num.DivInt(num.Add(num.Add(t.p1.y, t.p2.y), t.p3.y), 3)
	return &Point{x, y}
}

func (t *Tri) Area() *num.Frac {
	// Absf((t.p1.x*(t.p2.y-t.p3.y) + t.p2.x*(t.p3.y-t.p1.y) + t.p3.x*(t.p1.y-t.p2.y)) / 2.0)

	a1 := num.Sub(t.p2.y, t.p3.y)
	a1.Mul(t.p1.x)

	a2 := num.Sub(t.p3.y, t.p1.y)
	a2.Mul(t.p2.x)

	a3 := num.Sub(t.p1.y, t.p2.y)
	a3.Mul(t.p3.x)

	a := num.Add(num.Add(a1, a2), a3)
	a.DivInt(2)

	a.Abs()
	return a
}

// Return the area as a float64
func (t *Tri) Areaf() float64 {
	return t.Area().Float64()
}

// Return the area as an int
func (t *Tri) Areai() int {
	return t.Area().Int()
}

// AngleFromCenterTo returns the angle from the center of the triangle
// to the given point, as if the center of the triangle is the center
// of the unit circle.
func (t *Tri) AngleFromCenterTo(p *Point) float64 {
	//center := t.Center()
	//// Translate and normalize point, so that the center is 0,0
	//x := num.Sub(p.x, center.x)
	//y := num.Sub(p.y, center.y)
	//tp := &Point{x, y}
	//tp.Normalize()
	// Return the angle from (0,0) to the translated normalized point
	//return Anglef(Zero, tp)

	// Return the angle from the center to the given point
	return Anglef(t.Center(), p)
}
