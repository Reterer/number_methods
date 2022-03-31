package matrix

import "fmt"

type PMatrix struct {
	perm []int
	n    int
}

func MakePermutationMatrix(n int) *PMatrix {
	P := PMatrix{
		perm: make([]int, n),
		n:    n,
	}

	for i := 0; i < n; i++ {
		P.perm[i] = i
	}

	return &P
}

func (Pp *PMatrix) GetEl(i, j int) float64 {
	if i >= Pp.n || j >= Pp.n {
		// TODO Panic
		panic("out of rnage")
	}

	row := Pp.perm[i]

	if row == j {
		return 1
	} else {
		return 0
	}
}

func (Pp *PMatrix) SwapCol(i, j int) {
	Pp.perm[i], Pp.perm[j] = Pp.perm[j], Pp.perm[i]
}

func (Pp *PMatrix) Shape() (m, n int) {
	m = Pp.n
	n = Pp.n
	return
}

func (Pp *PMatrix) GetValue(i int) int {
	return Pp.perm[i]
}

func (Ap *PMatrix) MulByPMatrix(Bp *PMatrix) *PMatrix {
	if Ap.n != Bp.n {
		panic(fmt.Sprintf("mul P on P:\n\tmatrix A shape: %v;\n\tmatrix B shape: %v;", Ap.n, Bp.n))
	}

	C := PMatrix{
		perm: make([]int, Ap.n),
		n:    Ap.n,
	}

	// c(i) = b(a(i))
	for i := 0; i < C.n; i++ {
		C.perm[i] = Bp.perm[Ap.perm[i]]
	}

	return &C
}

// others matrix
func (Ap *PMatrix) MulByRMatrix(Bp *RMatrix) *RMatrix {
	if Ap.n != Bp.m {
		panic(fmt.Sprintf("mul P on R:\n\tmatrix A shape: %vx%v;\n\tmatrix B shape: %vx%v;", Ap.n, Ap.n, Bp.m, Bp.n))
	}

	for i := 0; i < Ap.n; i++ {
		// swap cols
		j := Ap.perm[i]
		// либо перетсановку уже сделали, либо делать ее не надо
		if j <= i {
			continue
		}

		Bp.cols[i], Bp.cols[j] = Bp.cols[j], Bp.cols[i]
	}

	return Bp
}
