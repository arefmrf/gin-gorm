package config

import (
	"github.com/spf13/viper"
	"log"
)

func Set() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&configurations); err != nil {
		log.Fatalf("Error parsing config file, %s", err)
	}
}
