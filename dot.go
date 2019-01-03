package shapes

// --- Dot product ---

func Dotf(a, b *Pointf) float64 {
	return a.x*b.x + a.y*b.y
}

func Dotif(a, b *Pointi) float64 {
	return float64(a.x*b.x + a.y*b.y)
}

func Doti(a, b *Pointi) int {
	return a.x*b.x + a.y*b.y
}

func (a *Pointf) Dot(b *Pointf) float64 {
	return a.x*b.x + a.y*b.y
}

func (a *Pointi) Dotf(b *Pointi) float64 {
	return float64(a.x*b.x + a.y*b.y)
}

func (a *Pointi) Dot(b *Pointi) int {
	return a.x*b.x + a.y*b.y
}
