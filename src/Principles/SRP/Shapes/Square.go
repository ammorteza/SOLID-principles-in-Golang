package Shapes

type Square struct{
	height	float64
}

func (s *Square)GetHeight() float64{
	return s.height
}

func (h *Square) Make(height float64){
	h.height = height
}