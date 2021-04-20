package util

import (
	"github.com/spf13/viper"
	"time"
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
	viper.AddConfigPath(path)
	viper.SetConfigName("test")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}

func LoadTestConfig() Config {
	var config Config
	config.DBDriver = "postgres"
	config.DBHost = "0.0.0.0"
	config.DBPort = "5432"
	config.DBUser = "hello"
	config.DBPassword = "Hello123@"
	config.DBName = "hello"
	config.ServerAddress = "0.0.0.0:8080"
	config.JWTTokenKey = "12345678901234567890123456789012"
	config.AccessTokenDuration = time.Minute * 30
	return config
}
