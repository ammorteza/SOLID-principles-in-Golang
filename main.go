package main

import (
	"./src/Principles/SRP"
	"./src/Principles/OCP"
	"./src/Principles/LSP"
	"./src/Principles/ISP"
	"./src/Principles/DIP"
)

func main(){
	SRP.Run()
	OCP.Run()
	LSP.Run()
	ISP.Run()
	DIP.Run()
}
