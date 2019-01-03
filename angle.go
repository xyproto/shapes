package shapes

import (
	"math"
)

const (
	twoPi    = 2.0 * math.Pi
	radToDeg = 180.0 / math.Pi
	degToRad = math.Pi / 180.0
)

func RadToDeg(r float64) float64 {
	return r * radToDeg
}

func DegToRad(d float64) float64 {
	return d * degToRad
}

// --- Angle between two unit vectors, in radians ---

func Anglef(a, b *Point) float64 {
	angle := 0.0
	if a.IsZero() {
		angle = math.Atan2(b.y.Float64(), b.x.Float64())
	} else {
		angle = math.Atan2(b.y.Float64()-a.y.Float64(), b.x.Float64()-a.x.Float64())
		//angle = math.Acos(Dotf(a, b))
	}
	for angle < 0 {
		angle += twoPi
	}
	return angle
}

func Anglei(a, b *Point) int {
	return int(Anglef(a, b))
}

// --- Angle between two unit vectors, in degrees ---

func DAnglef(a, b *Point) float64 {
	return Anglef(a, b) * radToDeg
}

func DAnglei(a, b *Point) int {
	return int(DAnglef(a, b))
}
