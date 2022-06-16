package main

import (
	"fmt"
	"math"

	"github.com/Reterer/number_methods/pkg/matrix"
)

// [x] Метод простых итераций
// [x] Метод Зейделя
// [ ] Анализ количество итераций, необходимых для достижения заданной точности

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
			fmt.Printf("%3.4f\t", mat.GetEl(i, j))
		}
		fmt.Println()
	}
}

func doYakobi(A *matrix.RMatrix, k int) (l []float64, x *matrix.RMatrix) {
	// в течении k итераций
	//		найти наибольший внедиагональный элемент
	//		посчитать для него матрицу поворота Uk
	//		A = Uk.T * A * Uk
	//		U = U * Uk
	//	повторить

	// TODO оптимизировать
	// TODO проверки на симметричность?

	m, n := A.Shape()
	x = matrix.MakeRealMatrix(m, n)
	for i := 0; i < n; i++ {
		x.SetEl(i, i, 1)
	}

	for iter := 0; iter < k; iter++ {
		printMatrix(x)
		printMatrix(A)
		fmt.Println("--------------\t", iter)

		var maxI, maxJ int = 0, 1
		var maxV float64
		for i := 0; i < m; i++ {
			aCol := A.GetCol(i)
			for j := i + 1; j < n; j++ {
				if math.Abs(aCol[j]) > maxV {
					maxV = math.Abs(aCol[j])
					maxI = i
					maxJ = j
				}
			}
		}

		Ui := matrix.MakeRealMatrix(m, n)
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

		theta := math.Atan((2*aij)/(aii-ajj)) / 2
		fmt.Println(theta, math.Cos(theta), math.Sin(theta), aij, aii, ajj, maxI, maxJ)
		Ui.SetEl(maxI, maxI, math.Cos(theta))
		Ui.SetEl(maxI, maxJ, -math.Sin(theta))
		Ui.SetEl(maxJ, maxI, math.Sin(theta))
		Ui.SetEl(maxJ, maxJ, math.Cos(theta))
		printMatrix(Ui)

		UiT.SetEl(maxI, maxI, math.Cos(theta))
		UiT.SetEl(maxI, maxJ, math.Sin(theta))
		UiT.SetEl(maxJ, maxI, -math.Sin(theta))
		UiT.SetEl(maxJ, maxJ, math.Cos(theta))
		printMatrix(UiT)

		// Применения
		x = x.MulByR(Ui)
		A = UiT.MulByR(A.MulByR(Ui))

	}

	l = make([]float64, n)
	for i := 0; i < n; i++ {
		l[i] = A.GetEl(i, i)
	}

	return l, x
}

func main() {
	A := readRMatrix()

	k := 10
	// x := doIteration(A, b, k)
	l, x := doYakobi(A, k)
	fmt.Println(l)
	printMatrix(x)
}
