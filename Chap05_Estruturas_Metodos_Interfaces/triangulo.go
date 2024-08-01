package geometria

type Triangulo struct {
	Altura float64
	Base float64
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) * 0.5
}
