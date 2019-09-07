package Shapes

type Square struct{
	name 	string
	height	float64
}

func (s *Square)GetHeight() float64{
	return s.height
}

func (s Square)GetName() string{
	return s.name
}

func (h *Square) Make(name string, height float64){
	h.name = name
	h.height = height
}

func (s Square)CalcArea() float64{
	return s.GetHeight() * 3.14 //for example this is wrong
}