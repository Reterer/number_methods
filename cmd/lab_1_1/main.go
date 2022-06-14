package main

import (
	"fmt"
	"strconv"

	"github.com/Reterer/number-methods/internal/lu_decompose"
	"github.com/Reterer/number-methods/internal/utils_lab_1_1/config"
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

func MakeLU(A *matrix.RMatrix) *lu_decompose.LU {
	cnfg := config.Get()
	permFunc := lu_decompose.PermMin

	switch cnfg.AutoPerm {
	case config.APM_nope:
		permFunc = nil
	case config.APM_min:
		permFunc = lu_decompose.PermMin
	case config.APM_all:
		permFunc = lu_decompose.PermEveryIteration
	}

	switch cnfg.UserPerm {
	case config.UPM_nope:
		// nothing
	case config.UPM_once:
		permFunc = lu_decompose.UPermOnceMake(permFunc, cnfg.UserPermVal)
	case config.UPM_everyime:
		permFunc = lu_decompose.UPermEverytime
	}

	return lu_decompose.MakeLU(permFunc, A)
}

func main() {
	config.Init()

	A := readRMatrix()
	LU := MakeLU(A)
	LU.Decompose()

	printMatrix(LU.P)
	printMatrix(LU.L)
	printMatrix(LU.U)

	fmt.Println("Проверка --- PA ?= LU")
	printMatrix(LU.P.MulByR(A))
	printMatrix(LU.L.MulByR(LU.U))

	b := readRMatrix()
	x := LU.Solve(b)
	printEq(A, x, b)
	printMatrix(x)
}
