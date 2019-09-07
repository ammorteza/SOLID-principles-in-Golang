package AreaCalculator

import "fmt"

type CalculatorOutputter struct {

}

func (co *CalculatorOutputter)PrintCircleArea(result AreaResult){
	fmt.Printf("circle area: %f\n", result.circleArea)
}

func (co *CalculatorOutputter)PrintSquareArea(result AreaResult)  {
	fmt.Printf("square area: %f\n", result.squareArea)
}