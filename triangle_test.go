package shapes

import (
	"fmt"
)

func ExampleAreaf() {
	t := NewTrianglei(15, 15, 23, 30, 50, 25)
	fmt.Printf("%3.2f\n", t.Areaf())
	fmt.Printf("%d\n", t.Area())
	// Output:
	// 222.50
	// 222
}

func ExampleCenterf() {
	t := NewTrianglef(15, 15, 47, 40, 65, 20)
	fmt.Println(t.Center())
	// Output:
	// (42.333, 25.000)
}

func ExampleAngleToCenter() {
	t := NewTrianglef(0, 0, 1, 2, 2, 0)
	p := NewPointf(1, 0)
	fmt.Println(RadToDeg(t.MustAngleToCenter(p)))
	// Output:
	// 90
}
