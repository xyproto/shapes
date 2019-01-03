package shapes

import (
	"github.com/xyproto/num"
)

// --- Cross product ---

func Cross(a, b *Point) *num.Frac {
	// a.x*b.y - a.y*b.x
	axby, err := num.Mul(a.x, b.y)
	if err != nil {
		panic(err)
	}
	aybx, err := num.Mul(a.y, b.x)
	if err != nil {
		panic(err)
	}
	return num.Sub(axby, aybx)
}

func Crossf(a, b *Point) float64 {
	return Cross(a, b).Float64()
}

func Crossi(a, b *Point) int {
	return Cross(a, b).Int()
}

func (a *Point) Cross(b *Point) *num.Frac {
	return Cross(a, b)
}

func (a *Point) Crossf(b *Point) float64 {
	return Cross(a, b).Float64()
}

func (a *Point) Crossi(b *Point) int {
	return Cross(a, b).Int()
}
