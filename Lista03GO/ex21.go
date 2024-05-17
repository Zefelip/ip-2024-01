package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y, Z float64
}

func main() {
	var N int
	fmt.Scanln(&N)

	var points []Point
	for i := 0; i < N; i++ {
		var x, y, z float64
		fmt.Scanln(&x, &y, &z)
		points = append(points, Point{x, y, z})
	}

	var maxAbsCoord float64

	for i := 0; i < N-1; i++ {
		vector := Vector(points[i], points[i+1])
		maxAbsCoord = math.Max(maxAbsCoord, math.Abs(vector.X))
		maxAbsCoord = math.Max(maxAbsCoord, math.Abs(vector.Y))
		maxAbsCoord = math.Max(maxAbsCoord, math.Abs(vector.Z))
		fmt.Printf("%.2f\n", maxAbsCoord)
	}

	fmt.Println() // Quebra de linha final
}

func Vector(p1, p2 Point) Point {
	return Point{p2.X - p1.X, p2.Y - p1.Y, p2.Z - p1.Z}
}
