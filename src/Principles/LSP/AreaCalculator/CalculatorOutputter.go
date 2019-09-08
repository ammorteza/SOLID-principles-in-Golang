package AreaCalculator

import (
	json2 "encoding/json"
)

type CalculatorOutputter struct {

}

func (co *CalculatorOutputter)JSON(result AreaResult) []byte{
	json, err := json2.Marshal(result)
	if err != nil{
		panic("error in convert map to json")
	}
	return json
}