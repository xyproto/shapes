package shapes

import (
	"math"
)

const (
	twoPi    = 2.0 * math.Pi
	radToDeg = 180.0 / math.Pi
	degToRad = math.Pi / 180.0
)

type Angle struct {
	rad float64 // radians
}

// Deg returns the angle, in degrees
func (a *Angle) Deg() float64 {
	return a.rad * radToDeg
}

// Rad returns the angle, in radians
func (a *Angle) Rad() float64 {
	return a.rad
}

// SetDeg sets the angle, in degrees
func (a *Angle) SetDeg(d float64) {
	a.rad = d * degToRad
}

// SetRad sets the angle, in radians
func (a *Angle) SetRad(r float64) {
	a.rad = r
}

func RadToDeg(r float64) float64 {
	return r * radToDeg
}

func DegToRad(d float64) float64 {
	return d * degToRad
}

// --- Angle between two unit vectors, in radians ---

func NewAngle(a, b *Point) *Angle {
	return &Angle{Anglef(a, b)}
}

func AngleD(degrees float64) *Angle {
	return &Angle{degrees * degToRad}
}

func AngleR(radians float64) *Angle {
	return &Angle{radians}
}

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
