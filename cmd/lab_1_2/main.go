package main

import (
	"fmt"
	"strconv"

	"github.com/Reterer/number_methods/internal/run_through"
	"github.com/Reterer/number_methods/pkg/matrix"
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

	for i := 0; i < m; i++ {
		col := mat.GetCol(i)
		for k := 0; k < 3; k++ {
			j := -1 + k + i
			if j < 0 || j >= n {
				continue
			}

			if _, err := fmt.Scan(&col[j]); err != nil {
				panic("can't read element")
			}
		}
	}

	return mat
}

func printMatrix(mat matrix.ShaperElGetter) {
	m, n := mat.Shape()
	fmt.Printf("%d %d\n", m, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%3.0f\t", mat.GetEl(i, j))
		}
		fmt.Println()
	}
}

func printEq(A, x, b matrix.ShaperElGetter) {
	m, n := A.Shape()

	for i := 0; i < m; i++ {
		sum := float64(0)
		for j := 0; j < n; j++ {
			if x == nil {
				fmt.Printf("%3.3f*x"+strconv.Itoa(j+1)+" ", A.GetEl(i, j))
			} else {
				fmt.Printf("%3.3f*%3.3f ", A.GetEl(i, j), x.GetEl(j, 0))
				sum += A.GetEl(i, j) * x.GetEl(j, 0)
			}
			if j+1 < n {
				fmt.Printf("+ ")
			}
		}

		fmt.Printf("= %3.3f\t (act %3.3f)\n", b.GetEl(i, 0), sum)
	}
}

func main() {
	A := readThreeDiagMatrix()
	b := readRMatrix()

	x := run_through.Do(A, b)
	printMatrix(x)
	printEq(A, x, b)
}
