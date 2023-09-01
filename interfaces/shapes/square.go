package shapes

type Square struct {
	Side uint
}

func (s Square) Area() float64 {
	return float64(s.Side * s.Side)
}
