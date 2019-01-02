package shapes

import (
	"fmt"
)

func ExampleNewPointi() {
	p := NewPointi(1, 2)
	fmt.Println(p)
	// Output:
	// (1, 2)
}

func ExampleNewPointf() {
	p := NewPointf(1, 2)
	fmt.Println(p)
	// Output:
	// (1, 2)
}
