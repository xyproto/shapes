package shapes

import (
	"fmt"
)

func ExampleNewPoint() {
	p := NewPoint(1, 2)
	fmt.Println(p.String())
	// Output:
	// (1, 2)
}

func ExampleNewPointf() {
	p := NewPointf(1.2, 3.4)
	fmt.Println(p.Stringf())
	// Output:
	// (1.200, 3.400)
}
