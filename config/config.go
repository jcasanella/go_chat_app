package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type ConfigValues struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func NewConfigValues() *ConfigValues {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	host := getValueString("postgres.address.host")
	port := getValueInt("postgres.address.port")
	username := getValueString("postgres.user.username")
	password := getValueString("postgres.user.password")
	database := getValueString("postgres.database")

	c := &ConfigValues{Host: host, Port: port, Username: username, Password: password, Database: database}
	return c
}

func getValueString(key string) string {
	value := viper.GetString(key)
	if value == "" {
		panic(fmt.Errorf("fatal error %s key not found", key))
	}

	return value
}

func getValueInt(key string) int {
	value := viper.GetInt(key)
	if value == 0 {
		panic(fmt.Errorf("fatal error %s key not found", key))
	}

	return value
}
