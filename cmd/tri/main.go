package main

import (
	"fmt"
	"github.com/xyproto/shapes"
	"math"
	"time"
)

const (
	twoPi = 2.0 * math.Pi

	// ASCII viewport
	fromX = -5
	toX   = 10
	fromY = -4
	toY   = 5

	// Terminal codes for the forground characters
	y = "\033[1;93mo\033[0m"
	b = "\033[1;94m@\033[0m"
	r = "\033[1;31m#\033[0m"

	// The background character
	bg = " "

	// Terminal code for clearing the rest of the line
	clearToEnd = "\033[K"
)

func init() {
	fmt.Print("\033[H\033[2J")
}

func reset() {
	// Reset the cursor
	fmt.Print("\033[H")
}

func main() {
	var (
		// The original triangle
		ot = shapes.NewTrianglef(0, 0, 4, 0, 4, 2)

		// The center position of the original triangle
		center = ot.Center()

		// Three triangles
		t1, t2, t3 *shapes.Tri

		// Angle, in radians
		angle = 0.0
	)
	for {
		// Create three new triangles, by rotating ot with different angles
		t1 = ot.RotateAround(angle, center)
		t2 = ot.RotateAround(angle-(math.Pi/3.0), center)
		t3 = ot.RotateAround(angle-(math.Pi/3.0)*2.0, center)

		// Increase the angle just a bit, with wraparound
		angle += 0.07
		if angle >= twoPi {
			angle -= twoPi
		}

		// Reset the cursor
		reset()

		// Draw three triangles, using ASCII graphics.
		fmt.Println(t1.Draw(fromX, toX, fromY, toY, y, bg), clearToEnd)
		fmt.Println(t2.Draw(fromX, toX, fromY, toY, b, bg), clearToEnd)
		fmt.Println(t3.Draw(fromX, toX, fromY, toY, r, bg), clearToEnd)

		// Output the coordinates of the third triangle, as floats
		shapes.FloatOutput = true
		fmt.Println(t3.Points(), clearToEnd)

		// Output the coordinates of the third triangle, as fractions
		shapes.FloatOutput = false
		fmt.Println(t3.Points(), clearToEnd)

		// Wait a bit
		time.Sleep(10 * time.Millisecond)
	}
}
