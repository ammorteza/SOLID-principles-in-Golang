package AreaCalculator

import (
	"../Shapes"
)

type AreaResult struct {
	circleArea float64
	squareArea float64
}

type AreaCalculator struct {
	circle 	Shapes.Circle
	square 	Shapes.Square
}

func (ac *AreaCalculator)SetCircle(c Shapes.Circle){
	ac.circle = c
}

func (ac *AreaCalculator)SetSquare(s Shapes.Square){
	ac.square = s
}

func (ac *AreaCalculator)CalcCircle(ar *AreaResult){
	ar.circleArea = ac.circle.GetRadius() * ac.circle.GetRadius()
}

func (ac *AreaCalculator)CalcSquare(ar *AreaResult){
	ar.squareArea = ac.square.GetHeight() * 3.14 //for example this is wrong
}