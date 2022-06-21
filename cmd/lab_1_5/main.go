package main

import (
	"fmt"
	"math"
	"math/cmplx"

	"github.com/Reterer/number_methods/internal/utils"
	"github.com/Reterer/number_methods/pkg/matrix"
)

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
	n, m := A.Shape()
	// TODO добавить проверки

	Q = matrix.MakeRealMatrix(n, m)
	for i := 0; i < n; i++ {
		Q.SetEl(i, i, 1)
	}

	for i := 0; i < n-1; i++ {
		var norm float64
		for j := i; j < m; j++ {
			norm += math.Pow(A.GetEl(j, i), 2)
		}
		norm = math.Pow(norm, 0.5)

		v := matrix.MakeRealMatrix(n, 1)
		vT := matrix.MakeRealMatrix(1, n)
		v.SetEl(i, 0, A.GetEl(i, i)+sign(A.GetEl(i, i))*norm)
		vT.SetEl(0, i, A.GetEl(i, i)+sign(A.GetEl(i, i))*norm)
		for j := i + 1; j < m; j++ {
			v.SetEl(j, 0, A.GetEl(j, i))
			vT.SetEl(0, j, A.GetEl(j, i))
		}

		H := matrix.MakeRealMatrix(n, m)
		for i := 0; i < n; i++ {
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

func calcL(A *matrix.RMatrix, i int) (l_0, l_1 complex128) {
	aii := A.GetEl(i, i)
	aioi := A.GetEl(i+1, i)
	aiio := A.GetEl(i, i+1)
	aioio := A.GetEl(i+1, i+1)

	a := complex(1, 0)
	b := complex(aii+aioio, 0)
	c := complex(aii*aioio-aiio*aioi, 0)

	d := cmplx.Pow(b*b-4*a*c, 0.5)
	l_0 = (-b - d) / (2 * a)
	l_1 = (-b + d) / (2 * a)

	return l_0, l_1
}

func getL(A *matrix.RMatrix, pl []complex128, eps float64) ([]complex128, bool) {
	ok := true
	n, _ := A.Shape()
	l := make([]complex128, n)
	for i := 0; i < n; i++ {
		// Комплексно сопряжонный
		if i+1 < n && math.Abs(A.GetEl(i+1, i)) > eps {
			l_0, l_1 := calcL(A, i)

			l[i] = l_0
			l[i+1] = l_1

			ok = ok && (cmplx.Abs(l[i]-pl[i]) < eps && cmplx.Abs(l[i+1]-pl[i+1]) < eps)
			fmt.Println(ok, l[i]-pl[i], l[i+1]-pl[i+1])
			i++

		} else {
			l[i] = complex(A.GetEl(i, i), 0)
			ok = ok && (cmplx.Abs(l[i]-pl[i]) < eps)
			fmt.Println(ok, l[i]-pl[i])
		}
	}
	return l, ok
}

func main() {
	var eps float64
	fmt.Scan(&eps)

	A := utils.ReadRMatrix()
	n, _ := A.Shape()
	l := make([]complex128, n)
	isRun := true
	for i := 0; isRun; i++ {
		Q, R := doQR(A)
		A = R.MulByR(Q)

		var ok bool
		l, ok = getL(A, l, eps)
		isRun = !ok
		fmt.Println(l, isRun)

		fmt.Println("------- it:", i)
		utils.PrintMatrix(A)
	}

	fmt.Println(l)
}
