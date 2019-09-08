package Shapes

type Circle struct{
		Shape
		radius	float64
}

func (r *Circle) GetRadius()  float64{
		return r.radius
}

func (r *Circle) Make(name string){
		r.name = name
}

func (r *Circle) SetRadius(s float64){
		r.radius = s
}

func (c Circle)CalcArea() float64{
		return c.GetRadius() * c.GetRadius()
}
