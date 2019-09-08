package Shapes

type Square struct{
	Shape
	height	float64
}

func (s *Square)GetHeight() float64{
	return s.height
}

func (h *Square) Make(name string){
	h.name = name
}

func (r *Square) SetHeight(h float64){
	r.height = h
}

func (s Square)CalcArea() float64{
	return s.GetHeight() * 3.14 //for example this is wrong
}
