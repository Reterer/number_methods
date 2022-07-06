package main

import (
	"fmt"
	"math"

	"github.com/Reterer/number_methods/internal/run_through"
	"github.com/Reterer/number_methods/pkg/matrix"
)

type fn func(x, y, z float64) float64
type Point struct {
	x, y float64
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

func shootingMethod(f, g fn, a, b, y0, y1, h, etaprev, etacurr, eps float64) []Point {
	zprev := etaprev
	yprev := rungeKuttaMethod(f, g, a, b, h, y0, zprev)
	zcurr := etacurr
	ycurr := rungeKuttaMethod(f, g, a, b, h, y0, zcurr)

	F := func(y []Point) float64 { return y[len(y)-1].y - y1 }

	for math.Abs(F(ycurr)) > eps {
		temp := zcurr
		zcurr = zcurr - (zcurr-zprev)/(F(ycurr)-F(yprev))*F(ycurr)
		zprev = temp

		yprev = ycurr
		ycurr = rungeKuttaMethod(f, g, a, b, h, y0, zcurr)
	}

	return ycurr
}

func finiteDifferenceMethod(p, q fn, a, b, y0, y1, h, alpha, beta, gamma, delta float64) []Point {
	n := int((b-a)/h) + 1
	diag := matrix.MakeRealMatrix(n, n)
	bias := matrix.MakeRealMatrix(n, 1)
	points := make([]Point, n)
	xk := a
	for i := 0; i < n; i++ {
		points[i].x = xk
		// A
		if i > 0 {
			if i == n-1 {
				diag.SetEl(i, i-1, -gamma)
			} else {
				diag.SetEl(i, i-1, 1-p(xk, 0, 0)*h/2)
			}
		}
		// B
		if i == 0 {
			diag.SetEl(i, i, alpha*h-beta)
		} else if i == n-1 {
			diag.SetEl(i, i, delta*h+gamma)
		} else {
			diag.SetEl(i, i, q(xk, 0, 0)*h*h-2)
		}

		// C
		if i < n-1 {
			if i == 0 {
				diag.SetEl(i, i+1, beta)
			} else {
				diag.SetEl(i, i+1, 1+p(xk, 0, 0)*h/2)
			}
		}
		xk += h
	}
	// bias
	bias.SetEl(0, 0, y0*h)
	bias.SetEl(n-1, 0, y1*h)

	ans := run_through.Do(diag, bias)
	for i := 0; i < n; i++ {
		points[i].y = ans.GetEl(i, 0)
	}

	return points
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

func main() {
	var (
		anf fn = func(x, y, z float64) float64 { return -math.Tan(x) }
		f   fn = func(x, y, z float64) float64 { return 2 * (1 + (math.Pow(math.Tan(x), 2))) * y }
		g   fn = func(x, y, z float64) float64 { return z }
		p   fn = func(x, y, z float64) float64 { return 0 }
		q   fn = func(x, y, z float64) float64 { return -2 * (1 + (math.Pow(math.Tan(x), 2))) }

		a       float64   = 0
		b       float64   = math.Pi / 6
		h       []float64 = []float64{math.Pi / 30, math.Pi / 60}
		y0      float64   = 0
		y1      float64   = anf(b, 0, 0)
		alpha   float64   = 1
		delta   float64   = 1
		gamma   float64   = 0
		beta    float64   = 0
		etaprev float64   = 0.8
		etacurr float64   = 1
		eps     float64   = 0.001
	)

	// fmt.Println(shootingMethod(f, g, a, b, y0, y1, h[0], etaprev, etacurr, eps))
	// fmt.Println(finiteDifferenceMethod(p, q, a, b, y0, y1, h[0], alpha, beta, gamma, delta))

	{
		p1 := shootingMethod(f, g, a, b, y0, y1, h[0], etaprev, etacurr, eps)
		fmt.Println("Метод стрельбы: h: ", h[0], "\n", p1)

		p2 := shootingMethod(f, g, a, b, y0, y1, h[1], etaprev, etacurr, eps)
		fmt.Println("Метод стрельбы: h: ", h[1], "\n", p2)

		fmt.Println("Точность:\n", "\tРунге-Ромберг:\t", rungeRomberg(h[0], h[1], p1, p2, 4), "\nОтностиельно точного решения:\n\th: ", h[0], " err:\t", err(p1, anf), "\n\th: ", h[1], " err:\t", err(p2, anf))
	}
	{
		p1 := finiteDifferenceMethod(p, q, a, b, y0, y1, h[0], alpha, beta, gamma, delta)
		fmt.Println("Конечно-разностный метод: h: ", h[0], "\n", p1)

		p2 := finiteDifferenceMethod(p, q, a, b, y0, y1, h[1], alpha, beta, gamma, delta)
		fmt.Println("Конечно-разностный метод: h: ", h[1], "\n", p2)

		fmt.Println("Точность:\n", "\tРунге-Ромберг:\t", rungeRomberg(h[0], h[1], p1, p2, 4), "\nОтностиельно точного решения:\n\th: ", h[0], " err:\t", err(p1, anf), "\n\th: ", h[1], " err:\t", err(p2, anf))
	}
}
