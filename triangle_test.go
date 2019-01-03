package shapes

import (
	"fmt"
)

func ExampleAreaf() {
	t := NewTriangle(15, 15, 23, 30, 50, 25)
	fmt.Printf("%3.2f\n", t.Areaf())
	fmt.Printf("%d\n", t.Areai())
	// Output:
	// 222.50
	// 222
}

func ExampleCenterf() {
	t := NewTriangle(15, 15, 47, 40, 65, 20)
	fmt.Println(t.Center())
	// Output:
	// (42.333, 25.000)
}

func ExampleAngleToCenter() {
	t := NewTriangle(0, 0, 1, 2, 2, 0)
	p := NewPoint(1, 0)
	fmt.Println(RadToDeg(t.AngleToCenter(p)))
	// Output:
	// 90
}
