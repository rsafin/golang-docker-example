package main

import (
	"exchangerate/internal/bot"
	"exchangerate/internal/config"
	"exchangerate/internal/exchangerate"
	"fmt"
	"log"
	"time"
)

func checkRate(c chan float64){
	var oldRate float64 = 0

	p := new(exchangerate.Profinance)
	for {
		rate, err := p.GetRate("")
		if err != nil {
			log.Panic(err)
		}

		if oldRate != rate {
			oldRate = rate
			c <- rate
		}

		time.Sleep(60 * time.Second)
	}
}

func main() {
	cfg, err := config.NewConfig("config.yml")
	if err != nil {
		log.Panic(err)
	}

	var c = make(chan float64)
	go checkRate(c)

	go bot.CreateBot(cfg.App.Bot.Key)

	for msg := range c {
		fmt.Printf("%f\n", msg)
	}
}