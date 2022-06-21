package main

import (
	"fmt"

	"github.com/Reterer/number_methods/internal/run_through"
	"github.com/Reterer/number_methods/internal/utils"
	"github.com/Reterer/number_methods/pkg/matrix"
)

func readThreeDiagMatrix() *matrix.RMatrix {
	var n, m int
	if _, err := fmt.Scan(&n, &m); err != nil {
		panic("can't read matrix shape")
	}

	mat := matrix.MakeRealMatrix(n, m)

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
	b := utils.ReadRMatrix()

	x := run_through.Do(A, b)
	utils.PrintMatrix(x)
	utils.PrintEq(A, x, b)
}
