package main

import "fmt"

type Calcutation struct {
	ID         string `json:"id"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
}

type CalcutationRequest struct {
	Expression string `json:"expression"`
}

var calculation = []Calcutation{} //nessesary to initilization like a slice

func main() {
	fmt.Println("hey its calculator")
}
