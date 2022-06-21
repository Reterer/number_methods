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
	if !(m == 1 && n == lu.n) {
		// TODO ERR
		return nil
	}

	// z := matrix.MakeRealMatrix(m, n)

	x := matrix.MakeRealMatrix(n, m)
	return x
}

func (lu *LU) CalcInverse() *matrix.RMatrix {
	if !lu.decomposed {
		// TODO ERR
		return nil
	}
	return nil
}
