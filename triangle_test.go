package shapes

import (
	"fmt"
)

func ExampleTri_Areaf() {
	t := NewTriangle(15, 15, 23, 30, 50, 25)
	fmt.Printf("%3.2f\n", t.Areaf())
	fmt.Printf("%d\n", t.Areai())
	// Output:
	// 222.50
	// 222
}

func ExampleTri_Center() {
	t := NewTriangle(15, 15, 47, 40, 65, 20)
	fmt.Println(t.Center())
	// Output:
	// (127/3, 25)
}

func ExampleTri_AngleFromCenterTo() {
	t := NewTriangle(0, 0, 1, 2, 2, 0)
	p := NewPoint(1, 0)
	angleInRadians := t.AngleFromCenterTo(p)
	fmt.Println(RadToDeg(angleInRadians))
	// Output:
	// 270
}
