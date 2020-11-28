package main

import (
	"exchangerate/internal/exchangerate"
	"fmt"
)

func main() {
	p := new(exchangerate.Profinance)

	rate, _ := p.GetRate("")

	fmt.Printf("%v\n", rate)
}