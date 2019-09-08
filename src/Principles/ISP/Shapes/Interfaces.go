package Shapes

type ShapeInterface interface {
		CalcArea()	float64
		GetName()		string
}

type I_ShapeCircle interface {
		ShapeInterface
		SetRadius(float64)
}

type I_ShapeRectangle interface {
		ShapeInterface
		SetWidth(float64)
		SetHeight(float64)
}

type I_ShapeSquare interface {
		ShapeInterface
		SetHeight(float64)
}
