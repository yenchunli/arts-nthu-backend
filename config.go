package main

import (
	"github.com/spf13/viper"
	"time"
	"os"
)

type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBHost              string        `mapstructure:"DB_HOST"`
	DBPort              string        `mapstructure:"DB_PORT"`
	DBUser              string        `mapstructure:"DB_USER"`
	DBPassword          string        `mapstructure:"DB_PASSWORD"`
	DBName              string        `mapstructure:"DB_NAME"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	JWTTokenKey         string        `mapstructure:"JWT_TOKEN_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	if os.Getenv("GIN_MODE") == "debug" {
		viper.AddConfigPath(path)
		viper.SetConfigName("test")
		viper.SetConfigType("env")
	} else if os.Getenv("GIN_MODE") == "release" {
		viper.AddConfigPath(path)
		viper.SetConfigName("prod")
		viper.SetConfigType("env")
	}
	

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
