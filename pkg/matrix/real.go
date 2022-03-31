package matrix

import "fmt"

type RMatrix struct {
	cols [][]float64
	m, n int
}

func MakeRealMatrix(m, n int) *RMatrix {
	A := RMatrix{
		cols: make([][]float64, m),
		m:    m,
		n:    n,
	}

	for i := 0; i < m; i++ {
		A.cols[i] = make([]float64, n)
	}

	return &A
}

func (Ap *RMatrix) Shape() (m, n int) {
	return Ap.m, Ap.n
}

func (Ap *RMatrix) SwapCol(i, j int) {
	Ap.cols[i], Ap.cols[j] = Ap.cols[j], Ap.cols[i]
}

func (Ap *RMatrix) MulByR(Bp *RMatrix) *RMatrix {
	if Ap.n != Bp.m {
		panic(fmt.Sprintf("mul:\n\tmatrix A shape: %vx%v;\n\tmatrix B shape: %vx%v;", Ap.m, Ap.n, Bp.m, Bp.n))
	}

	C := RMatrix{
		cols: make([][]float64, Ap.m),
		m:    Ap.m,
		n:    Bp.n,
	}
	for i := 0; i < C.m; i++ {
		C.cols[i] = make([]float64, C.n)
		for j := 0; j < C.n; j++ {
			sum := float64(0)
			for k := 0; k < Ap.n; k++ {
				sum += Ap.cols[i][k] * Bp.cols[k][j]
			}
			C.cols[i][j] = sum
		}
	}

	return &C
}

func (Ap *RMatrix) MulByConstant(c float64) *RMatrix {
	for i := 0; i < Ap.m; i++ {
		for j := 0; j < Ap.n; j++ {
			Ap.cols[i][j] *= c
		}
	}
	return Ap
}

func (Ap *RMatrix) Add(Bp *RMatrix) *RMatrix {
	if Ap.n != Bp.n || Ap.m != Bp.m {
		panic(fmt.Sprintf("add:\n\tmatrix A shape: %vx%v;\n\tmatrix B shape: %vx%v;", Ap.m, Ap.n, Bp.m, Bp.n))
	}

	C := RMatrix{
		cols: make([][]float64, Ap.m),
		m:    Ap.m,
		n:    Ap.n,
	}

	for i := 0; i < C.m; i++ {
		C.cols[i] = make([]float64, C.n)
		for j := 0; j < C.n; j++ {
			C.cols[i][j] = Ap.cols[i][j] + Bp.cols[i][j]
		}
	}

	return &C
}

func (Ap *RMatrix) GetEl(i, j int) float64 {
	return Ap.cols[i][j]
}

func (Ap *RMatrix) SetEl(i, j int, val float64) {
	Ap.cols[i][j] = val
}

func (Ap *RMatrix) GetCol(i int) []float64 {
	return Ap.cols[i]
}

// others matrix
func (Ap *RMatrix) MulByPMatrix(Pp *PMatrix) *RMatrix {
	if Ap.n != Pp.n {
		panic(fmt.Sprintf("mul R on P:\n\tmatrix A shape: %vx%v;\n\tmatrix B shape: %vx%v;", Ap.m, Ap.n, Pp.n, Pp.n))
	}

	for i := 0; i < Pp.n; i++ {
		// swap rows
		j := Pp.perm[i]
		// либо перетсановку уже сделали, либо делать ее не надо
		if j <= i {
			continue
		}

		for k := 0; k < Ap.m; k++ {
			Ap.cols[k][i], Ap.cols[k][j] = Ap.cols[k][j], Ap.cols[k][i]
		}
	}

	return Ap
}
