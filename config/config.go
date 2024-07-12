package config

import (
	"github.com/spf13/viper"
	"log"
)

func InitConfig(filename string) (*viper.Viper, error) {
	config := viper.New()
	config.SetConfigFile(".env")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("Error parsing config file", err)
	}
	return config, nil
}
