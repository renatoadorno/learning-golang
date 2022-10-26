package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Environment string
	Port        string
	Mongo       MongoConfig
}

type MongoConfig struct {
	Server   string
	Database string
}

func GetConfig() Config {
	conf := Config{}

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	return conf
}
