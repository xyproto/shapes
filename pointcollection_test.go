package shapes

import (
	"fmt"
)

func ExamplePointCollection() {
	p1 := NewPoint(1, 2)
	p2 := NewPointf(3.1, 4.3)

	fmt.Println("p1", p1)
	fmt.Println("p2", p2)

	pc := NewPointCollection()
	pc.Push(p1)
	pc.Push(p2)

	fmt.Println(len(*pc))

	fmt.Println(pc.String())

	// Output:
	// p1 (1, 2)
	// p2 (31/10, 43/10)
	// 2
	// (1, 2), (31/10, 43/10)
}
