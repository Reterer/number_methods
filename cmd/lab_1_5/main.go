package main

import (
	"fmt"
	"math"
	"math/cmplx"

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

func sign(num float64) float64 {
	if num == 0 {
		return 0
	} else if num < 0 {
		return -1
	} else {
		return 1
	}
}

func doQR(A *matrix.RMatrix) (Q, R *matrix.RMatrix) {
	m, n := A.Shape()
	// TODO добавить проверки

	Q = matrix.MakeRealMatrix(m, n)
	for i := 0; i < m; i++ {
		Q.SetEl(i, i, 1)
	}

	for i := 0; i < m-1; i++ {
		var norm float64
		for j := i; j < n; j++ {
			norm += math.Pow(A.GetEl(j, i), 2)
		}
		norm = math.Pow(norm, 0.5)

		v := matrix.MakeRealMatrix(m, 1)
		vT := matrix.MakeRealMatrix(1, m)
		v.SetEl(i, 0, A.GetEl(i, i)+sign(A.GetEl(i, i))*norm)
		vT.SetEl(0, i, A.GetEl(i, i)+sign(A.GetEl(i, i))*norm)
		for j := i + 1; j < n; j++ {
			v.SetEl(j, 0, A.GetEl(j, i))
			vT.SetEl(0, j, A.GetEl(j, i))
		}

		H := matrix.MakeRealMatrix(m, n)
		for i := 0; i < m; i++ {
			H.SetEl(i, i, 1)
		}
		c := -2 / vT.MulByR(v).GetEl(0, 0)
		H = H.Add(v.MulByR(vT).MulByConstant(c))
		Q = Q.MulByR(H)
		A = H.MulByR(A)

	}
	R = A

	return Q, R
}

func main() {
	A := readRMatrix()
	m, _ := A.Shape()
	l := make([]complex128, m)
	for i := 0; i < 10; i++ {
		Q, R := doQR(A)
		A = R.MulByR(Q)
		fmt.Println("------- it:", i)
		printMatrix(A)
	}

	eps := float64(0.01)
	for i := 0; i < m; i++ {
		// Комплексно сопряжонный
		if i+1 < m && math.Abs(A.GetEl(i+1, i)) > eps {
			aii := A.GetEl(i, i)
			aioi := A.GetEl(i+1, i)
			aiio := A.GetEl(i, i+1)
			aioio := A.GetEl(i+1, i+1)

			a := complex(1, 0)
			b := complex(aii+aioio, 0)
			c := complex(aii*aioio-aiio*aioi, 0)

			d := cmplx.Pow(b*b-4*a*c, 0.5)
			l_0 := (-b - d) / (2 * a)
			l_1 := (-b + d) / (2 * a)

			l[i] = l_0
			l[i+1] = l_1
			i++
		} else {
			l[i] = complex(A.GetEl(i, i), 0)
		}
	}
	fmt.Println(l)
}
