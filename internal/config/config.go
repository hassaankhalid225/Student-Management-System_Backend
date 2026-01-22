package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	DBName         string `mapstructure:"DB_NAME"`
	RedisHost      string `mapstructure:"REDIS_HOST"`
	RedisPort      string `mapstructure:"REDIS_PORT"`
	JWTSecret      string `mapstructure:"JWT_SECRET"`
	RefreshSecret  string `mapstructure:"REFRESH_SECRET"`
	AppPort        string `mapstructure:"APP_PORT"`
	Environment    string `mapstructure:"ENVIRONMENT"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("Warning: .env file not found, using environment variables: %v", err)
	}

	err = viper.Unmarshal(&config)
	return
}
