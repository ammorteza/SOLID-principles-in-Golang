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

func (h *Square) Make(name string){
	h.name = name
}

func (r *Square) SetHeight(h float64){
	r.height = h
}

func (r *Square) SetWidth(w float64){}

func (r *Square) SetRadius(s float64){}

func (s Square)CalcArea() float64{
	return s.GetHeight() * 3.14 //for example this is wrong
}
