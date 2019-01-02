package shapes

import (
	"math"
)

var (
	Zeroi = &Pointi{0, 0}
	Zerof = &Pointf{0.0, 0.0}
)

// --- Length and distance ---

func Distancei(a, b *Pointi) int {
	return int(math.Sqrt(float64((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y))))
}

func (a *Pointi) Length() int {
	return Distancei(Zeroi, a)
}

func Distancef(a, b *Pointf) float64 {
	return math.Sqrt((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y))
}

func (a *Pointf) Length() float64 {
	return Distancef(Zerof, a)
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

// Normalize, return an error if the length is zero
func (p *Pointi) Normalizei() (*Pointi, error) {
	return Divi(p, Newi(p.Length()))
}

// Normalize, return an error if the length is zero
func (p *Pointi) MustNormalizei() *Pointi {
	return MustDivi(p, Newi(p.Length()))
}

// Normalize, return an error if the length is zero
func (p *Pointf) Normalizef() (*Pointf, error) {
	return Divf(p, Newf(p.Length()))
}

// Normalize, return an error if the length is zero
func (p *Pointf) MustNormalizef() *Pointf {
	return MustDivf(p, Newf(p.Length()))
}
