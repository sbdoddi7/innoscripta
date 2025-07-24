package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port        string
	PostgresDSN string
}

// LoadConfig retrices variables from envirnoment
func LoadConfig() *Config {
	viper.SetDefault("PORT", "8080")
	viper.AutomaticEnv()

	cfg := &Config{
		Port:        viper.GetString("PORT"),
		PostgresDSN: viper.GetString("POSTGRES_DSN"),
	}

	log.Printf("Loaded config: %+v", cfg)
	return cfg
}
