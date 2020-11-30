package main

import (
	"exchangerate/internal/config"
	"exchangerate/internal/exchangerate"
	"fmt"
	/*"github.com/go-telegram-bot-api/telegram-bot-api"
	"fmt"*/
	"log"
	"time"
)

type Config struct {
	App struct{
		Bot struct {
			Key string `yaml:"key"`
		} `yaml:"bot"`
	} `yaml:"app"`
}

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
	cfg, err := config.NewConfig("config.yml")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%+v\n", cfg)


	/*bot*/
	/*bot, err := tgbotapi.NewBotAPI()
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}*/

	/*rate*/
	/*var c = make(chan float64)
	go checkRate(c)

	for msg := range c {
		fmt.Printf("%f\n", msg)
	}*/
}