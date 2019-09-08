package Shapes

type Circle struct{
	name string
	radius	float64
}

func (r *Circle) GetRadius()  float64{
	return r.radius
}

func (c Circle) GetName() string{
	return c.name
}

func (r *Circle) Make(name string){
	r.name = name

}

func (r *Circle) SetHeight(h float64){

}

func (r *Circle) SetWidth(w float64){

}

func (r *Circle) SetRadius(s float64){
		r.radius = s
}

func (c Circle)CalcArea() float64{
		return c.GetRadius() * c.GetRadius()
}
