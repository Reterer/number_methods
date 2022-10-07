package main

import (
	"fmt"
	"math"
)

type fn func(float64) float64

func newtonMethod(f, df fn, a, b, eps float64) (x float64, itcnt int) {
	var px float64
	x = a + (b-a)/2.

	for ; math.Abs(x-px) >= eps || itcnt == 0; itcnt++ {
		px = x
		x = x - f(x)/df(x)
		fmt.Println("iter: ", itcnt, "x: ", x)
	}

	return x, itcnt
}

func iterationMethod(f, df fn, a, b, eps float64) (x float64, itcnt int) {
	var px, q float64
	x = a + (b-a)/2.
	q = math.Max(math.Abs(df(a)), math.Abs(df(b)))
	fmt.Println("q: ", q)
	q = q / (1 - q)

	for ; math.Abs(x-px)*q >= eps || itcnt == 0; itcnt++ {
		px = x
		x = f(x)
		fmt.Println("iter: ", itcnt, "x: ", x)
	}

	return x, itcnt
}

func main() {
	var eps float64
	fmt.Scan(&eps)

	// f(x) = e ^ (2x) + 3x - 4
	f := func(x float64) float64 {
		return math.Exp(2*x) + 3*x - 4
	}
	df := func(x float64) float64 {
		return math.Exp(2*x) + 3
	}

	// f(x) = e ^ (2x) + 3x - 4 => x = ln(4-3x)/2
	phi := func(x float64) float64 {
		return math.Log(4-3*x) / 2
	}
	dphi := func(x float64) float64 {
		return -3 / (2 * (4 - 3*x))
	}

	{
		fmt.Println("Метод простых итераций")
		x, itcnt := iterationMethod(phi, dphi, 0.4, 0.6, eps)
		fmt.Println("Количество итераций: ", itcnt, "\tx:", x)
	}
	{
		fmt.Println("Метод Ньютона")
		x, itcnt := newtonMethod(f, df, 0.4, 0.6, eps)
		fmt.Println("Количество итераций: ", itcnt, "\tx:", x)
	}
}
