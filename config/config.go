package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AddConfigPath(".")
	viper.AddConfigPath("app")
	viper.AddConfigPath("config")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("could not find env file: %w", err))
	}
}

func BasicAuthUsername() string {
	return viper.GetString("BASIC_USER")
}

func BasicAuthPasswd() string {
	return viper.GetString("BASIC_PASSWORD")
}

func DBPath() string {
	return viper.GetString("DB_CONN")
}
