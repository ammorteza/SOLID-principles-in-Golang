package Shapes

type Circle struct{
	name string
	radius	float64
}

func (r *Circle)GetRadius()  float64{
	return r.radius
}

func (c Circle)GetName() string{
	return c.name
}

func (r *Circle) Make(name string, radius float64){
	r.name = name
	r.radius = radius
}

func (c Circle)CalcArea() float64{
	return c.GetRadius() * c.GetRadius()
}