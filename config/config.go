package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Conf = new(config)

type config struct {
	Level    string `mapstructure:"level"`
	Interval int    `mapstructure:"interval"` // 单个任务之间时间间隔，单位：秒
	Catnap   struct {
		Number   int `mapstructure:"number"`   // 短睡眠数量
		Duration int `mapstructure:"duration"` // 短睡眠时间，单位：分钟
	} `mapstructure:"catnap"`
	Sleep struct {
		Number   int `mapstructure:"number"`   // 长睡眠数量
		Duration int `mapstructure:"duration"` // 长睡眠时间，单位：分钟
	} `mapstructure:"sleep"`
	Playlist []int `mapstructure:"playlist"` // 自定义歌单
	Accounts []struct {
		Phone    string `mapstructure:"phone"`     // 手机号
		Passwd   string `mapstructure:"passwd"`    // 密码
		Expired  string `mapstructure:"expired"`   // 任务过期时间
		OnlySign bool   `mapstructure:"only_sign"` // 只进行签到
		Unstable bool   `mapstructure:"unstable"`  // 不稳定刷歌
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
