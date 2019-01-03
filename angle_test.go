package shapes

import (
	"fmt"
	"math"
)

func ExampleAnglef() {

	a := NewPointf(0.0, 1.0)
	a.Normalize()

	b := NewPointf(5.0, 0.0)
	b.Normalize()

	// Find the angle, in degrees, between the two unit vectors
	fmt.Println(DAnglef(a, b))

	// 90 degrees
	fmt.Println(RadToDeg(math.Pi / 2.0))

	// Output:
	// 315
	// 90
}
