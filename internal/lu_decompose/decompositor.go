package lu_decompose

import (
	"github.com/Reterer/number_methods/pkg/matrix"
)

type LU struct {
	perm       Permutator
	P          *matrix.PMatrix
	L          *matrix.RMatrix
	U          *matrix.RMatrix
	n          int
	decomposed bool
}

func MakeLU(perm Permutator, A *matrix.RMatrix) *LU {
	n, m := A.Shape()
	if n != m {
		// TODO ERR
		panic("n != m")
	}
	dec := LU{
		perm: perm,
		P:    matrix.MakePermutationMatrix(n),
		L:    matrix.MakeRealMatrix(n, n),
		U:    matrix.MakeRealMatrix(n, n),
		n:    n,
	}

	for i := 0; i < n; i++ {
		dec.L.SetEl(i, i, 1)
		copy(dec.U.GetCol(i), A.GetCol(i))
	}

	return &dec
}

func (lu *LU) Decompose() {
	if lu.decomposed {
		// TODO ERR
		return
	}

	niter := lu.n - 1
	for i := 0; i < niter; i++ {
		// Сначала делаем permutation
		if lu.perm != nil {
			lu.perm(lu, i)
		}

		// Затем обновляем L и U
		// Строка, ниже которой мы будем обнулять i столбец
		mainCol := lu.U.GetCol(i)
		if mainCol[i] == 0 {
			// todo handle
			panic("Main element is eq 0")
		}
		// Для каждой более нижней строки
		for j := i + 1; j < lu.n; j++ {
			currCol := lu.U.GetCol(j)
			del := currCol[i] / mainCol[i]
			// update U
			for k := i; k < lu.n; k++ {
				currCol[k] -= del * mainCol[k]
			}
			// update L
			lu.L.SetEl(j, i, del)
		}
	}

	lu.decomposed = true
}

func (lu *LU) Solve(b *matrix.RMatrix) *matrix.RMatrix {
	if !lu.decomposed {
		// TODO ERR
		return nil
	}
	n, m := b.Shape()
	if !(n == lu.n) {
		// TODO ERR
		return nil
	}

	nB := lu.P.MulByR(b)

	z := matrix.MakeRealMatrix(n, m)
	for k := 0; k < m; k++ {
		z.SetEl(0, k, nB.GetEl(0, k))
		for i := 1; i < n; i++ {
			var sum float64
			for j := 0; j < i; j++ {
				sum += lu.L.GetEl(i, j) * z.GetEl(j, k)
			}
			z.SetEl(i, k, nB.GetEl(i, k)-sum)
		}
	}
	x := matrix.MakeRealMatrix(n, m)
	for k := 0; k < m; k++ {
		x.SetEl(n-1, k, z.GetEl(n-1, k)/lu.U.GetEl(n-1, n-1))
		for i := n - 2; i >= 0; i-- {
			var sum float64
			for j := i + 1; j < n; j++ {
				sum += lu.U.GetEl(i, j) * x.GetEl(j, k)
			}
			x.SetEl(i, k, (z.GetEl(i, k)-sum)/lu.U.GetEl(i, i))
		}
	}

	return x
}
