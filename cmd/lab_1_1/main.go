package main

import (
	"fmt"

	"github.com/Reterer/number_methods/internal/lu_decompose"
	"github.com/Reterer/number_methods/internal/utils"
	"github.com/Reterer/number_methods/internal/utils_lab_1_1/config"
	"github.com/Reterer/number_methods/pkg/matrix"
)

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

	A := utils.ReadRMatrix()
	LU := MakeLU(A)
	LU.Decompose()

	utils.PrintMatrix(LU.P)
	utils.PrintMatrix(LU.L)
	utils.PrintMatrix(LU.U)

	fmt.Println("Проверка --- PA ?= LU")
	utils.PrintMatrix(LU.P.MulByR(A))
	utils.PrintMatrix(LU.L.MulByR(LU.U))

	// Нахождение обратной матирцы AX = E
	fmt.Println("Нахождение обратной матрицы")
	n, _ := A.Shape()
	E := matrix.MakeRealMatrix(n, n)
	for i := 0; i < n; i++ {
		E.SetEl(i, i, 1)
	}
	invA := LU.Solve(E)
	utils.PrintMatrix(invA)

	fmt.Println("Решение СЛАУ")
	// Решение СЛАУ
	b := utils.ReadRMatrix()
	x := LU.Solve(b)
	utils.PrintEq(A, x, b)
	utils.PrintMatrix(x)
}
