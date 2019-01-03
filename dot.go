package shapes

import (
	"github.com/xyproto/num"
)

// --- Dot product ---

func Dot(a, b *Point) *num.Frac {
	// a.x*b.x + a.y*b.y
	axbx, err := num.Mul(a.x, b.x)
	if err != nil {
		panic(err)
	}
	ayby, err := num.Mul(a.y, b.y)
	if err != nil {
		panic(err)
	}
	return num.Add(axbx, ayby)
}

func Dotf(a, b *Point) float64 {
	return Dot(a, b).Float64()
}

func Doti(a, b *Point) int {
	return Dot(a, b).Int()
}

func (a *Point) Dotf(b *Point) float64 {
	return Dot(a, b).Float64()
}

func (a *Point) Doti(b *Point) int {
	return Dot(a, b).Int()
}
