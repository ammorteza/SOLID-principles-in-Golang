package Shapes

type Rectangle struct{
	Shape
	width	float64
	height	float64
}

func (r *Rectangle)GetW_and_H()  (float64, float64){
	return r.width, r.height
}

func (r *Rectangle) Make(name string){
	r.name = name
}

func (r *Rectangle) SetHeight(h float64){
	r.height = h
}

func (r *Rectangle) SetWidth(w float64){
	r.width = w
}

func (r Rectangle)CalcArea() float64{
	return r.width * r.height
}
