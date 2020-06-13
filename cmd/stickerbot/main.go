package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/millfort/imgfit/bot"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("bot")
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exeDir := filepath.Dir(exePath)
	viper.AddConfigPath(filepath.Join(exeDir, "config"))
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	tgToken := viper.GetString("token")
	fitbot := bot.New(tgToken)
	fitbot.Start()
}
