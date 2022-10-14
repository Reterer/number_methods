package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

type Point struct {
	x, y float64
}

type ftype func(float64) float64

func MakeLangrangeInterpolation(points []Point) func(float64) float64 {
	n := len(points)
	l := make([]float64, n)
	for i := 0; i < n; i++ {
		var w float64 = 1
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			w *= points[i].x - points[j].x
		}

		l[i] = points[i].y / w
	}

	return func(x float64) float64 {
		var res float64
		for i := 0; i < n; i++ {
			var xx float64 = 1
			for j := 0; j < n; j++ {
				if j == i {
					continue
				}
				xx *= x - points[j].x
			}
			res += l[i] * xx
		}
		return res
	}
}

func MakeNewtonInterpolation(points []Point) func(float64) float64 {
	n := len(points)
	mat := make([][]float64, n)
	mat[0] = make([]float64, n)
	for i := 0; i < n; i++ {
		mat[0][i] = points[i].y
	}

	for i := 1; i < n; i++ {
		mat[i] = make([]float64, n-i)
		for j := i; j < n; j++ {
			mat[i][j-i] = (mat[i-1][j-i] - mat[i-1][j-i+1]) / (points[j-i].x - points[j].x)
		}
	}

	return func(x float64) float64 {
		ans := float64(0)
		pow := float64(1)
		for i := 0; i < n; i++ {
			ans += pow * mat[i][0]
			pow *= x - points[i].x
		}
		return ans
	}
}

func f(x float64) float64 {
	return math.Log(x)
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

func genPlot(path string, lf ftype, nf ftype, f ftype, a float64, b float64, h float64) {
	p := plot.New()

	p.Title.Text = "Interpolation"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	steps := int((b - a) / h)
	o_p := make(plotter.XYs, steps)
	l_p := make(plotter.XYs, steps)
	n_p := make(plotter.XYs, steps)
	x := a
	for step := 0; step < steps; step++ {
		o_p[step].X = x
		o_p[step].Y = f(x)
		l_p[step].X = x
		l_p[step].Y = lf(x)
		n_p[step].X = x
		n_p[step].Y = nf(x)

		x += h
	}

	err := plotutil.AddLinePoints(p,
		"Original", o_p,
		"Langrange", l_p,
		"Newton", n_p)
	if err != nil {
		panic(err)
	}

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
	lf, nf := MakeLangrangeInterpolation(points), MakeNewtonInterpolation(points)
	eps := math.Abs(f(0.8) - lf(0.8))

	fmt.Println("Значение интерполяционного многочлена: ", lf(0.8), "Значение погрешности: ", eps)
	fmt.Println("Значение интерполяционного многочлена: ", nf(0.8), "Значение погрешности: ", eps)

	genPlot(outputFile, lf, nf, f, points[0].x, points[len(points)-1].x, 0.01)
}
