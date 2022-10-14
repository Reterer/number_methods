package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

type Point struct {
	x, y float64
}
type ftype func(float64) float64

func firstDerivative(points []Point) func(float64) float64 {
	n := len(points) - 2
	a := make([]float64, n)
	b := make([]float64, n)

	for i := 0; i < n; i++ {
		dy1 := points[i+1].y - points[i].y
		dy2 := points[i+2].y - points[i+1].y
		dx1 := points[i+1].x - points[i].x
		dx2 := points[i+2].x - points[i+1].x
		dxm := points[i+2].x - points[i].x

		a[i] = dy1 / dx1
		b[i] = (dy2/dx2 - dy1/dx1) / dxm
	}
	return func(x float64) float64 {
		i := 0
		for ; points[i+1].x < x; i++ {
		}
		return a[i] + b[i]*(2*x-points[i].x-points[i+1].x)
	}
}

func secondDerivative(points []Point) func(float64) float64 {
	n := len(points) - 2
	a := make([]float64, n)

	for i := 0; i < n; i++ {
		dy1 := points[i+1].y - points[i].y
		dy2 := points[i+2].y - points[i+1].y
		dx1 := points[i+1].x - points[i].x
		dx2 := points[i+2].x - points[i+1].x
		dxm := points[i+2].x - points[i].x

		a[i] = 2 * (dy2/dx2 - dy1/dx1) / dxm
	}

	return func(x float64) float64 {
		i := 0
		for ; points[i+1].x < x; i++ {
		}
		return a[i]
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

func genPlot(path string, f ftype, ff ftype, points []Point, a float64, b float64, h float64) {
	p := plot.New()

	p.Title.Text = "Interpolation"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	steps := int((b - a) / h)
	f_p := make(plotter.XYs, steps)
	s_p := make(plotter.XYs, steps)
	x := a
	for step := 0; step < steps; step++ {
		f_p[step].X = x
		f_p[step].Y = f(x)
		s_p[step].X = x
		s_p[step].Y = ff(x)

		x += h
	}
	err := plotutil.AddLinePoints(p,
		"first derivative", f_p,
		"second derivative", s_p,
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
	fmt.Println(points)
	f := firstDerivative(points)
	fmt.Println("f'", f(0.2))

	ff := secondDerivative(points)
	fmt.Println("f\"", f(0.2))

	genPlot(outputFile, f, ff, points, points[0].x, points[len(points)-3].x, 0.01)
}
