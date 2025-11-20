package configs

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	AppName    string `mapstructure:"app_name"`
	AppEnv     string `mapstructure:"app_env"`
	ServerPort string `mapstructure:"server_port"`

	Database struct {
		Host    string `mapstructure:"host"`
		Port    int    `mapstructure:"port"`
		User    string `mapstructure:"user"`
		Pass    string `mapstructure:"pass"`
		Name    string `mapstructure:"name"`
		SSLMode string `mapstructure:"sslmode"`
	} `mapstructure:"database"`

	Log struct {
		Level string `mapstructure:"level"`
	} `mapstructure:"log"`
}

func Load() Config {
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

	return cfg
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
