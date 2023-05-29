package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("could not find env file: %w", err))
	}
}

func BasicAuthUsername() string {
	return viper.GetString("BASIC_USERNAME")
}

func BasicAuthPasswd() string {
	return viper.GetString("BASIC_PASSWORD")
}

func DBPath() string {
	return viper.GetString("DB_CONN")
}
