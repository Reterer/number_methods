package main

import (
	"fmt"

	"github.com/Reterer/number-methods/pkg/matrix"
)

func readRMatrix() *matrix.RMatrix {
	var m, n int
	if _, err := fmt.Scan(&m, &n); err != nil {
		panic("can't read matrix shape")
	}

	mat := matrix.MakeRealMatrix(m, n)
	fillRMatrix(mat)

	return mat
}

func fillRMatrix(mat *matrix.RMatrix) {
	m, n := mat.Shape()
	for i := 0; i < m; i++ {
		col := mat.GetCol(i)
		for j := 0; j < n; j++ {
			if _, err := fmt.Scan(&col[j]); err != nil {
				panic("can't read element")
			}
		}
	}
}

func readThreeDiagMatrix() *matrix.RMatrix {
	var m, n int
	if _, err := fmt.Scan(&m, &n); err != nil {
		panic("can't read matrix shape")
	}

	mat := matrix.MakeRealMatrix(m, n)

	for i := 0; i < n; i++ {
		col := mat.GetCol(i)
		for k := 0; k < 3; k++ {
			j := -1 + k + i
			if j < 0 || j >= m {
				continue
			}

			if _, err := fmt.Scan(&col[j]); err != nil {
				panic("can't read element")
			}
		}
	}

	return mat
}

func main() {
	A := readThreeDiagMatrix()
	b := readRMatrix()

	// x := run_through.do(A, b)

}
