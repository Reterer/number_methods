package main

import (
	"fmt"

	"github.com/Reterer/number_methods/internal/run_through"
	"github.com/Reterer/number_methods/internal/utils"
	"github.com/Reterer/number_methods/pkg/matrix"
)

type Point struct {
	x, y float64
}

func MakeSplainInterpolation(points []Point) func(float64) float64 {
	n := len(points) - 1

	c := make([]float64, n)
	{
		mat := matrix.MakeRealMatrix(n-1, n-1)
		b := matrix.MakeRealMatrix(n-1, 1)
		for i := 0; i < n-1; i++ {
			hc := points[i+2].x - points[i+1].x
			hp := points[i+1].x - points[i].x

			if i > 0 {
				mat.SetEl(i, i-1, hp)
			}

			mat.SetEl(i, i, 2*(hp+hc))

			if i < n-2 {
				mat.SetEl(i, i+1, hp)
			}

			fc := points[i+2].y - points[i+1].y
			fp := points[i+1].y - points[i].y
			b.SetEl(i, 0, 3*(fc/hc-fp/hp))

		}

		utils.PrintMatrix(mat)
		utils.PrintMatrix(b)
		c_2n := run_through.Do(mat, b)
		utils.PrintMatrix(c_2n)

		for i := 0; i < n-1; i++ {
			c[i+1] = c_2n.GetEl(i, 0)
		}
	}

	a := make([]float64, n)
	for i := 0; i < n; i++ {
		a[i] = points[i].y
	}

	b := make([]float64, n)
	for i := 0; i < n-1; i++ {
		fcurr := points[i+1].y - points[i].y
		hcurr := points[i+1].x - points[i].x
		b[i] = fcurr/hcurr - 1./3.*hcurr*(c[i+1]+2*c[i])
	}
	b[n-1] = (points[n].y-points[n-1].y)/(points[n].x-points[n-1].x) - 2./3.*(points[n].x-points[n-1].x)*c[n-1]

	d := make([]float64, n)
	for i := 0; i < n-1; i++ {
		hcurr := points[i+1].x - points[i].x
		d[i] = (c[i+1] - c[i]) / (3 * hcurr)
	}
	d[n-1] = -c[n-1] / (3 * (points[n].x - points[n-1].x))

	fmt.Println("A: ", a)
	fmt.Println("B: ", b)
	fmt.Println("C: ", c)
	fmt.Println("D: ", d)

	// TODO вывести все коэффиценты, сплайн как объект

	return func(x float64) float64 {
		// find interval
		i := 0
		for ; points[i+1].x < x; i++ {
		}
		dx := x - points[i].x
		return a[i] + b[i]*dx + c[i]*dx*dx + d[i]*dx*dx*dx
	}
}

func main() {
	{
		points := []Point{
			{0, 0},
			{1, 1.8415},
			{2, 2.9093},
			{3, 3.1411},
			{4, 3.2432},
		}
		lf := MakeSplainInterpolation(points)
		// eps := math.Abs(f(0.8) - lf(0.8))
		fmt.Println("Значение интерполяционного многочлена: ", lf(1.5))
	}
}
