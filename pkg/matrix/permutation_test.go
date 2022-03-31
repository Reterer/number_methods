package matrix

import (
	"reflect"
	"testing"
)

func TestGetEl(t *testing.T) {
	P := MakePermutationMatrix(3)

	elGetter := ElGetter(P)

	if !(elGetter.GetEl(0, 0) == float64(1)) {
		t.Errorf("P.GetEl(0,0) is not eq 1")
	}
	if !(elGetter.GetEl(1, 1) == float64(1)) {
		t.Errorf("P.GetEl(1,1) is not eq 1")
	}
	if !(elGetter.GetEl(2, 2) == float64(1)) {
		t.Errorf("P.GetEl(2,2) is not eq 1")
	}
	if !(elGetter.GetEl(2, 1) == float64(0)) {
		t.Errorf("P.GetEl(2,1) is not eq 0")
	}
}

func TestShape(t *testing.T) {
	P := MakePermutationMatrix(3)

	shaper := Shaper(P)

	if m, n := shaper.Shape(); !(m == 3 && n == 3) {
		t.Errorf("P.Shape() is not eq 3,3")
	}
}

func TestColSwapperAndGetValue(t *testing.T) {
	P := MakePermutationMatrix(3)

	colSwapper := ColSwapper(P)

	colSwapper.SwapCol(0, 1)
	if !(P.GetValue(0) == 1 && P.GetValue(1) == 0 && P.GetValue(2) == 2) {
		t.Errorf("P.perm is not eq [1, 0, 2]")
	}
	colSwapper.SwapCol(0, 1)
	if !(P.GetValue(0) == 0 && P.GetValue(1) == 1 && P.GetValue(2) == 2) {
		t.Errorf("P.perm is not eq [0, 1, 2]")
	}
}

func TestMulByPMatrix(t *testing.T) {
	A := MakePermutationMatrix(3)
	B := MakePermutationMatrix(3)

	A.perm = []int{2, 0, 1}
	B.perm = []int{2, 1, 0}
	ansPerm := []int{0, 2, 1}

	C := A.MulByPMatrix(B)
	if !(reflect.DeepEqual(C.perm, ansPerm)) {
		t.Errorf("C.perm is not eq [0, 2, 1]")
	}
}

func TestPMulR(t *testing.T) {
	P := MakePermutationMatrix(2)
	A := MakeRealMatrix(2, 3)

	copy(A.cols[0], []float64{0, 1, 2})
	copy(A.cols[1], []float64{3, 4, 5})
	P.SwapCol(0, 1)

	AAns := MakeRealMatrix(2, 3)
	copy(AAns.cols[0], []float64{3, 4, 5})
	copy(AAns.cols[1], []float64{0, 1, 2})

	if !(reflect.DeepEqual(AAns.cols, P.MulByRMatrix(A).cols)) {
		t.Errorf("P * A is not correct")
	}
}
