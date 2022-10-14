package main

import (
	"fmt"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

type Point struct {
	x, y float64
}
type fn func(x, y, dydx float64) float64

func eulerMethod(f, g fn, a, b, h, y0, dydx0 float64) []Point {
	res := make([]Point, 1)
	res[0] = Point{
		x: a,
		y: y0,
	}

	b -= h
	for zk, yk, xk := dydx0, y0, a; xk < b; xk += h {
		zk += h * f(xk, yk, zk)
		yk += h * g(xk, yk, zk)

		res = append(res, Point{
			x: xk + h,
			y: yk,
		})
	}

	return res
}

func rungeKuttaMethod(f, g fn, a, b, h, y0, dydx0 float64) []Point {
	res := make([]Point, 1)
	res[0] = Point{
		x: a,
		y: y0,
	}

	b -= h
	for zk, yk, xk := dydx0, y0, a; xk <= b; xk += h {
		k1 := h * f(xk, yk, zk)
		l1 := h * g(xk, yk, zk)

		k2 := h * f(xk+h/2, yk+l1/2, zk+k1/2)
		l2 := h * g(xk+h/2, yk+l1/2, zk+k1/2)

		k3 := h * f(xk+h/2, yk+l2/2, zk+k2/2)
		l3 := h * g(xk+h/2, yk+l2/2, zk+k2/2)

		k4 := h * f(xk+h, yk+l3, zk+k3)
		l4 := h * g(xk+h, yk+l3, zk+k3)

		zk += (k1 + 2*k2 + 2*k3 + k4) / 6
		yk += (l1 + 2*l2 + 2*l3 + l4) / 6

		res = append(res, Point{
			x: xk + h,
			y: yk,
		})
	}

	return res
}

func adamsMethod(f, g fn, a, b, h, y0, dydx0 float64) []Point {
	res := make([]Point, 1)
	res[0] = Point{
		x: a,
		y: y0,
	}

	x := make([]float64, 1)
	y := make([]float64, 1)
	z := make([]float64, 1)
	x[0] = a
	y[0] = y0
	z[0] = dydx0

	// b -= h
	for zk, yk, xk, b := dydx0, y0, a, a+2*h; xk <= b; xk += h {
		k1 := h * f(xk, yk, zk)
		l1 := h * g(xk, yk, zk)

		k2 := h * f(xk+h/2, yk+l1/2, zk+k1/2)
		l2 := h * g(xk+h/2, yk+l1/2, zk+k1/2)

		k3 := h * f(xk+h/2, yk+l2/2, zk+k2/2)
		l3 := h * g(xk+h/2, yk+l2/2, zk+k2/2)

		k4 := h * f(xk+h, yk+l3, zk+k3)
		l4 := h * g(xk+h, yk+l3, zk+k3)

		zk += (k1 + 2*k2 + 2*k3 + k4) / 6
		yk += (l1 + 2*l2 + 2*l3 + l4) / 6

		x = append(x, xk+h)
		y = append(y, yk)
		z = append(z, zk)
		res = append(res, Point{
			x: xk + h,
			y: yk,
		})
	}

	b -= h
	for xk, yk, zk := x[3], y[3], z[3]; xk <= b; xk += h {
		n := len(x) - 1
		zk += (55*f(x[n], y[n], z[n]) -
			59*f(x[n-1], y[n-1], z[n-1]) +
			37*f(x[n-2], y[n-2], z[n-2]) -
			9*f(x[n-3], y[n-3], z[n-3])) * h / 24
		yk += (55*g(x[n], y[n], z[n]) -
			59*g(x[n-1], y[n-1], z[n-1]) +
			37*g(x[n-2], y[n-2], z[n-2]) -
			9*g(x[n-3], y[n-3], z[n-3])) * h / 24

		x = append(x, xk+h)
		y = append(y, yk)
		z = append(z, zk)
		res = append(res, Point{
			x: xk + h,
			y: yk,
		})
	}

	return res
}

func rungeRomberg(h1, h2 float64, p1, p2 []Point, p float64) []Point {
	k := h1 / h2
	div := math.Pow(k, p) - 1
	res := make([]Point, len(p1))
	for i := 0; i < len(p1); i++ {
		res[i].x = p1[i].x
		res[i].y = (p2[i].y - p1[i].y) / div
	}

	return res
}

func err(p []Point, anf fn) []Point {
	res := make([]Point, len(p))
	for i := 0; i < len(p); i++ {
		res[i].x = p[i].x
		res[i].y = math.Abs(p[i].y - anf(p[i].x, 0, 0))
	}
	return res
}

func genPlot(path string, anf fn, pointsB []Point, a float64, b float64, h float64) {
	p := plot.New()

	p.Title.Text = "Interpolation"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	steps := int((b - a) / h)
	anf_p := make(plotter.XYs, steps)
	c_p := make(plotter.XYs, len(pointsB))
	x := a
	for step := 0; step < steps; step++ {
		anf_p[step].X = x
		anf_p[step].Y = anf(x, 0, 0)

		x += h
	}
	for i := 0; i < len(pointsB); i++ {
		c_p[i].X = pointsB[i].x
		c_p[i].Y = pointsB[i].y
	}
	err := plotutil.AddLinePoints(p,
		"Original", anf_p,
		"Calc", c_p,
	)
	if err != nil {
		panic(err)
	}

	// Scatter
	// s_ad := make(plotter.XYs, len(pointsA))
	// s_bd := make(plotter.XYs, len(pointsB))
	// for i := 0; i < len(pointsA); i++ {
	// 	s_ad[i].X = pointsA[i].x
	// 	s_ad[i].Y = pointsA[i].y
	// 	s_bd[i].X = pointsB[i].x
	// 	s_bd[i].Y = pointsB[i].y
	// }
	// s_a, err := plotter.NewScatter(s_ad)
	// if err != nil {
	// 	panic(err)
	// }
	// s_b, err := plotter.NewScatter(s_ad)
	// if err != nil {
	// 	panic(err)
	// }

	// p.Add(s_a)
	// p.Legend.Add("Original", s_a)
	// p.Add(s_b)
	// p.Legend.Add("Original", s_b)
	// Save the plot to a PNG file.
	if err := p.Save(2000, 2000, path); err != nil {
		panic(err.Error())
	}
}

func main() {

	var (
		anf   fn        = func(x, y, dydx float64) float64 { return (1 + x) * math.Exp(x*x) }
		f     fn        = func(x, y, z float64) float64 { return 4*x*z - (4*x*x-2)*y }
		g     fn        = func(x, y, z float64) float64 { return z }
		a     float64   = 0
		b     float64   = 1
		h     []float64 = []float64{0.1, 0.05}
		y0    float64   = 1
		dydx0 float64   = 1
	)
	{
		p1 := eulerMethod(f, g, a, b, h[0], y0, dydx0)
		fmt.Println("Метод эйлера: h: ", h[0], "\n", p1)

		p2 := eulerMethod(f, g, a, b, h[1], y0, dydx0)
		fmt.Println("Метод эйлера: h: ", h[1], "\n", p2)
		fmt.Println("Точность:\n", "\tРунге-Ромберг:\t", rungeRomberg(h[0], h[1], p1, p2, 4), "\nОтностиельно точного решения:\n\th: ", h[0], " err:\t", err(p1, anf), "\n\th: ", h[1], " err:\t", err(p2, anf))
		genPlot("Euler.png", anf, p2, a, b, 0.05)
	}
	{
		p1 := rungeKuttaMethod(f, g, a, b, h[0], y0, dydx0)
		fmt.Println("Метод Рунге-Кутты четвертого порядка: h: ", h[0], "\n", p1)

		p2 := rungeKuttaMethod(f, g, a, b, h[1], y0, dydx0)
		fmt.Println("Метод Рунге-Кутты четвертого порядка: h: ", h[1], "\n", p2)

		fmt.Println("Точность:\n", "\tРунге-Ромберг:\t", rungeRomberg(h[0], h[1], p1, p2, 4), "\nОтностиельно точного решения:\n\th: ", h[0], " err:\t", err(p1, anf), "\n\th: ", h[1], " err:\t", err(p2, anf))
		genPlot("RungeKutt.png", anf, p2, a, b, 0.05)
	}
	{
		p1 := adamsMethod(f, g, a, b, h[0], y0, dydx0)
		fmt.Println("Метод Адамса четвертого порядка: h: ", h[0], "\n", p1)

		p2 := adamsMethod(f, g, a, b, h[1], y0, dydx0)
		fmt.Println("Метод Адамса четвертого порядка: h: ", h[1], "\n", p2)

		fmt.Println("Точность:\n", "\tРунге-Ромберг:\t", rungeRomberg(h[0], h[1], p1, p2, 4), "\nОтностиельно точного решения:\n\th: ", h[0], " err:\t", err(p1, anf), "\n\th: ", h[1], " err:\t", err(p2, anf))
		genPlot("Adams.png", anf, p2, a, b, 0.05)
	}
}
