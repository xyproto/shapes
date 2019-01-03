package shapes

type Trif struct {
	p1 *Pointf
	p2 *Pointf
	p3 *Pointf
}

type Trii struct {
	p1 *Pointi
	p2 *Pointi
	p3 *Pointi
}

func NewTrianglei(x1, y1, x2, y2, x3, y3 int) *Trii {
	return &Trii{&Pointi{x1, y1}, &Pointi{x2, y2}, &Pointi{x3, y3}}
}

func NewTrianglef(x1, y1, x2, y2, x3, y3 float64) *Trif {
	return &Trif{&Pointf{x1, y1}, &Pointf{x2, y2}, &Pointf{x3, y3}}
}

func Absi(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Absf(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func (t *Trif) Area() float64 {
	return Absf((t.p1.x*(t.p2.y-t.p3.y) + t.p2.x*(t.p3.y-t.p1.y) + t.p3.x*(t.p1.y-t.p2.y)) / 2.0)
}

// Return the area as a float64
func (t *Trii) Areaf() float64 {
	return Absf(float64(t.p1.x*(t.p2.y-t.p3.y)+t.p2.x*(t.p3.y-t.p1.y)+t.p3.x*(t.p1.y-t.p2.y)) / 2.0)
}

// Return the area as an int
func (t *Trii) Area() int {
	return Absi((t.p1.x*(t.p2.y-t.p3.y) + t.p2.x*(t.p3.y-t.p1.y) + t.p3.x*(t.p1.y-t.p2.y)) / 2)
}

func (t *Trii) Center() *Pointi {
	x := (t.p1.x + t.p2.x + t.p3.x) / 3
	y := (t.p1.y + t.p2.y + t.p3.y) / 3
	return &Pointi{x, y}
}

func (t *Trif) Center() *Pointf {
	x := (t.p1.x + t.p2.x + t.p3.x) / 3.0
	y := (t.p1.y + t.p2.y + t.p3.y) / 3.0
	return &Pointf{x, y}
}

// Given a point, returns the angle to the center of the triangle, in radians,
// as if the triangle was a unit circle.
// Returns an error if the distance between the two points is zero.
func (t *Trif) AngleToCenter(p *Pointf) (float64, error) {
	p.Sub(t.Center())
	if err := p.Normalize(); err != nil {
		return 0.0, err
	}
	rads := Anglef(Zerof, p)
	return rads, nil
}

func (t *Trif) MustAngleToCenter(p *Pointf) float64 {
	rad, err := t.AngleToCenter(p)
	if err != nil {
		panic(err)
	}
	return rad
}
