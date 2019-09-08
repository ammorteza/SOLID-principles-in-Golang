package AreaCalculator

import (
	"../Shapes"
)

type AreaResult map[string]float64

type AreaCalculator struct {
	shapes []Shapes.ShapeInterface
}

func (ac *AreaCalculator)AddShape(si Shapes.ShapeInterface){
	ac.shapes = append(ac.shapes,si)
}

func (ac *AreaCalculator)GetShapesArea() AreaResult{
	result := make(AreaResult)
	for i := 0; i < len(ac.shapes); i++ {
		result[ac.shapes[i].GetName()] = ac.shapes[i].CalcArea()
	}
	return result
}
