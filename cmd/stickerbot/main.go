package main

import (
	"log"

	"github.com/millfort/imgfit/bot"
	"github.com/spf13/viper"
)

func init() {
	viper.SetEnvPrefix("sb")
	err := viper.BindEnv("token")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	token := viper.GetString("token")
	fitbot := bot.New(token)
	fitbot.Start()
}
