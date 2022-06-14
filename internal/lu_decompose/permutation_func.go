package lu_decompose

import (
	"math"
)

type Permutator func(lu *LU, iter int)

func PermDefault(lu *LU, iter int, j int) {

	// Обновляем матрицы
	lu.P.SwapCol(iter, j)
	lu.U.SwapCol(iter, j)

	lu.L.SwapCol(iter, j)

	lu.L.SetEl(iter, j, 0)
	lu.L.SetEl(iter, iter, 1)
	lu.L.SetEl(j, iter, 0)
	lu.L.SetEl(j, j, 1)
}

func PermMin(lu *LU, i int) {
	if lu.U.GetEl(i, i) != 0 {
		return
	}
	for j := i; j < lu.n; j++ {
		if lu.U.GetEl(j, i) != 0 {
			PermDefault(lu, i, j)
			return
		}
	}
	// todo handle
	panic("эту матрицу нельзя разложить в LU")
}

func PermEveryIteration(lu *LU, i int) {
	maxJ := i
	maxV := lu.U.GetEl(i, i)
	for j := i; j < lu.n; j++ {
		el := lu.U.GetEl(j, i)
		if math.Abs(float64(1)-el) < math.Abs(float64(1)-maxV) {
			maxV = el
			maxJ = j
		}
	}
	if maxV != 0 {
		PermDefault(lu, i, maxJ)
		return
	}
	// todo handle
	panic("эту матрицу нельзя разложить в LU")
}

// user permutation
func UPermOnceMake(perm Permutator, col int) Permutator {
	return func(lu *LU, i int) {
		if i == 0 {
			PermDefault(lu, i, col)
		} else {
			perm(lu, i)
			lu.perm = perm
		}
	}
}

func UPermEverytime(dec *LU, i int) {
	// TODO
	panic("Не определена")
}
