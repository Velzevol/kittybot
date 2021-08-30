package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	TelegramBotToken string
}

func main() {
	configfile, _ := os.Open("config/config.json")
	decoder := json.NewDecoder(configfile)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(configuration.TelegramBotToken)
}
