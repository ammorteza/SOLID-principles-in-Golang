package SRP

import (
	"./Shapes"
	"fmt"
	"./AreaCalculator"
)

func Run(){
	fmt.Println("Single responsibility principle (SRP)")

	var circle Shapes.Circle
	var square Shapes.Square
	var areaResult AreaCalculator.AreaResult
	var calcOutputter AreaCalculator.CalculatorOutputter

	circle.Make(6.63)
	square.Make(12.5)

	var areaCalc AreaCalculator.AreaCalculator
	areaCalc.SetCircle(circle)
	areaCalc.SetSquare(square)

	areaCalc.CalcCircle(&areaResult)
	areaCalc.CalcSquare(&areaResult)

	calcOutputter.PrintCircleArea(areaResult)
	calcOutputter.PrintSquareArea(areaResult)

	fmt.Println("........................................................")
}