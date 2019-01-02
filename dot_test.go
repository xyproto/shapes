package shapes

import (
	"fmt"
)

func AnglefExample() {

	a := NewPointf(0.0, 2.0)
	a.Normalize()

	b := NewPointf(2.0, 0.0)
	b.Normalize()

	// Find the angle, in degrees, between the two unit vectors
	fmt.Println(DAnglef(a, b))

	// Output:
	// 90
}
