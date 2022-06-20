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

func doIteration(A, b *matrix.RMatrix, eps float64) *matrix.RMatrix {
	// TODO aii == 0 ?
	m, n := A.Shape()
	mm, nn := b.Shape()
	if m != n && m > 0 {
		// TODO PANIC
		return nil
	} else if nn != 1 {
		// TODO PANIC
		return nil
	} else if mm != m {
		// TODO PANIC
		return nil
	}

	beta := matrix.MakeRealMatrix(m, 1)
	alpha := matrix.MakeRealMatrix(m, n)

	for i := 0; i < m; i++ {
		aCol := A.GetCol(i)
		aii := aCol[i]
		alphaCol := alpha.GetCol(i)
		if aii == 0 {
			// TODO PANIC
			return nil
		}
		beta.SetEl(i, 0, b.GetEl(i, 0)/aii)

		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			alphaCol[j] = -aCol[j] / aii
		}
	}

	// TODO COPY matrix
	x := beta.Add(matrix.MakeRealMatrix(m, 1))
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

	m, n := A.Shape()
	for i := 0; i < m; i++ {
		colA := A.GetCol(i)
		for j := 0; j < n; j++ {
			norm += colA[j] * colA[j]
		}
	}
	norm = math.Sqrt(norm)

	return norm
}

func doZeidel(A, b *matrix.RMatrix, eps float64) *matrix.RMatrix {
	// TODO aii == 0 ?
	m, n := A.Shape()
	mm, nn := b.Shape()
	if m != n && m > 0 {
		// TODO PANIC
		return nil
	} else if nn != 1 {
		// TODO PANIC
		return nil
	} else if mm != m {
		// TODO PANIC
		return nil
	}

	beta := matrix.MakeRealMatrix(m, 1)
	alpha := matrix.MakeRealMatrix(m, n)

	for i := 0; i < m; i++ {
		aCol := A.GetCol(i)
		aii := aCol[i]
		alphaCol := alpha.GetCol(i)
		if aii == 0 {
			// TODO PANIC
			return nil
		}
		beta.SetEl(i, 0, b.GetEl(i, 0)/aii)

		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			alphaCol[j] = -aCol[j] / aii
		}
	}

	// TODO COPY matrix
	x := beta.Add(matrix.MakeRealMatrix(m, 1))
	norm := calcNorm(x)

	for iter := 0; norm > eps; iter++ {
		// Придется работать вручную
		// TODO add func in pkg
		norm = 0
		for i := 0; i < n; i++ {
			alphaCol := alpha.GetCol(i)
			var summ float64
			for j := 0; j < n; j++ {
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

func main() {
	var eps float64
	fmt.Scan(&eps)

	A := readRMatrix()
	b := readRMatrix()

	x := doIteration(A, b, eps)
	printMatrix(x)
	x = doZeidel(A, b, eps)
	printMatrix(x)
}
