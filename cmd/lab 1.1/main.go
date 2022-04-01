package main

import (
	"fmt"
	"strconv"

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

type ShaperElGetter interface {
	matrix.ElGetter
	matrix.Shaper
}

func printMatrix(mat ShaperElGetter) {
	m, n := mat.Shape()
	fmt.Printf("%d %d\n", m, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%3.3f ", mat.GetEl(i, j))
		}
		fmt.Println()
	}
}

func printEq(A, x, b ShaperElGetter) {
	m, n := A.Shape()

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if x == nil {
				fmt.Printf("%3.3f*x"+strconv.Itoa(j+1)+" ", A.GetEl(i, j))
			} else {
				fmt.Printf("%3.3f*%3.3f ", A.GetEl(i, j), x.GetEl(j, 0))
			}
			if j+1 < n {
				fmt.Printf("+ ")
			}
		}

		fmt.Printf("= %3.3f\n", b.GetEl(i, 0))
	}
}

func main() {
	A := readRMatrix()
	b := readRMatrix()

	printEq(A, nil, b)
}
