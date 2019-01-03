package shapes

import (
	"errors"
	"fmt"

	"github.com/xyproto/num"
)

// Addition, subtraction, multiplication and division for
// 2D points that are (int, int) and 2D points that are (float64, float64)
// The functions starting with "Must" will not return an error but
// panic instead. Functions that return errors are also provided.

type Point struct {
	x *num.Frac
	y *num.Frac
}

var (
	Zero = &Point{num.Zero, num.Zero}

	ErrDivZero = errors.New("divide by zero")
)

func NewPoint(x, y int) *Point {
	return NewPointFromInt(x, y)
}

func NewPointn(x, y *num.Frac) *Point {
	return &Point{x, y}
}

func NewPointFromInt(x, y int) *Point {
	return &Point{num.NewFromInt(x), num.NewFromInt(y)}
}

func NewPointFromFloat(x, y float64, maxIterations int) *Point {
	xn := num.NewFromFloat64(x, maxIterations)
	yn := num.NewFromFloat64(y, maxIterations)
	return &Point{xn, yn}
}

func NewPointf(x, y float64) *Point {
	return NewPointFromFloat(x, y, num.DefaultMaxIterations)
}

func (p *Point) XY() (*num.Frac, *num.Frac) {
	return p.x, p.y
}

func (p *Point) XYi() (int, int) {
	return p.x.Int(), p.y.Int()
}

func (p *Point) XYf() (float64, float64) {
	return p.x.Float64(), p.y.Float64()
}

// Setf is for setting the X and Y coordinates from two floats.
// maxIterations is how many iterations should be performed when
// converting the two floats to fractional numbers, per float.
func (p *Point) Setf(x, y float64, maxIterations int) {
	xn := num.NewFromFloat64(x, maxIterations)
	yn := num.NewFromFloat64(y, maxIterations)
	p.x = xn
	p.y = yn
}

func (p *Point) Set(x, y *num.Frac) {
	p.x, p.y = x, y
}

// --- Add ---

// Add and return
func Add(a, b *Point) *Point {
	return &Point{num.Add(a.x, b.x), num.Add(a.y, b.y)}
}

// Add in place
func (a *Point) Add(b *Point) {
	a.x.Add(b.x)
	a.y.Add(b.y)
}

// --- Sub ---

// Sub and return
func Sub(a, b *Point) *Point {
	return &Point{num.Sub(a.x, b.x), num.Sub(a.y, b.y)}
}

// Sub in place
func (a *Point) Sub(b *Point) {
	a.x.Sub(b.x)
	a.y.Sub(b.y)
}

// --- Mul ---

// Mul and return
func Mul(a, b *Point) *Point {
	axbx, err := num.Mul(a.x, b.x)
	if err != nil {
		panic(err)
	}
	ayby, err := num.Mul(a.y, b.y)
	if err != nil {
		panic(err)
	}
	return &Point{axbx, ayby}
}

// Mul in place
func (a *Point) Mul(b *Point) {
	a.x.Mul(b.x)
	a.y.Mul(b.y)
}

// --- Div ---

// Div and return
func Div(a, b *Point) (*Point, error) {
	x, err := num.Div(a.x, b.x)
	if err != nil {
		return nil, err
	}
	y, err := num.Div(a.y, b.y)
	if err != nil {
		return nil, err
	}
	return &Point{x, y}, nil
}

// Div and return
func MustDiv(a, b *Point) *Point {
	x, err := num.Div(a.x, b.x)
	if err != nil {
		panic(err)
	}
	y, err := num.Div(a.y, b.y)
	if err != nil {
		panic(err)
	}
	return &Point{x, y}
}

// Div in place
func (a *Point) Div(b *Point) error {
	if b.x.IsZero() || b.y.IsZero() {
		return ErrDivZero
	}
	a.x.Div(b.x)
	a.y.Div(b.y)
	return nil
}

// Div in place
func (a *Point) MustDiv(b *Point) {
	if b.x.IsZero() || b.y.IsZero() {
		panic(ErrDivZero)
	}
	a.x.Div(b.x)
	a.y.Div(b.y)
}

// --- Create points where x == y ---

func Newi(x int) *Point {
	return NewPointFromInt(x, x)
}

func Newf(x float64, maxIterations int) *Point {
	return NewPointFromFloat(x, x, maxIterations)
}

// --- Strings ---

func (p *Point) String() string {
	return fmt.Sprintf("(%.3f, %.3f)", p.x.Float64(), p.y.Float64())
}

// --- Zero checks ---

func (p *Point) IsZero() bool {
	return p.x.IsZero() && p.y.IsZero()
}
