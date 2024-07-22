package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Database struct {
		User     string
		Password string
		DBName   string
		Host     string
		Port     int
		SSLMode  string
	}
	Kafka struct {
		Brokers []string
	}
	Server struct {
		Port int
	}
}

var AppConfig Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}
