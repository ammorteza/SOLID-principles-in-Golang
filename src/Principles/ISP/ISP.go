package ISP

import (
	"./Shapes"
	"fmt"
	"./AreaCalculator"
)

func Run(){
	fmt.Println("Interface Segregation principle (ISP)")

	var circle Shapes.Circle
	var square Shapes.Square
	var rec	   Shapes.Rectangle

	var calcOutputter AreaCalculator.CalculatorOutputter
	var areaCalc AreaCalculator.AreaCalculator

	circle.Make("circle")
	square.Make("square")
	rec.Make("rectangle")

	var _circle Shapes.I_ShapeCircle = &circle
	var _square Shapes.I_ShapeSquare = &square
	var _rec Shapes.I_ShapeRectangle = &rec

	_circle.SetRadius(12.6)
	_square.SetHeight(4.3)
	_rec.SetHeight(6.5)
	_rec.SetWidth(4.8)

	areaCalc.AddShape(_circle)
	areaCalc.AddShape(_square)
	areaCalc.AddShape(_rec)

	areaResult := areaCalc.GetShapesArea()
	fmt.Println(string(calcOutputter.JSON(areaResult)))
	fmt.Println("........................................................")
}
