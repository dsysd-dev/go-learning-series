package shapes

type Rectangle struct {
	Length, Breadth uint
}

func (s Rectangle) Area() float64 {
	return float64(s.Length * s.Breadth)
}
