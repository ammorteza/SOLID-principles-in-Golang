package OCP

import (
	"./Shapes"
	"fmt"
	"./AreaCalculator"
)

func Run(){
	fmt.Println("Open-Closed principle (OCP)")

	var circle Shapes.Circle
	var square Shapes.Square
	var rec	   Shapes.Rectangle

	var calcOutputter AreaCalculator.CalculatorOutputter
	var areaCalc AreaCalculator.AreaCalculator

	circle.Make("circle",6.63)
	square.Make("square", 12.5)
	rec.Make("rectangle", 2.6, 4.78)

	var _circle Shapes.ShapeInterface = circle
	var _square Shapes.ShapeInterface = square
	var _rec Shapes.ShapeInterface = rec

	areaCalc.AddShape(_circle)
	areaCalc.AddShape(_square)
	areaCalc.AddShape(_rec)

	areaResult := areaCalc.GetShapesArea()
	fmt.Println(string(calcOutputter.JSON(areaResult)))
	fmt.Println("........................................................")
}