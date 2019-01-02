package shapes

import "fmt"

// --- Strings ---

func (p *Pointi) String() string {
	return fmt.Sprintf("(%v, %v)", p.x, p.y)
}

func (p *Pointf) String() string {
	return fmt.Sprintf("(%v, %v)", p.x, p.y)
}
