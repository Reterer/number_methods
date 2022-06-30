package main

import "fmt"

type Point struct {
	x, y float64
}

func firstDerivative(points []Point) func(float64) float64 {
	n := len(points) - 2
	a := make([]float64, n)
	b := make([]float64, n)

	for i := 0; i < n; i++ {
		dy1 := points[i+1].y - points[i].y
		dy2 := points[i+2].y - points[i+1].y
		dx1 := points[i+1].x - points[i].x
		dx2 := points[i+2].x - points[i+1].x
		dxm := points[i+2].x - points[i].x

		a[i] = dy1 / dx1
		b[i] = (dy2/dx2 - dy1/dx1) / dxm
	}

	return func(x float64) float64 {
		i := 0
		for ; points[i+1].x < x; i++ {
		}
		return a[i] + b[i]*(2*x-points[i].x-points[i+1].x)
	}
}

func secondDerivative(points []Point) func(float64) float64 {
	n := len(points) - 2
	a := make([]float64, n)

	for i := 0; i < n; i++ {
		dy1 := points[i+1].y - points[i].y
		dy2 := points[i+2].y - points[i+1].y
		dx1 := points[i+1].x - points[i].x
		dx2 := points[i+2].x - points[i+1].x
		dxm := points[i+2].x - points[i].x

		a[i] = 2 * (dy2/dx2 - dy1/dx1) / dxm
	}

	return func(x float64) float64 {
		i := 0
		for ; points[i+1].x < x; i++ {
		}
		return a[i]
	}
}

func main() {
	{
		points := []Point{
			{0, 1},
			{0.1, 1.1052},
			{0.2, 1.2214},
			{0.3, 1.3499},
			{0.4, 1.3499},
		}
		f := firstDerivative(points)
		fmt.Println("f'", f(0.2))
	}
	{
		points := []Point{
			{0, 1},
			{0.1, 1.1052},
			{0.2, 1.2214},
			{0.3, 1.3499},
			{0.4, 1.3499},
		}
		f := secondDerivative(points)
		fmt.Println("f\"", f(0.2))
	}
}
