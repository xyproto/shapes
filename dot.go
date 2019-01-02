package shapes

import (
	"math"
)

// --- Dot product ---

func Dotf(a, b *Pointf) float64 {
	return a.x*b.x + a.y*b.y
}

func Doti(a, b *Pointi) int {
	return a.x*b.x + a.y*b.y
}

func (a *Pointf) Dot(b *Pointf) float64 {
	return a.x*b.x + a.y*b.y
}

func (a *Pointi) Dot(b *Pointi) int {
	return a.x*b.x + a.y*b.y
}

// --- Angle between two unit vectors, in radians ---

func Anglef(a, b *Pointf) float64 {
	return math.Acos(Dotf(a, b))
}

func Anglei(a, b *Pointi) int {
	return int(math.Acos(float64(Doti(a, b))))
}

// --- Angle between two unit vectors, in degrees ---

func DAnglef(a, b *Pointf) float64 {
	return (Anglef(a, b) / (2.0 * math.Pi)) * 360.0
}

func DAnglei(a, b *Pointi) int {
	return int((float64(Anglei(a, b)) / (2.0 * math.Pi)) * 360.0)
}
