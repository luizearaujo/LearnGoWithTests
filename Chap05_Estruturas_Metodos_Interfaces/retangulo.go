package geometria

type Retangulo struct{
	Largura float64
	Altura float64
}

func (r Retangulo) Perimetro() float64 {
	return (r.Largura + r.Altura) * 2
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}
