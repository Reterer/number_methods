package run_through

import "github.com/Reterer/number_methods/pkg/matrix"

func Do(A *matrix.RMatrix, B *matrix.RMatrix) *matrix.RMatrix {
	n, m := A.Shape()
	nn, mm := B.Shape()
	if n != m && n > 0 {
		// TODO PANIC
		return nil
	} else if mm != 1 {
		// TODO PANIC
		return nil
	} else if nn != n {
		// TODO PANIC
		return nil
	}

	P := make([]float64, n+1)
	Q := make([]float64, n+1)

	getElFromCol := func(col []float64, j int) float64 {
		if j < 0 || j >= m {
			return 0
		}
		return col[j]
	}

	for i := 0; i < n; i++ {
		col := A.GetCol(i)
		a := getElFromCol(col, i-1)
		b := getElFromCol(col, i)
		c := getElFromCol(col, i+1)
		d := B.GetEl(i, 0)
		P[i+1] = -c / (b + a*P[i])
		Q[i+1] = (d - a*Q[i]) / (b + a*P[i])
	}

	x := matrix.MakeRealMatrix(n, 1)
	prevX := float64(0)
	for i := n - 1; i >= 0; i-- {
		x.SetEl(i, 0, P[i+1]*prevX+Q[i+1])
		prevX = x.GetEl(i, 0)
	}

	return x
}
