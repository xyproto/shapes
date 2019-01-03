package shapes

import (
	"math"
)

var (
	Zeroi = &Pointi{0, 0}
	Zerof = &Pointf{0.0, 0.0}
)

// --- Length and distance ---

// Returns the distance as an int
func Distancei(a, b *Pointi) int {
	return int(math.Sqrt(float64((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y))))
}

// Returns the length as an int
func (a *Pointi) Length() int {
	return Distancei(Zeroi, a)
}

// Returns the length as an int
func (a *Pointi) Distance(b *Pointi) int {
	return Distancei(a, b)
}

func Distancef(a, b *Pointf) float64 {
	return math.Sqrt((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y))
}

func (a *Pointf) Length() float64 {
	return Distancef(Zerof, a)
}

func (a *Pointf) Distance(b *Pointf) float64 {
	return Distancef(a, b)
}

// --- Normalization ---

// Normalize in-place, return an error if the length is zero
func (p *Pointi) Normalize() error {
	return p.Div(Newi(p.Length()))
}

// Normalize in-place, panic if the length is zero
func (p *Pointi) MustNormalize() {
	p.MustDiv(Newi(p.Length()))
}

// Normalize in-place, return an error if the length is zero
func (p *Pointf) Normalize() error {
	return p.Div(Newf(p.Length()))
}

// Normalize in-place, panic if the length is zero
func (p *Pointf) MustNormalize() {
	p.MustDiv(Newf(p.Length()))
}

// Normalizei, return an error if the length is zero
func (p *Pointi) Normalizei() (*Pointi, error) {
	return Divi(p, Newi(p.Length()))
}

// Normalize, panic if the length is zero
func (p *Pointi) MustNormalizei() *Pointi {
	return MustDivi(p, Newi(p.Length()))
}

// Normalizef, return an error if the length is zero
func (p *Pointf) Normalizef() (*Pointf, error) {
	return Divf(p, Newf(p.Length()))
}

// Normalize, panic if the length is zero
func (p *Pointf) MustNormalizef() *Pointf {
	return MustDivf(p, Newf(p.Length()))
}
