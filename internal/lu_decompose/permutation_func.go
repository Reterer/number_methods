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
	maxJ := i
	valueJ := lu.U.GetEl(i, i)
	for j := i; j < lu.n; j++ {
		if math.Abs(lu.U.GetEl(j, i)) > math.Abs(valueJ) {
			maxJ = j
			valueJ = lu.U.GetEl(j, i)
		}
	}
	if valueJ != 0 {
		PermDefault(lu, i, maxJ)
		return
	}
	panic("эту матрицу нельзя разложить в LU")
}

func PermEveryIteration(lu *LU, i int) {
	maxJ := i
	maxV := lu.U.GetEl(i, i)
	for j := i; j < lu.n; j++ {
		el := lu.U.GetEl(j, i)
		if math.Abs(el) < math.Abs(maxV) {
			maxV = el
			maxJ = j
		}
	}
	if maxV != 0 {
		PermDefault(lu, i, maxJ)
		return
	}
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
