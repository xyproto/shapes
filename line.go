package shapes

type Line struct {
	a *Point
	b *Point
}

func (l *Line) Length() float64 {
	return Distance(l.a, l.b).Float64()
}

func (l *Line) Lengthi() int {
	return Distance(l.a, l.b).Int()
}

func (l *Line) Points() []*Point {
	return []*Point{l.a, l.b}
}
