package shapes

type Circle struct {
	Radius uint
}

func (c Circle) Area() float64 {
	return float64(3.14) * float64(c.Radius) * float64(c.Radius)
}
