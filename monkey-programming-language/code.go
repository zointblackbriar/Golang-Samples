package main

type Instructions []byte

type Opcode byte

import "fmt"

type Definition struct {
	Name string 
	OperandWidths []int
}

var definitions = map[Opcode]*Definition {
	OpConstant: {"OpConstant", []int{2}}
}

func main() {

}