package matrix

type Shaper interface {
	Shape() (m, n float64)
}

type ColSwaper interface {
	SwapCol(i, j int)
}

type ElGetter interface {
	GetEl(i, j int) float64
}
