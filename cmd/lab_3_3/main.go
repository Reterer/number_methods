package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"

	"github.com/Reterer/number_methods/internal/lu_decompose"
	"github.com/Reterer/number_methods/pkg/matrix"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

type Point struct {
	x, y float64
}
type ftype func(float64) float64

func squareError(points []Point, f func(float64) float64) float64 {
	var err float64

	for i := 0; i < len(points); i++ {
		err += math.Pow(f(points[i].x)-points[i].y, 2)
	}

	return err
}

func lsm(points []Point, n int) func(float64) float64 {
	N := len(points)
	n1 := n + 1
	a := make([]float64, n1)

	{
		// Делаем систему
		A := matrix.MakeRealMatrix(n1, n1)
		b := matrix.MakeRealMatrix(n1, 1)

		for k := 0; k < n1; k++ {
			for i := 0; i < n1; i++ {
				var sumA, sumB float64
				for j := 0; j < N; j++ {
					sumA += math.Pow(points[j].x, float64(k+i))
					sumB += math.Pow(points[j].x, float64(k)) * points[j].y
				}
				A.SetEl(k, i, sumA)
				b.SetEl(k, 0, sumB)
			}
		}
		// Решаем систему
		LU := lu_decompose.MakeLU(lu_decompose.PermMin, A)
		LU.Decompose()
		aMat := LU.Solve(b)
		for i := 0; i < n1; i++ {
			a[i] = aMat.GetEl(i, 0)
		}
	}
	return func(x float64) float64 {
		var ans float64
		var xk float64 = 1

		for i := 0; i < n1; i++ {
			ans += xk * a[i]
			xk *= x
		}

		return ans
	}
}

func readFromFile(filePath string) []Point {
	f, err := os.Open(filePath)
	if err != nil {
		panic("Unable to read input file " + filePath + " " + err.Error())
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic("Unable to parse file as CSV for " + filePath + " " + err.Error())
	}

	points := make([]Point, len(records))
	for i := 0; i < len(records); i++ {
		_, err := fmt.Sscanf(records[i][0], "%f", &points[i].x)
		if err != nil {
			panic(err.Error())
		}
		_, err = fmt.Sscanf(records[i][1], "%f", &points[i].y)
		if err != nil {
			panic(err.Error())
		}
	}
	return points
}

func genPlot(path string, f_1 ftype, f_2 ftype, points []Point, a float64, b float64, h float64) {
	p := plot.New()

	p.Title.Text = "Interpolation"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	steps := int((b - a) / h)
	f1_p := make(plotter.XYs, steps)
	f2_p := make(plotter.XYs, steps)
	x := a
	for step := 0; step < steps; step++ {
		f1_p[step].X = x
		f1_p[step].Y = f_1(x)
		f2_p[step].X = x
		f2_p[step].Y = f_2(x)

		x += h
	}
	err := plotutil.AddLinePoints(p,
		"lsm-1", f1_p,
		"lsm-2", f2_p,
	)
	if err != nil {
		panic(err)
	}

	// Scatter
	scatter_data := make(plotter.XYs, len(points))
	for i := 0; i < len(points); i++ {
		scatter_data[i].X = points[i].x
		scatter_data[i].Y = points[i].y
	}
	s, err := plotter.NewScatter(scatter_data)
	if err != nil {
		panic(err)
	}
	s.GlyphStyle.Radius = 10
	p.Add(s)
	p.Legend.Add("Points", s)
	// Save the plot to a PNG file.
	if err := p.Save(2000, 2000, path); err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		panic("Аргументов должно быть два")
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	points := readFromFile(inputFile)

	f_1 := lsm(points, 1)
	fmt.Println("a + bx: ", f_1(5.1), " serr: ", squareError(points, f_1))

	f_2 := lsm(points, 2)
	fmt.Println("a + bx + cx^2: ", f_2(5.1), " serr: ", squareError(points, f_2))

	genPlot(outputFile, f_1, f_2, points, points[0].x, points[len(points)-1].x, 0.1)

}
