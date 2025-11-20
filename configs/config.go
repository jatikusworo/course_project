package configs

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	AppName     string `mapstructure:"appName"`
	Env         string `mapstructure:"env"`
	DatabaseURL string `mapstructure:"databaseURL"`
	Port        string `mapstructure:"port"`
}

func Load() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	viper.SetConfigName("config." + env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Error parsing config: %v", err)
	}

	return &cfg
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
