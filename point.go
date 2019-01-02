package shapes

import (
	"errors"
)

// Addition, subtraction, multiplication and division for
// 2D points that are (int, int) and 2D points that are (float64, float64)
// The functions starting with "Must" will not return an error but
// panic instead. Functions that return errors are also provided.

type Pointi struct {
	x int
	y int
}

type Pointf struct {
	x float64
	y float64
}

var (
	ErrDivZero = errors.New("divide by zero")
)

func NewPointi(x, y int) *Pointi {
	return &Pointi{x, y}
}

func NewPointf(x, y float64) *Pointf {
	return &Pointf{x, y}
}

func (p *Pointf) XY() (float64, float64) {
	return p.x, p.y
}

func (p *Pointi) XY() (int, int) {
	return p.x, p.y
}

func (p *Pointf) Set(x, y float64) {
	p.x, p.y = x, y
}

func (p *Pointi) Set(x, y int) {
	p.x, p.y = x, y
}

// --- Add ---

// Add and return
func Addf(a, b *Pointf) *Pointf {
	return &Pointf{a.x + b.x, a.y + b.y}
}

// Add in place
func (a *Pointf) Add(b *Pointf) {
	a.x += b.x
	a.y += b.y
}

// Add and return
func Addi(a, b *Pointi) *Pointi {
	return &Pointi{a.x + b.x, a.y + b.y}
}

// Add in place
func (a *Pointi) Add(b *Pointi) {
	a.x += b.x
	a.y += b.y
}

// --- Sub ---

// Sub and return
func Subf(a, b *Pointf) *Pointf {
	return &Pointf{a.x - b.x, a.y - b.y}
}

// Sub in place
func (a *Pointf) Sub(b *Pointf) {
	a.x -= b.x
	a.y -= b.y
}

// Sub and return
func Subi(a, b *Pointi) *Pointi {
	return &Pointi{a.x - b.x, a.y - b.y}
}

// Sub in place
func (a *Pointi) Sub(b *Pointi) {
	a.x -= b.x
	a.y -= b.y
}

// --- Mul ---

// Mul and return
func Mulf(a, b *Pointf) *Pointf {
	return &Pointf{a.x * b.x, a.y * b.y}
}

// Mul in place
func (a *Pointf) Mul(b *Pointf) {
	a.x *= b.x
	a.y *= b.y
}

// Mul and return
func Muli(a, b *Pointi) *Pointi {
	return &Pointi{a.x * b.x, a.y * b.y}
}

// Mul in place
func (a *Pointi) Mul(b *Pointi) {
	a.x *= b.x
	a.y *= b.y
}

// --- Div ---

// Div and return
func Divf(a, b *Pointf) (*Pointf, error) {
	if b.x == 0 || b.y == 0 {
		return nil, ErrDivZero
	}
	return &Pointf{a.x / b.x, a.y / b.y}, nil
}

// Div and return
func MustDivf(a, b *Pointf) *Pointf {
	if b.x == 0 || b.y == 0 {
		panic(ErrDivZero.Error())
	}
	return &Pointf{a.x / b.x, a.y / b.y}
}

// Div in place
func (a *Pointf) Div(b *Pointf) error {
	if b.x == 0 || b.y == 0 {
		return ErrDivZero
	}
	a.x /= b.x
	a.y /= b.y
	return nil
}

// Div in place
func (a *Pointf) MustDiv(b *Pointf) {
	if b.x == 0 || b.y == 0 {
		panic(ErrDivZero.Error())
	}
	a.x /= b.x
	a.y /= b.y
}

// Div and return
func Divi(a, b *Pointi) (*Pointi, error) {
	if b.x == 0 || b.y == 0 {
		return nil, ErrDivZero
	}
	return &Pointi{a.x / b.x, a.y / b.y}, nil
}

// Div and return
func MustDivi(a, b *Pointi) *Pointi {
	if b.x == 0 || b.y == 0 {
		panic(ErrDivZero.Error())
	}
	return &Pointi{a.x / b.x, a.y / b.y}
}

// Div in place
func (a *Pointi) Div(b *Pointi) error {
	if b.x == 0 || b.y == 0 {
		return ErrDivZero
	}
	a.x /= b.x
	a.y /= b.y
	return nil
}

// Div in place
func (a *Pointi) MustDiv(b *Pointi) {
	if b.x == 0 || b.y == 0 {
		panic(ErrDivZero.Error())
	}
	a.x /= b.x
	a.y /= b.y
}

// --- Create points where x == y ---

func Newi(x int) *Pointi {
	return &Pointi{x, x}
}

func Newf(x float64) *Pointf {
	return &Pointf{x, x}
}
