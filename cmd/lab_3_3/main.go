package main

import (
	"fmt"
	"math"

	"github.com/Reterer/number_methods/internal/lu_decompose"
	"github.com/Reterer/number_methods/pkg/matrix"
)

type Point struct {
	x, y float64
}

func squareError(points []Point, f func(float64) float64) float64 {
	var err float64

	for i := 0; i < len(points); i++ {
		err += math.Pow(f(points[i].x)-points[i].y, 2)
	}

	return err
}

func lsm(points []Point, n int) func(float64) float64 {
	N := len(points)
	n1 := n + 1
	a := make([]float64, n1)

	{
		// Делаем систему
		A := matrix.MakeRealMatrix(n1, n1)
		b := matrix.MakeRealMatrix(n1, 1)

		for k := 0; k < n1; k++ {
			for i := 0; i < n1; i++ {
				var sumA, sumB float64
				for j := 0; j < N; j++ {
					sumA += math.Pow(points[j].x, float64(k+i))
					sumB += math.Pow(points[j].x, float64(k)) * points[j].y
				}
				A.SetEl(k, i, sumA)
				b.SetEl(k, 0, sumB)
			}
		}
		// Решаем систему
		LU := lu_decompose.MakeLU(lu_decompose.PermMin, A)
		LU.Decompose()
		aMat := LU.Solve(b)
		for i := 0; i < n1; i++ {
			a[i] = aMat.GetEl(i, 0)
		}
	}
	return func(x float64) float64 {
		var ans float64
		var xk float64 = 1

		for i := 0; i < n1; i++ {
			ans += xk * a[i]
			xk *= x
		}

		return ans
	}
}

func main() {
	{
		points := []Point{
			{0, 0},
			{1.7, 1.3038},
			{3.4, 1.8439},
			{5.1, 2.2583},
			{6.8, 2.6077},
			{8.5, 2.9155},
		}
		f := lsm(points, 1)
		fmt.Println("a + bx: ", f(5.1), " serr: ", squareError(points, f))
	}
	{
		points := []Point{
			{0, 0},
			{1.7, 1.3038},
			{3.4, 1.8439},
			{5.1, 2.2583},
			{6.8, 2.6077},
			{8.5, 2.9155},
		}
		f := lsm(points, 2)
		fmt.Println("a + bx + cx^2: ", f(5.1), " serr: ", squareError(points, f))
	}
}
