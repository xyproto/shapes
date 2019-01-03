package shapes

import (
	"math"
)

const (
	twoPi    = 2.0 * math.Pi
	radToDeg = 360.0 / twoPi
)

func RadToDeg(r float64) float64 {
	return r * radToDeg
}

// --- Angle between two unit vectors, in radians ---

func Anglef(a, b *Pointf) float64 {
	//angle := 0.0
	//if (a.x == 0) && (a.y == 0) {
	//	angle = math.Atan2(b.y, b.x)
	//} else {
	//	angle = math.Acos(Dotf(a, b))
	//}
	angle := math.Acos(Dotf(a, b))
	for angle < 0 {
		angle += twoPi
	}
	return angle
}

func Angleif(a, b *Pointi) float64 {
	//angle := 0.0
	//if (a.x == 0) && (a.y == 0) {
	//	angle = math.Atan2(float64(b.y), float64(b.x))
	//} else {
	//	angle = math.Acos(float64(Doti(a, b)))
	//}
	angle := math.Acos(float64(Doti(a, b)))
	for angle < 0 {
		angle += twoPi
	}
	return angle
}

func Anglei(a, b *Pointi) int {
	anglei := int(Angleif(a, b))
	for anglei < 0 {
		anglei += 360
	}
	return anglei
}

// --- Angle between two unit vectors, in degrees ---

func DAnglef(a, b *Pointf) float64 {
	return Anglef(a, b) * radToDeg
}

func DAngleif(a, b *Pointi) float64 {
	return Angleif(a, b) * radToDeg
}

func DAnglei(a, b *Pointi) int {
	return int(DAngleif(a, b))
}
