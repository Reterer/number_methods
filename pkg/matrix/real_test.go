package matrix

import (
	"fmt"
	"reflect"
	"testing"
)

func TestShaper(t *testing.T) {
	A := MakeRealMatrix(2, 3)
	shaper := Shaper(A)

	if n, m := shaper.Shape(); !(n == 2 && m == 3) {
		t.Errorf("shape A is not 2x3")
	}
}

func TestColSwapperAndSetElAndGetCol(t *testing.T) {
	A := MakeRealMatrix(2, 3)

	A.SetEl(0, 0, 1)
	A.SetEl(0, 1, 2)
	A.SetEl(0, 2, 3)
	copy(A.GetCol(1), []float64{4, 5, 6})

	colSwapper := ColSwapper(A)
	colSwapper.SwapCol(0, 1)

	if !(reflect.DeepEqual(A.cols[0], []float64{4, 5, 6})) {
		t.Errorf("A first col is not eq [4,5,6]")
	}
	if !(reflect.DeepEqual(A.cols[1], []float64{1, 2, 3})) {
		t.Errorf("A second col is not eq [1,2,3]")
	}
}

func TestGetter(t *testing.T) {
	A := MakeRealMatrix(2, 3)
	elGet := ElGetter(A)

	A.SetEl(1, 2, 42)
	if !(elGet.GetEl(1, 2) == 42) {
		t.Errorf("A[1][2] is not eq 42")
	}
	if !(elGet.GetEl(0, 2) == 0) {
		t.Errorf("A[0][2] is not eq 0")
	}
}

func TestMulByR(t *testing.T) {
	A := MakeRealMatrix(3, 2)
	B := MakeRealMatrix(2, 3)

	copy(A.cols[0], []float64{0, 1})
	copy(A.cols[1], []float64{2, 3})
	copy(A.cols[2], []float64{4, 5})

	copy(B.cols[0], []float64{0, 1, 2})
	copy(B.cols[1], []float64{3, 4, 5})

	CAns := MakeRealMatrix(3, 3)
	copy(CAns.cols[0], []float64{3, 4, 5})
	copy(CAns.cols[1], []float64{9, 14, 19})
	copy(CAns.cols[2], []float64{15, 24, 33})

	if !(reflect.DeepEqual(CAns.cols, A.MulByR(B).cols)) {
		t.Errorf("A * B is not correct")
		fmt.Println(A.cols, B.cols, A.MulByR(B).cols, CAns.cols)
	}
}

func TestMulByConstant(t *testing.T) {
	A := MakeRealMatrix(2, 2)

	copy(A.cols[0], []float64{0, 1})
	copy(A.cols[1], []float64{2, 3})
	consant := float64(2)

	AAns := MakeRealMatrix(2, 2)
	copy(AAns.cols[0], []float64{0, 2})
	copy(AAns.cols[1], []float64{4, 6})

	if !(reflect.DeepEqual(AAns.cols, A.MulByConstant(consant).cols)) {
		t.Errorf("A * Constant is not correct")
	}
}

func TestAdd(t *testing.T) {
	A := MakeRealMatrix(3, 2)
	B := MakeRealMatrix(3, 2)

	copy(A.cols[0], []float64{0, 1})
	copy(A.cols[1], []float64{2, 3})
	copy(A.cols[2], []float64{4, 5})

	copy(B.cols[0], []float64{0, 1})
	copy(B.cols[1], []float64{2, 3})
	copy(B.cols[2], []float64{4, 5})

	CAns := MakeRealMatrix(3, 2)
	copy(CAns.cols[0], []float64{0, 2})
	copy(CAns.cols[1], []float64{4, 6})
	copy(CAns.cols[2], []float64{8, 10})

	if !(reflect.DeepEqual(CAns.cols, A.Add(B).cols)) {
		t.Errorf("A * B is not correct")
	}
}

func TestRMulP(t *testing.T) {
	A := MakeRealMatrix(3, 2)
	P := MakePermutationMatrix(2)

	copy(A.cols[0], []float64{0, 1})
	copy(A.cols[1], []float64{2, 3})
	copy(A.cols[2], []float64{4, 5})
	P.SwapCol(0, 1)

	CAns := MakeRealMatrix(3, 2)
	copy(CAns.cols[0], []float64{1, 0})
	copy(CAns.cols[1], []float64{3, 2})
	copy(CAns.cols[2], []float64{5, 4})

	if !(reflect.DeepEqual(CAns.cols, A.MulByPMatrix(P).cols)) {
		t.Errorf("A * P is not correct")
	}
}
