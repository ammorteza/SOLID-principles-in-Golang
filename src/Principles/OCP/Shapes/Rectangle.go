package Shapes

type Rectangle struct{
	name string
	width	float64
	height	float64
}

func (r *Rectangle)GetW_and_H()  (float64, float64){
	return r.width, r.height
}

func (r Rectangle)GetName() string{
	return r.name
}

func (r *Rectangle) Make(name string, width float64, height float64){
	r.name = name
	r.width = width
	r.height = height
}

func (r Rectangle)CalcArea() float64{
	return r.width * r.height
}