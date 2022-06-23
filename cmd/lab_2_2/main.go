package main

import (
	"fmt"
	"math"

	"github.com/Reterer/number_methods/internal/lu_decompose"
	"github.com/Reterer/number_methods/internal/utils"
	"github.com/Reterer/number_methods/pkg/matrix"
)

type fn func(x *matrix.RMatrix) float64
type sfn struct {
	n   int
	fns []fn
	jak [][]fn
}

func (s *sfn) calc(x *matrix.RMatrix) *matrix.RMatrix {
	res := matrix.MakeRealMatrix(s.n, 1)
	for i := 0; i < s.n; i++ {
		res.SetEl(i, 0, s.fns[i](x))
	}
	return res
}

func (s *sfn) Jak(x *matrix.RMatrix) *matrix.RMatrix {
	jak := matrix.MakeRealMatrix(s.n, s.n)
	for i := 0; i < s.n; i++ {
		coljak := s.jak[i]
		colres := jak.GetCol(i)
		for j := 0; j < s.n; j++ {
			colres[j] = coljak[j](x)
		}
	}

	return jak
}

func norm(x *matrix.RMatrix) float64 {
	var ans float64
	n, m := x.Shape()
	for i := 0; i < n; i++ {
		var sum float64
		colX := x.GetCol(i)
		for j := 0; j < m; j++ {
			sum += math.Abs(colX[j])
		}
		if sum > ans {
			ans = sum
		}
	}

	return ans
}

func calcDet(x *matrix.RMatrix) float64 {
	n, _ := x.Shape()
	lu := lu_decompose.MakeLU(lu_decompose.PermMin, x)
	det := float64(1)
	for i := 0; i < n; i++ {
		det *= lu.U.GetEl(i, i)
	}
	return det
}

func newtonMethod(a, b *matrix.RMatrix, s *sfn, eps float64) (x *matrix.RMatrix, itcnt int) {
	n := s.n
	px := matrix.MakeRealMatrix(n, 1)
	x = b.Add(a).MulByConstant(1. / 2.) // x = (a + b) / 2

	for ; norm(x.Add(px.MulByConstant(-1))) > eps || itcnt == 0; itcnt++ {
		jak := s.Jak(x)
		detJak := calcDet(jak)

		dsf := matrix.MakeRealMatrix(n, 1)
		temp := matrix.MakeRealMatrix(n, 1)
		for j := 0; j < n; j++ {
			for i := 0; i < n; i++ {
				temp.SetEl(i, 0, jak.GetEl(i, j))
				jak.SetEl(i, j, s.fns[i](x))
			}
			dsf.SetEl(j, 0, -calcDet(jak)/detJak)
			for i := 0; i < n; i++ {
				jak.SetEl(i, j, temp.GetEl(i, 0))
			}
		}

		px = x
		x = x.Add(dsf)
		fmt.Println("iter:", itcnt, "\tx: ")
		utils.PrintMatrix(x)
	}

	return x, itcnt
}

func iterationMethod(a, b *matrix.RMatrix, s *sfn, eps float64) (x *matrix.RMatrix, itcnt int) {
	n := s.n
	px := matrix.MakeRealMatrix(n, 1)
	x = b.Add(a).MulByConstant(1. / 2.) // x = (a + b) / 2

	var q float64
	{
		jak := s.Jak(x)
		q = norm(jak)
		fmt.Println(q)
		q = q / (1 - q)
	}

	for ; norm(x.Add(px.MulByConstant(-1)))*q > eps || itcnt == 0; itcnt++ {
		px = x
		x = s.calc(x)
		fmt.Println("iter:", itcnt, "\tx: ")
		utils.PrintMatrix(x)
	}

	return x, itcnt
}

func fisrtSfn() *sfn {
	return &sfn{
		n: 2,
		fns: []fn{
			func(x *matrix.RMatrix) float64 { return math.Pow(x.GetEl(0, 0), 2) + math.Pow(x.GetEl(1, 0), 2) - 4 }, // f1
			func(x *matrix.RMatrix) float64 { return x.GetEl(0, 0) - math.Exp(x.GetEl(1, 0)) + 2 },                 // f2
		},
		jak: [][]fn{
			{
				func(x *matrix.RMatrix) float64 { return 2 * x.GetEl(0, 0) },
				func(x *matrix.RMatrix) float64 { return 2 * x.GetEl(1, 0) },
			},
			{
				func(x *matrix.RMatrix) float64 { return 1 },
				func(x *matrix.RMatrix) float64 { return -math.Exp(x.GetEl(1, 0)) },
			},
		},
	}
}

func secondSfn() *sfn {
	return &sfn{
		n: 2,
		fns: []fn{
			func(x *matrix.RMatrix) float64 { return math.Sqrt(4 - math.Pow(x.GetEl(1, 0), 2)) }, // f1
			func(x *matrix.RMatrix) float64 { return math.Log(x.GetEl(0, 0) + 2) },               // f2
		},
		jak: [][]fn{
			{
				func(x *matrix.RMatrix) float64 { return 0 },
				func(x *matrix.RMatrix) float64 { return -x.GetEl(1, 0) / math.Sqrt(4-math.Pow(x.GetEl(1, 0), 2)) },
			},
			{
				func(x *matrix.RMatrix) float64 { return 1. / (x.GetEl(0, 0) + 2) },
				func(x *matrix.RMatrix) float64 { return 0 },
			},
		},
	}
}

func main() {
	var eps float64
	fmt.Scan(&eps)

	{
		system := fisrtSfn()
		a := matrix.MakeRealMatrix(2, 1)
		b := matrix.MakeRealMatrix(2, 1)
		a.SetEl(0, 0, 1.0)
		b.SetEl(0, 0, 2.0)
		a.SetEl(1, 0, 0.5)
		b.SetEl(1, 0, 1.5)

		fmt.Println("Newton Method")
		newtonMethod(a, b, system, eps)
	}

	{
		system := secondSfn()
		a := matrix.MakeRealMatrix(2, 1)
		b := matrix.MakeRealMatrix(2, 1)
		a.SetEl(0, 0, 1.0)
		b.SetEl(0, 0, 2.0)
		a.SetEl(1, 0, 0.5)
		b.SetEl(1, 0, 1.5)

		fmt.Println("Iteration Method")
		iterationMethod(a, b, system, eps)
	}
}
