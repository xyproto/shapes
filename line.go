package shapes

type Linei struct {
	a *Pointi
	b *Pointi
}

type Linef struct {
	a *Pointf
	b *Pointf
}

func (l *Linei) Length() int {
	return Distancei(l.a, l.b)
}

func (l *Linef) Length() float64 {
	return Distancef(l.a, l.b)
}

func (l *Linef) Points() []*Pointf {
	return []*Pointf{l.a, l.b}
}

func (l *Linei) Points() []*Pointi {
	return []*Pointi{l.a, l.b}
}
