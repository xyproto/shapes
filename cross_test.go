package shapes

import (
	"fmt"
)

func ExampleCrossf() {
	a := NewPointf(1, 2)
	b := NewPointf(3, 4)
	c := a.Cross(b)
	fmt.Println(c)

	// Output:
	// -2
}
