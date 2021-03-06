package matrix

type Shaper interface {
	Shape() (n, m int)
}

type ColSwapper interface {
	SwapCol(i, j int)
}

type ElGetter interface {
	GetEl(i, j int) float64
}

type ShaperElGetter interface {
	ElGetter
	Shaper
}
