package shapes

import (
	"fmt"
)

func ExamplePointCollection() {
	p1 := NewPointi(1, 2)
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
	// p2 (3.100, 4.300)
	// 2
	// (1, 2), (3.1, 4.3)
}
