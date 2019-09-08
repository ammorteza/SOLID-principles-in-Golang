package Shapes

type ShapeInterface interface {
		CalcArea()	float64
		GetName()		string
		SetWidth(float64)
		SetHeight(float64)
		SetRadius(float64)
}
