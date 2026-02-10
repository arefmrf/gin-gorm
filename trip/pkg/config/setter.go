package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

func Set() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	//viper.AddConfigPath("config")

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("cannot get current file path")
	}
	projectRoot := filepath.Join(filepath.Dir(filename), "..", "..") // adjust to reach project root
	configPath := filepath.Join(projectRoot, "config")
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file, %s", err)
	}
	if err := viper.Unmarshal(&configurations); err != nil {
		log.Fatalf("Error parsing config file, %s", err)
	}
}
