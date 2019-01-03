package shapes

import (
	"strings"
)

var FloatOutput bool

type PointCollection []Point

func NewPointCollection() *PointCollection {
	pc := PointCollection(make([]Point, 0))
	return &pc
}

func (pc *PointCollection) Push(p *Point) {
	*pc = append(*pc, *p)
}

func (pc *PointCollection) Pop() *Point {
	lastIndex := len(*pc) - 1
	p := (*pc)[lastIndex]
	*pc = (*pc)[:lastIndex]
	return &p
}

func (pc *PointCollection) String() string {
	var sb strings.Builder
	for i, p := range *pc {
		if i > 0 {
			sb.WriteString(", ")
		}
		if FloatOutput {
			sb.WriteString(p.Stringf())
		} else {
			sb.WriteString(p.String())
		}
	}
	return sb.String()
}
