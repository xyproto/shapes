package shapes

import (
	"github.com/xyproto/num"
)

// --- Length and distance ---

// Distancef returns the distance between two points
func Distance(a, b *Point) *num.Frac {
	// sqrt((a.x-b.x)^2 + (a.y-b.y)^2)
	x2 := num.Square(num.Sub(a.x, b.x))
	y2 := num.Square(num.Sub(a.y, b.y))
	return num.Sqrt(num.Add(x2, y2))
}

// Returns the distance as an int
func Distancei(a, b *Point) int {
	return Distance(a, b).Int()
}

// Returns the distance as a float
func Distancef(a, b *Point) float64 {
	return Distance(a, b).Float64()
}

func (a *Point) Length() *num.Frac {
	// sqrt((a.x-0)^2 + (a.y-0)^2)
	return num.Sqrt(num.Add(num.Square(a.x), num.Square(a.y)))
}

// Returns the length as a float
func (a *Point) Lengthf() float64 {
	return a.Length().Float64()
}

// Returns the length as an int
func (a *Point) Lengthi() int {
	return a.Length().Int()
}

// --- Normalization ---

// Normalize
func (p *Point) Normalize() {
	l := p.Length()
	if l.IsZero() {
		// The normalized form of a null vector is a null vector
		return
	}
	p.x.Div(l)
	p.y.Div(l)
}
