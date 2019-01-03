package shapes

import (
	"errors"
	"fmt"
	"strings"
)

type Point interface {
	GetXY() (interface{}, interface{})
	SetXY(x, y interface{}) error
}

type PointCollection []Point

func (p *Pointi) GetXY() (interface{}, interface{}) {
	return p.XY()
}

func (p *Pointf) GetXY() (interface{}, interface{}) {
	return p.XY()
}

func (p *Pointi) SetXY(x, y interface{}) error {
	xi, ok := x.(int)
	if !ok {
		return errors.New("invalid type for x: not an int")
	}
	yi, ok := y.(int)
	if !ok {
		return errors.New("invalid type for y: not an int")
	}
	p.Set(xi, yi)
	return nil
}

func (p *Pointf) SetXY(x, y interface{}) error {
	xf, ok := x.(float64)
	if !ok {
		return errors.New("invalid type for x: not a float64")
	}
	yf, ok := y.(float64)
	if !ok {
		return errors.New("invalid type for y: not a float64")
	}
	p.Set(xf, yf)
	return nil
}

func NewPointCollection() *PointCollection {
	pc := PointCollection(make([]Point, 0))
	return &pc
}

func (pc *PointCollection) Push(p Point) {
	*pc = append(*pc, p)
}

func (pc *PointCollection) Pop() Point {
	lastIndex := len(*pc) - 1
	p := (*pc)[lastIndex]
	*pc = (*pc)[:lastIndex]
	return p
}

func (pc *PointCollection) String() string {
	var sb strings.Builder
	for i, p := range *pc {
		if i > 0 {
			sb.WriteString(", ")
		}
		x, y := p.GetXY()
		sb.WriteString(fmt.Sprintf("(%v, %v)", x, y))
	}
	return sb.String()
}
