package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port         string `mapstructure:"PORT"`
	Username     string `mapstructure:"USER_NAME"`
	Email        string `mapstructure:"EMAIL"`
	SmtpHost     string `mapstructure:"SMTP_HOST"`
	SmtpPort     string `mapstructure:"SMTP_PORT"`
	SmtpPassword string `mapstructure:"SMTP_PASSWORD"`
}

func InitConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	return

}

func GetConfig() Config {
	config, err := InitConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	return config
}
