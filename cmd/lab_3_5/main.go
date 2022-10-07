package main

import (
	"fmt"
	"math"
)

type fn func(x float64) float64

func rectangleMethod(xs, xe, h float64, f fn) float64 {
	var res float64
	for x := xs + h; x <= xe; x += h {
		res += h * f(x-h/2)
	}
	return res
}

func trapezeMethod(xs, xe, h float64, f fn) float64 {
	var res float64
	for x := xs + h; x <= xe; x += h {
		res += h * (f(x) + f(x-h))
	}
	return res / 2
}

func simpsonMethod(xs, xe, h float64, f fn) float64 {
	var res float64
	h *= 2
	for x := xs + h; x <= xe; x += h {
		res += h * (f(x) + 4*f(x-h/2) + f(x-h))
	}
	return res / 6
}

// TODO повотрить как работает этот метод.
func rungeRombergMethod(fh1, fh2, h1, h2 float64) float64 {
	k := h1 / h2
	return fh2 + (fh2-fh1)/(math.Pow(k, 2)-1)
}

func main() {
	var x0 float64 = -1
	var xk float64 = 1
	h1 := 0.5
	h2 := 0.25
	f := func(x float64) float64 {
		return x / math.Pow(3*x+4, 2)
	}

	{
		f1 := rectangleMethod(x0, xk, h1, f)
		f2 := rectangleMethod(x0, xk, h2, f)
		fmt.Println("Метод прямоугольников:", "h1: ", f1, "h2 ", f2)
		fmt.Println("Рунге-Ромберг-Ричардсон: ", rungeRombergMethod(f1, f2, h1, h2))

	}
	{
		f1 := trapezeMethod(x0, xk, h1, f)
		f2 := trapezeMethod(x0, xk, h2, f)
		fmt.Println("Метод трапеций:", "h1: ", f1, "h2 ", f2)
		fmt.Println("Рунге-Ромберг-Ричардсон: ", rungeRombergMethod(f1, f2, h1, h2))
	}
	{
		f1 := simpsonMethod(x0, xk, h1, f)
		f2 := simpsonMethod(x0, xk, h2, f)
		fmt.Println("Метод Симпсона:", "h1: ", f1, "h2 ", f2)
		fmt.Println("Рунге-Ромберг-Ричардсон: ", rungeRombergMethod(f1, f2, h1, h2))
	}
}
