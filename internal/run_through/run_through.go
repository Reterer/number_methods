package run_through

import "github.com/Reterer/number_methods/pkg/matrix"

func Do(A *matrix.RMatrix, B *matrix.RMatrix) *matrix.RMatrix {
	m, n := A.Shape()
	mm, nn := B.Shape()
	if m != n && m > 0 {
		// TODO PANIC
		return nil
	} else if nn != 1 {
		// TODO PANIC
		return nil
	} else if mm != m {
		// TODO PANIC
		return nil
	}

	P := make([]float64, m+1)
	Q := make([]float64, m+1)

	getElFromCol := func(col []float64, j int) float64 {
		if j < 0 || j >= n {
			return 0
		}
		return col[j]
	}

	for i := 0; i < m; i++ {
		col := A.GetCol(i)
		a := getElFromCol(col, i-1)
		b := getElFromCol(col, i)
		c := getElFromCol(col, i+1)
		d := B.GetEl(i, 0)
		P[i+1] = -c / (b + a*P[i])
		Q[i+1] = (d - a*Q[i]) / (b + a*P[i])
	}

	x := matrix.MakeRealMatrix(m, 1)
	prevX := float64(0)
	for i := m - 1; i >= 0; i-- {
		x.SetEl(i, 0, P[i+1]*prevX+Q[i+1])
		prevX = x.GetEl(i, 0)
	}

	return x
}
