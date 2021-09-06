package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Conf = new(config)

type config struct {
	Playlist []int `mapstructure:"playlist"`
	Accounts []struct {
		Phone    string `mapstructure:"phone"`
		Passwd   string `mapstructure:"passwd"`
		Expired  string `mapstructure:"expired"`
		OnlySign bool   `mapstructure:"only_sign"`
	} `mapstructure:"accounts"`
}

func InitConfig() {
	// Set configuration file path
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	// Read configuration
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error reading configuration file: %s \n", err))
	}
	// Unmarshal configuration
	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("unmarshal configuration failed, err: %s \n", err))
	}
}
