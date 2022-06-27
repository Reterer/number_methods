package main

import (
	"fmt"
	"math"
)

/*
	[ ] - Лангранж
	[ ] - Ньютон
*/
type Point struct {
	x, y float64
}

func MakeLangrangeInterpolation(points []Point) func(float64) float64 {
	n := len(points)
	l := make([]float64, n)
	for i := 0; i < n; i++ {
		var w float64 = 1
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			w *= points[i].x - points[j].x
		}

		l[i] = points[i].y / w
	}

	return func(x float64) float64 {
		var res float64
		for i := 0; i < n; i++ {
			var xx float64 = 1
			for j := 0; j < n; j++ {
				if j == i {
					continue
				}
				xx *= x - points[j].x
			}
			res += l[i] * xx
		}
		return res
	}
}

func MakeNewtonInterpolation(points []Point) func(float64) float64 {
	n := len(points)
	mat := make([][]float64, n)
	mat[0] = make([]float64, n)
	for i := 0; i < n; i++ {
		mat[0][i] = points[i].y
	}

	for i := 1; i < n; i++ {
		mat[i] = make([]float64, n-i)
		for j := i; j < n; j++ {
			mat[i][j-i] = (mat[i-1][j-i] - mat[i-1][j-i+1]) / (points[j-i].x - points[j].x)
		}
	}

	return func(x float64) float64 {
		ans := float64(0)
		pow := float64(1)
		for i := 0; i < n; i++ {
			ans += pow * mat[i][0]
			pow *= x - points[i].x
		}
		return ans
	}
}

func f(x float64) float64 {
	return math.Log(x)
}

func main() {
	{
		points := []Point{
			{0.1, -2.30259},
			{0.5, -0.69315},
			{0.9, -0.10536},
			{1.3, 0.26236},
		}
		lf := MakeLangrangeInterpolation(points)
		eps := math.Abs(f(0.8) - lf(0.8))
		fmt.Println("Значение интерполяционного многочлена: ", lf(0.8), "Значение погрешности: ", eps)
	}
	{
		points := []Point{
			{0.1, -2.30259},
			{0.5, -0.69315},
			{0.9, -0.10536},
			{1.3, 0.26236},
		}
		nf := MakeNewtonInterpolation(points)
		eps := math.Abs(f(0.8) - nf(0.8))
		fmt.Println("Значение интерполяционного многочлена: ", nf(0.8), "Значение погрешности: ", eps)
	}
}
