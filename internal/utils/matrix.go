package utils

import (
	"fmt"
	"strconv"

	"github.com/Reterer/number_methods/pkg/matrix"
)

func ReadRMatrix() *matrix.RMatrix {
	var n, m int
	if _, err := fmt.Scan(&n, &m); err != nil {
		panic("can't read matrix shape")
	}

	mat := matrix.MakeRealMatrix(n, m)
	fillRMatrix(mat)

	return mat
}

func fillRMatrix(mat *matrix.RMatrix) {
	n, m := mat.Shape()
	for i := 0; i < n; i++ {
		col := mat.GetCol(i)
		for j := 0; j < m; j++ {
			if _, err := fmt.Scan(&col[j]); err != nil {
				panic("can't read element")
			}
		}
	}
}

func PrintMatrix(mat matrix.ShaperElGetter) {
	n, m := mat.Shape()
	fmt.Printf("%d %d\n", n, m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%8.4f ", mat.GetEl(i, j))
		}
		fmt.Println()
	}
}

func PrintEq(A, x, b matrix.ShaperElGetter) {
	n, m := A.Shape()

	for i := 0; i < n; i++ {
		sum := float64(0)
		for j := 0; j < m; j++ {
			if x == nil {
				fmt.Printf("%8.4f*x"+strconv.Itoa(j+1)+" ", A.GetEl(i, j))
			} else {
				fmt.Printf("%8.4f*%8.4f ", A.GetEl(i, j), x.GetEl(j, 0))
				sum += A.GetEl(i, j) * x.GetEl(j, 0)
			}
			if j+1 < m {
				fmt.Printf("+ ")
			}
		}

		fmt.Printf("= %8.4f\t (act %8.4f)\n", b.GetEl(i, 0), sum)
	}
}
