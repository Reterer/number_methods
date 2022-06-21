package main

import (
	"fmt"
	"math"

	"github.com/Reterer/number_methods/internal/utils"
	"github.com/Reterer/number_methods/pkg/matrix"
)

func doIteration(A, b *matrix.RMatrix, eps float64) *matrix.RMatrix {
	n, m := A.Shape()
	nn, mm := b.Shape()
	if n != m && n > 0 {
		// TODO PANIC
		return nil
	} else if mm != 1 {
		// TODO PANIC
		return nil
	} else if nn != n {
		// TODO PANIC
		return nil
	}

	beta := matrix.MakeRealMatrix(n, 1)
	alpha := matrix.MakeRealMatrix(n, m)

	for i := 0; i < n; i++ {
		aCol := A.GetCol(i)
		aii := aCol[i]
		alphaCol := alpha.GetCol(i)
		if aii == 0 {
			// TODO aii == 0 ?
			return nil
		}
		beta.SetEl(i, 0, b.GetEl(i, 0)/aii)

		for j := 0; j < m; j++ {
			if j == i {
				continue
			}
			alphaCol[j] = -aCol[j] / aii
		}
	}

	// TODO COPY matrix
	x := beta.Add(matrix.MakeRealMatrix(n, 1))
	norm := calcNorm(x)
	for iter := 0; norm > eps; iter++ {
		nx := beta.Add(alpha.MulByR(x))
		norm = calcNorm(nx.Add(x.MulByConstant(-1)))
		x = nx
		fmt.Println(norm)
	}

	return x
}

func calcNorm(A *matrix.RMatrix) float64 {
	var norm float64

	n, m := A.Shape()
	for i := 0; i < n; i++ {
		colA := A.GetCol(i)
		for j := 0; j < m; j++ {
			norm += colA[j] * colA[j]
		}
	}
	norm = math.Sqrt(norm)

	return norm
}

func doZeidel(A, b *matrix.RMatrix, eps float64) *matrix.RMatrix {
	// TODO aii == 0 ?
	n, m := A.Shape()
	nn, mm := b.Shape()
	if n != m && n > 0 {
		// TODO PANIC
		return nil
	} else if mm != 1 {
		// TODO PANIC
		return nil
	} else if nn != n {
		// TODO PANIC
		return nil
	}

	beta := matrix.MakeRealMatrix(n, 1)
	alpha := matrix.MakeRealMatrix(n, m)

	for i := 0; i < n; i++ {
		aCol := A.GetCol(i)
		aii := aCol[i]
		alphaCol := alpha.GetCol(i)
		if aii == 0 {
			// TODO PANIC
			return nil
		}
		beta.SetEl(i, 0, b.GetEl(i, 0)/aii)

		for j := 0; j < m; j++ {
			if j == i {
				continue
			}
			alphaCol[j] = -aCol[j] / aii
		}
	}

	// TODO COPY matrix
	x := beta.Add(matrix.MakeRealMatrix(n, 1))
	norm := calcNorm(x)

	for iter := 0; norm > eps; iter++ {
		// Придется работать вручную
		// TODO add func in pkg
		norm = 0
		for i := 0; i < n; i++ {
			alphaCol := alpha.GetCol(i)
			var summ float64
			for j := 0; j < m; j++ {
				summ += x.GetEl(j, 0) * alphaCol[j]
			}
			prev := x.GetEl(i, 0)
			x.SetEl(i, 0, summ+beta.GetEl(i, 0))
			norm += math.Pow(prev-x.GetEl(i, 0), 2)
		}
		norm = math.Sqrt(norm)
		fmt.Println(norm)

	}

	return x
}

func main() {
	var eps float64
	fmt.Scan(&eps)

	A := utils.ReadRMatrix()
	b := utils.ReadRMatrix()

	x := doIteration(A, b, eps)
	fmt.Println("Метод итераций: ")
	utils.PrintMatrix(x)
	x = doZeidel(A, b, eps)
	fmt.Println("Метод Зейделя")
	utils.PrintMatrix(x)
}
