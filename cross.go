package shapes

// --- Cross product ---

func Crossf(a, b *Pointf) float64 {
	return a.x*b.y - a.y*b.x
}

func Crossif(a, b *Pointi) float64 {
	return float64(a.x*b.y - a.y*b.x)
}

func Crossi(a, b *Pointi) int {
	return a.x*b.y - a.y*b.x
}

func (a *Pointf) Cross(b *Pointf) float64 {
	return a.x*b.y - a.y*b.x
}

func (a *Pointi) Crossf(b *Pointi) float64 {
	return float64(a.x*b.y - a.y*b.x)
}

func (a *Pointi) Cross(b *Pointi) int {
	return a.x*b.y - a.y*b.x
}
