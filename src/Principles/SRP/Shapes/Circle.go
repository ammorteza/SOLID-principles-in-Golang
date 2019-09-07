package Shapes

type Circle struct{
	radius	float64
}

func (r *Circle)GetRadius()  float64{
	return r.radius
}

func (r *Circle) Make(radius float64){
	r.radius = radius
}