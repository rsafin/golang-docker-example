package main

import (
	"exchangerate/internal/exchangerate"
	"fmt"
	"log"
	"time"
)

func checkRate(c chan float64){
	p := new(exchangerate.Profinance)
	for {
		rate, err := p.GetRate("")
		if err != nil {
			log.Panic(err)
		}
		c <- rate
		time.Sleep(5 * time.Second)
	}
}

func main() {
	var c = make(chan float64)
	go checkRate(c)

	for msg := range c {
		fmt.Printf("%f\n", msg)
	}
}