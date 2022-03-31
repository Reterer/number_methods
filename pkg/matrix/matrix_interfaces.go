package matrix

type Shaper interface {
	Shape() (m, n int)
}

type ColSwapper interface {
	SwapCol(i, j int)
}

type ElGetter interface {
	GetEl(i, j int) float64
}
