package main

import (
	"fmt"
	"math"

	"github.com/Reterer/number_methods/internal/utils"
	"github.com/Reterer/number_methods/pkg/matrix"
)

func accYakobi(A *matrix.RMatrix) float64 {
	acc := float64(0)
	n, m := A.Shape()
	for i := 0; i < n; i++ {
		colA := A.GetCol(i)
		for j := i + 1; j < m; j++ {
			acc += math.Pow(colA[j], 2)
		}
	}
	acc = math.Sqrt(acc)
	return acc
}

func doYakobi(A *matrix.RMatrix, eps float64) (l []float64, x *matrix.RMatrix) {

	n, m := A.Shape()
	x = matrix.MakeRealMatrix(n, m)
	for i := 0; i < m; i++ {
		x.SetEl(i, i, 1)
	}

	currEps := accYakobi(A)
	for iter := 0; currEps > eps; iter++ {
		fmt.Println(currEps)
		utils.PrintMatrix(x)
		utils.PrintMatrix(A)
		fmt.Println("--------------\t", iter)

		var maxI, maxJ int = 0, 1
		var maxV float64
		for i := 0; i < n; i++ {
			aCol := A.GetCol(i)
			for j := i + 1; j < m; j++ {
				if math.Abs(aCol[j]) > maxV {
					maxV = math.Abs(aCol[j])
					maxI = i
					maxJ = j
				}
			}
		}

		Ui := matrix.MakeRealMatrix(n, m)
		UiT := matrix.MakeRealMatrix(m, n)
		for i := 0; i < n; i++ {
			if i == maxI || i == maxJ {
				continue
			}
			Ui.SetEl(i, i, 1)
			UiT.SetEl(i, i, 1)
		}
		// Расчет углов
		aij := A.GetEl(maxI, maxJ)
		aii := A.GetEl(maxI, maxI)
		ajj := A.GetEl(maxJ, maxJ)
		theta := math.Pi / 4
		if aii != ajj {
			theta = math.Atan((2*aij)/(aii-ajj)) / 2
		}
		Ui.SetEl(maxI, maxI, math.Cos(theta))
		Ui.SetEl(maxI, maxJ, -math.Sin(theta))
		Ui.SetEl(maxJ, maxI, math.Sin(theta))
		Ui.SetEl(maxJ, maxJ, math.Cos(theta))

		UiT.SetEl(maxI, maxI, math.Cos(theta))
		UiT.SetEl(maxI, maxJ, math.Sin(theta))
		UiT.SetEl(maxJ, maxI, -math.Sin(theta))
		UiT.SetEl(maxJ, maxJ, math.Cos(theta))

		// Применения
		x = x.MulByR(Ui)
		A = UiT.MulByR(A.MulByR(Ui))

		currEps = accYakobi(A)
	}

	l = make([]float64, n)
	for i := 0; i < n; i++ {
		l[i] = A.GetEl(i, i)
	}

	return l, x
}

func main() {
	var eps float64
	fmt.Scan(&eps)
	A := utils.ReadRMatrix()

	l, x := doYakobi(A, eps)
	fmt.Println(l)
	utils.PrintMatrix(x)
}
