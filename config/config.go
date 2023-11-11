package config

import "github.com/spf13/viper"

type Config struct {
	ApiKey string
}

var cfg Config

func LoadConfig() {
	viper.AutomaticEnv()

	cfg.ApiKey = viper.GetString("APIKEY")

}

func GetConfig() Config {
	return cfg
}
